package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mchu1966/cli/cmd"
	"github.com/spf13/cobra"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestWordCmd(t *testing.T) {
	tt := []struct {
		args []string
		err  error
		out  string
	}{
		// {
		// 	args: nil,
		// 	err:  errors.New("flag needs an argument"),
		// 	out:  "",
		// },
		{
			args: []string{"--str", "a", "--mode", "1"},
			err:  nil,
			out:  "A",
		},
	}

	root := &cobra.Command{
		Use: "word",
		Run: cmd.WordCmdRun,
	}
	cmd.SetWordCmdFlags(root)

	for _, tc := range tt {
		out, err := execute(t, root, tc.args...)
		if err != tc.err {
			t.Errorf("word cmd error - Expected %s, got:%s", tc.err.Error(), err.Error())
		}
		if out != tc.out {
			t.Errorf("word cmd - Expected %s, got:%s", tc.out, out)
		}
	}
}
