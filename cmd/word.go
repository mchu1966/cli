package cmd

import (
	"log"
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
var str string
var mode int8
var file string
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format conversion",
	Long:  desc,
	Run:   WordCmdRun,
}

func init() {
	SetWordCmdFlags(wordCmd)
}

func SetWordCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&str, "str", "s", "", "the target word")
	// TODO
	// wordCmd.Flags().StringVarP(&file, "file", "f", "", "the target file")
	cmd.Flags().Int8VarP(&mode, "mode", "m", 0, "word conversion mode")
}

func WordCmdRun(cmd *cobra.Command, args []string) {
	var content string
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
		log.Fatalf("mode is not supported, word --help for usage")
	}

	cmd.Println(content)
}
