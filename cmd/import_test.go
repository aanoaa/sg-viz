package cmd_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	sgviz "github.com/aanoaa/sgviz/cmd"
)

func TestImport(t *testing.T) {
	cmd := sgviz.NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"import", "--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
