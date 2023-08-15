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
	"fmt"
	"strings"

	// import sqlite3.
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	sqlite "github.com/aanoaa/sg-viz/internal/db"
	"github.com/aanoaa/sg-viz/repo"
)

// graphCmd represents the graph command.
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Print policies as dot format.",
	Long: `Print policies as dot format.
For example:

  $ sg-viz graph
  digraph {
      "foo" -> "bar.example.com" [label="8080"]
  }

  $ sg-viz graph | dot -Tsvg > example.svg
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		db, err = sqlite.Conn("./sg.db")
		if err != nil {
			return errors.Wrap(err, "db conn fail")
		}

		ctx := context.Background()
		pr := repo.NewPolicyRepo(db)
		rows, err := pr.ListGroup(ctx)
		if err != nil {
			return errors.Wrap(err, "list fail")
		}

		b := new(strings.Builder)
		fmt.Fprintln(b, "digraph {")
		for _, row := range rows {
			fmt.Fprintf(b, "  \"%s\" -> \"%s\" \n", row.Src, row.Dst)
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
