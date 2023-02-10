package cmd_test

import (
	"testing"

	"github.com/mchu1966/cli/cmd"
	"github.com/spf13/cobra"
)

func TestTimeCmd(t *testing.T) {
	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: []string{"--str", "a", "--mode", "1"},
			err:  nil,
			out:  "A",
		},
	}

	root := &cobra.Command{
		Use: "word",
		Run: cmd.NowTimeCmdRun,
	}

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
