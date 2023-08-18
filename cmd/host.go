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
	"encoding/csv"
	"os"

	_ "github.com/mattn/go-sqlite3" // import sqlite3.
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"

	sqlite "github.com/aanoaa/sgviz/internal/db"
	"github.com/aanoaa/sgviz/repo"
)

// hostCmd represents the host command.
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Export policies as csv format for host.",
	Long: `Export policies as csv format for host.
For example:

  $ sgviz host
  src,src_addr,dst,dst_addr,port,protocol
  foo-host01,192.168.0.1,bar-vserver,10.0.0.1,8080,tcp
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

		ctx := context.Background()
		pr := repo.NewPolicyRepo(db)
		rows, err := pr.ListHost(ctx)
		if err != nil {
			return errors.Wrap(err, "list fail")
		}

		w := csv.NewWriter(cmd.OutOrStdout())
		for _, row := range rows {
			if err := w.Write(row.ToStrings()); err != nil {
				return errors.Wrap(err, "write fail")
			}
		}

		w.Flush()
		return errors.Wrap(w.Error(), "error writing record to csv")
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
