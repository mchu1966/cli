package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	root := &cobra.Command{
		Use: "root",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	err := root.Execute()
	if err != nil {
		t.Errorf("root cmd failed, err:%v", err)
	}
}
