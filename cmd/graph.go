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
	"regexp"
	"strings"

	// import sqlite3.
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/aanoaa/sgviz/internal/reader"
)

// graphCmd represents the graph command.
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Print policies as dot format.",
	Long: `Print policies as dot format.
For example:

  $ sgviz graph
  digraph {
      "foo" -> "bar.example.com" [label="8080"]
  }

  $ sgviz graph | dot -Tsvg > example.svg
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		re := regexp.MustCompile(`src,dst,port,protocol\n`)
		records, err := reader.StdinToCsv(re)
		if err != nil {
			return errors.Wrap(err, "stdin to csv fail")
		}

		if len(records) == 0 {
			return errors.New("Empty input")
		}

		if len(records[0]) != 4 {
			return errors.New("Unexpected format")
		}

		b := new(strings.Builder)
		fmt.Fprintln(b, "digraph {")
		for _, record := range records {
			fmt.Fprintf(b, "  \"%s\" -> \"%s\" \n", record[0], record[1])
		}
		fmt.Fprintln(b, "}")
		stdout := cmd.OutOrStdout()
		fmt.Fprintln(stdout, b.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
