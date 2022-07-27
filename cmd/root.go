package cmd

import (
	"github.com/guilhermocc/test-case-generator/internal/generator"
	"os"

	"github.com/spf13/cobra"
)

var inputFilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github.com/guilhermocc/test-case-generator",
	Short: "A brief description of your application",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		generator.GenerateTestCases(inputFilePath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFilePath, "input-file-path", "i", "", "input csv file describing the state machine to be tested")
	rootCmd.MarkFlagRequired("input-file-path")

}
