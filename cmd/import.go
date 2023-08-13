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
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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

  $ cat <<EOL | sg-viz import --host
  foo-host01,192.168.0.1
  foo-host02,192.168.0.2
  bar-vserver,10.0.0.1
  EOL

  $ cat <<EOL | sg-viz import --host
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

  $ cat <<EOL | sg-viz import --group
  group,hostname
  foo,foo-host01
  foo,foo-host02
  bar.example.com,bar-vserver
  EOL

and The POLICY CSV format should look like this

  src,dst,port
  foo,bar.example.com,8080

For example:

  $ cat <<EOL | sg-viz import --policy
  src,dst,port,desc
  foo,bar.example.com,8080,'foo to bar.example.com:8080'
  EOL

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(isHost || isGroup || isPolicy) {
			fmt.Fprintln(cmd.ErrOrStderr(), "at least one of host|group|policy is required")
			os.Exit(1)
		}

		var err error
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
	isHost   bool
	isGroup  bool
	isPolicy bool
)

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().BoolVarP(&isHost, "host", "", false, "given csv should be HOST format")
	importCmd.Flags().BoolVarP(&isGroup, "group", "", false, "given csv should be GROUP format")
	importCmd.Flags().BoolVarP(&isPolicy, "policy", "", false, "given csv should be POLICY format")
	importCmd.MarkFlagsMutuallyExclusive("host", "group", "policy")
}

func importHost(cmd *cobra.Command, _ []string) error {
	stdout := cmd.OutOrStdout()
	fmt.Fprintln(stdout, "<ctrl+d> to submit, <ctrl+c> to abort.")

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return errors.Wrap(err, "read fail")
	}

	// 첫번째 줄이 header 인지 판단.
	content := string(b)
	content = strings.Trim(content, "\n")
	lines := strings.Split(content, "\n")
	re := regexp.MustCompile(`hostname,ipaddr(,desc)?`)
	if re.MatchString(lines[0]) {
		lines = lines[1:]
	}
	fmt.Println(lines)
	return nil
}

func importGroup(_ *cobra.Command, _ []string) error {
	return nil
}

func importPolicy(_ *cobra.Command, _ []string) error {
	return nil
}
