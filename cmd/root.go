package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	CmdAddCommand(rootCmd, []*cobra.Command{wordCmd, timeCmd})
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func CmdAddCommand(parent *cobra.Command, cmd []*cobra.Command) {
	parent.AddCommand(cmd...)
}
