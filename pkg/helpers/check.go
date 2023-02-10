package helpers

import "github.com/spf13/cobra"

func Check(cmd *cobra.Command, err error) {
	if err != nil {
		cmd.PrintErr(err)
	}
}
