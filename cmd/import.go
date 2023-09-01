/*
Copyright © 2023 Hyungsuk Hong

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	_ "github.com/mattn/go-sqlite3" // import sqlite3.
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"

	sqlite "github.com/aanoaa/sgviz/internal/db"
	"github.com/aanoaa/sgviz/internal/reader"
	"github.com/aanoaa/sgviz/repo"
)

// importCmd represents the import command.
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Reads and stores HOST, GROUP, and POLICY information from a given <stdin> csv.",
	Long: `Reads and stores HOST, GROUP, and POLICY information from a given <stdin> csv.
The HOST CSV format should look like this

  foo-host01,192.168.0.1
  foo-host02,192.168.0.2
  bar-vserver,10.0.0.1

The header and desc columns are optional.

  hostname,ipaddr,desc
  foo-host01,192.168.0.1,brief of host
  foo-host02,192.168.0.2,
  bar-vserver,10.0.0.1,

For example:

  $ cat <<EOL | sgviz import --host
  foo-host01,192.168.0.1
  foo-host02,192.168.0.2
  bar-vserver,10.0.0.1
  EOL

  $ cat <<EOL | sgviz import --host
  hostname,ipaddr,desc
  foo-host01,192.168.0.1,brief of host
  foo-host02,192.168.0.2,
  bar-vserver,10.0.0.1,
  EOL

The GROUP CSV format should look like this

  group,hostname
  foo,foo-host01
  foo,foo-host02
  bar.example.com,bar-vserver

For example:

  $ cat <<EOL | sgviz import --group
  group,hostname
  foo,foo-host01
  foo,foo-host02
  bar.example.com,bar-vserver
  EOL

and The POLICY CSV format should look like this

  src,dst,port
  foo,bar.example.com,8080

For example:

  $ cat <<EOL | sgviz import --policy
  src,dst,port,desc
  foo,bar.example.com,8080,'foo to bar.example.com:8080'
  EOL

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(isHost || isGroup || isPolicy) {
			return errors.New("at least one of host|group|policy is required")
		}

		if os.Getenv("DEBUG") != "" {
			boil.DebugMode = true
		}

		var err error
		db, err = sqlite.Conn()
		if err != nil {
			return errors.Wrap(err, "db conn fail")
		}

		if isHost {
			err = importHost(cmd, args)
		} else if isGroup {
			err = importGroup(cmd, args)
		} else if isPolicy {
			err = importPolicy(cmd, args)
		}
		return errors.Wrap(err, "import fail")
	},
}

var (
	isHost              bool
	isGroup             bool
	isPolicy            bool
	fromHost            bool
	similarityThreshold float64
	zone                string

	db *sql.DB
)

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().BoolVarP(&isHost, "host", "", false, "given csv should be HOST format")
	importCmd.Flags().BoolVarP(&isGroup, "group", "", false, "given csv should be GROUP format")
	importCmd.Flags().BoolVarP(&isPolicy, "policy", "", false, "given csv should be POLICY format")
	importCmd.Flags().BoolVarP(&fromHost, "from-host", "", false, "generate group records from host data")
	importCmd.Flags().StringVarP(&zone, "zone", "z", "", "zone name where host located")

	similarityThreshold = *importCmd.Flags().Float64P("similarity", "", 1.0, "similarity threshold")
	importCmd.MarkFlagsMutuallyExclusive("host", "group", "policy")
}

func importHost(cmd *cobra.Command, _ []string) error {
	stdout := cmd.OutOrStdout()
	fmt.Fprintln(stdout, "<ctrl+d> to submit, <ctrl+c> to abort.")

	re := regexp.MustCompile(`hostname,ipaddr(,desc)?\n`)
	records, err := reader.StdinToCsv(re)
	if err != nil {
		return errors.Wrap(err, "stdin to csv fail")
	}

	ctx := context.Background()
	for _, record := range records {
		hr := repo.NewHostRepo(db)
		if err := hr.Upsert(ctx, record, zone); err != nil {
			return errors.Wrap(err, "upsert fail")
		}
	}

	fmt.Fprintf(stdout, "Updated %d records successfully\n", len(records))
	return nil
}

func importGroup(cmd *cobra.Command, _ []string) error {
	stdout := cmd.OutOrStdout()
	stderr := cmd.ErrOrStderr()
	hr := repo.NewHostRepo(db)
	gr := repo.NewGroupRepo(db)
	ctx := context.Background()

	if fromHost {
		list, err := hr.List(ctx, "")
		if err != nil {
			return errors.Wrap(err, "list fail")
		}

		hosts := make([]string, len(list))
		zones := make([]string, len(list))
		groupMap := make(map[string]bool)
		for i, host := range list {
			hosts[i] = host.Hostname
			zones[i] = host.Zone
			groupMap[strings.TrimFunc(host.Hostname, trimDigit)] = true
		}
		groups := make([]string, 0, len(groupMap))
		for k := range groupMap {
			groups = append(groups, k)
		}

		var similarity float64
		hamming := metrics.NewHamming()
		fmt.Fprintln(stdout, "group,hostname,zone")
		for i := range groups {
			for j := range hosts {
				zone := zones[j]
				group := groups[i]                                 // foo-bar-baz
				hostname := hosts[j]                               // foo-bar-baz01
				hostnameT := strings.TrimFunc(hostname, trimDigit) // foo-bar-baz

				groupSplit := strings.Split(group, "-")    // foo,bar,baz01
				hostSplit := strings.Split(hostnameT, "-") // foo,bar,baz

				// foo-bar-baz
				similarity = strutil.Similarity(group, hostnameT, hamming)
				if similarity >= similarityThreshold {
					if os.Getenv("DEBUG") != "" {
						fmt.Fprintf(stderr, "%s|%s: %.2f\n", group, hostnameT, similarity)
					}
					fmt.Fprintf(stdout, "%s,%s,%s\n", group, hostname, zone)
				}

				// bar-baz
				a := strings.Join(groupSplit[1:], "-")
				b := strings.Join(hostSplit[1:], "-")
				similarity = strutil.Similarity(a, b, hamming)
				if similarity >= similarityThreshold {
					fmt.Fprintf(stdout, "%s,%s,%s\n", a, hostname, zone)
				}

				// foo-bar
				a = strings.Join(groupSplit[:len(groupSplit)-1], "-")
				b = strings.Join(hostSplit[:len(hostSplit)-1], "-")
				similarity = strutil.Similarity(a, b, hamming)
				if similarity >= similarityThreshold {
					fmt.Fprintf(stdout, "%s,%s,%s\n", a, hostname, zone)
				}
			}
		}

		return nil
	}

	fmt.Fprintln(stdout, "<ctrl+d> to submit, <ctrl+c> to abort.")
	re := regexp.MustCompile(`group,hostname,zone\n`)
	records, err := reader.StdinToCsv(re)
	if err != nil {
		return errors.Wrap(err, "csv read fail")
	}

	for _, record := range records {
		if err := gr.Upsert(ctx, record); err != nil {
			return errors.Wrapf(err, "upsert fail: %v", record)
		}
	}

	fmt.Fprintf(stdout, "Updated %d records successfully\n", len(records))
	return nil
}

func importPolicy(cmd *cobra.Command, _ []string) error {
	stdout := cmd.OutOrStdout()
	fmt.Fprintln(stdout, "<ctrl+d> to submit, <ctrl+c> to abort.")

	re := regexp.MustCompile(`src,dst,port\n`)
	records, err := reader.StdinToCsv(re)
	if err != nil {
		return errors.Wrap(err, "csv read fail")
	}

	ctx := context.Background()
	for _, record := range records {
		pr := repo.NewPolicyRepo(db)
		if err := pr.Upsert(ctx, record); err != nil {
			return errors.Wrap(err, "upsert fail")
		}
		fmt.Fprintln(stdout, len(record), record)
	}

	return nil
}

func trimDigit(r rune) bool {
	return unicode.IsNumber(r)
}
