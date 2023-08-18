/*
Copyright © 2023 Hyunksuk Hong

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
	"fmt"
	"os"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3" // import sqlite3.
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"

	sqlite "github.com/aanoaa/sgviz/internal/db"
	"github.com/aanoaa/sgviz/internal/reader"
	"github.com/aanoaa/sgviz/repo"
)

// markdownCmd represents the markdown command.
var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Print policies as confluence markdown format from `$ sgviz policy`",
	Long: `Print policies as confluence markdown format from $ sgviz policy
For example:

  $ sgviz policy | sgviz markdown
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("DEBUG") != "" {
			boil.DebugMode = true
		}

		var err error
		db, err = sqlite.Conn()
		if err != nil {
			return errors.Wrap(err, "db conn fail")
		}

		stdout := cmd.OutOrStdout()
		fmt.Fprintln(stdout, "<ctrl+d> to submit, <ctrl+c> to abort.")

		re := regexp.MustCompile(`src,src_zone,dst,dst_zone,port,protocol\n`)
		records, err := reader.StdinToCsv(re)
		if err != nil {
			return errors.Wrap(err, "stdin to csv fail")
		}

		if len(records) == 0 {
			return errors.New("Empty input")
		}

		if len(records[0]) != 6 {
			return errors.New("Unexpected format")
		}

		b := new(strings.Builder)
		fmt.Fprintln(b, "||zone||group||host||dst||ip||port||domain||app||desc||remaining tasks||")

		ctx := context.Background()
		hr := repo.NewHostRepo(db)
		gr := repo.NewGroupRepo(db)
		for _, record := range records {
			src, srcZone, dst, _, port := record[0], record[1], record[2], record[3], record[4]
			group, err := gr.FindByName(ctx, src)
			if err != nil {
				return errors.Wrapf(err, "not found group: %s", src)
			}
			hosts, err := hr.ListByGroupID(ctx, group.ID)
			if err != nil {
				return errors.Wrapf(err, "list hosts fail: %d", group.ID)
			}
			fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | | | | |\n", srcZone, "", src, dst, "", "")
			for _, host := range hosts {
				fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | | | | |\n", srcZone, src, host.Hostname, "", host.Ipaddr, port)
			}
		}

		fmt.Fprintln(stdout, b.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(markdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
