package cmd

import (
	"fmt"
	"os"

	"github.com/veritone/tagger/tagger"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const taggerVersion = "1.2.1"

var (
	logLevel  string
	taggerCfg = tagger.Config{}
	rootCmd   = &cobra.Command{
		Use:     "tagger",
		Short:   "Tag Git, Github & Docker with all the same semver",
		Example: "tagger 0.1.0",
		Run:     runRoot,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureLogging() {
	if level, err := log.ParseLevel(logLevel); err != nil {
		log.Error("log-level argument malformed: ", logLevel, ": ", err)
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Set log level (debug, info, warn, error, fatal)")

	rootCmd.Flags().StringVarP(&taggerCfg.TaggerFile, "file", "f", "./tagger.yml", "Tagger configuration file")
	rootCmd.Flags().IntVarP(&taggerCfg.Concurrency, "concurrency", "c", 1, "Number of concurrent tagging")
}

func runRoot(cmd *cobra.Command, args []string) {
	configureLogging()

	if len(args) != 1 {
		log.Error("Wrong number of args, run: tagger --help")
		os.Exit(1)
	}

	arg := args[0]
	if arg == "version" {
		fmt.Println(taggerVersion)
		return
	}

	taggerCfg.GlobalTag = arg
	tagger.Tag(taggerCfg)
}
