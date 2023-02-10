package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/mchu1966/cli/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"mode can be",
	"1: to upper case",
	"2: to lower case",
	"3: underscore to upper camel case",
	"4: underscore to lower camel case",
	"5: camel to underscore",
}, "\n")
var str, file string
var mode int8
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format conversion",
	Long:  desc,
	Run:   WordCmdRun,
	Args:  cobra.NoArgs,
}

func init() {
	SetWordCmdFlags(wordCmd)
}

func SetWordCmdFlags(cmd *cobra.Command) {
	cmd.Flags().Int8VarP(&mode, "mode", "m", 0, "word conversion mode")
	cmd.Flags().StringVarP(&str, "str", "s", "", "the target word")
	cmd.Flags().StringVarP(&file, "file", "f", "", "the target file")
	cmd.MarkFlagRequired("mode")
	cmd.MarkFlagFilename("file", "txt")
	cmd.MarkFlagsMutuallyExclusive("str", "file")
}

func WordCmdRun(cmd *cobra.Command, args []string) {
	var content string
	if str != "" {
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			cmd.PrintErrln("mode is not supported, word --help for usage")
			return
		}
		cmd.Println(content)
	} else if file != "" {
		readFile, err := os.Open(file)
		if err != nil {
			cmd.PrintErr(err)
		}
		defer readFile.Close()

		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			line := fileScanner.Text()
			switch mode {
			case ModeUpper:
				line = word.ToUpper(line)
			case ModeLower:
				line = word.ToLower(line)
			case ModeUnderscoreToUpperCamelCase:
				line = word.UnderscoreToUpperCamelCase(line)
			case ModeUnderscoreToLowerCamelCase:
				line = word.UnderscoreToLowerCamelCase(line)
			case ModeCamelCaseToUnderscore:
				line = word.CamelCaseToUnderscore(line)
			default:
				cmd.PrintErrln("mode is not supported, word --help for usage")
				return
			}
			cmd.Println(line)
		}
	}
}
