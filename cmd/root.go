package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var taggerFile string
var rootCmd = &cobra.Command{
	Use:     "tagger",
	Short:   "Tag Git, Github & Docker with all the same semver",
	Example: "tagger 0.1.0",
	Run:     runRoot,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&taggerFile, "file", "f", "./tagger.yml", "Tagger configuration file")
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Println(args)
	fmt.Println(taggerFile)
}
