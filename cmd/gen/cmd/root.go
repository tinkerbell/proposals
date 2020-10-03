package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use: "gen",
	Long: `gen - The tinkerbell proposal site generator
gen is a command line tool used to generate the tinkerbell proposals
webiste.  Tinkebell has a mechanism based on a Git repository
https://github.com/tinkerbell/proposals where community memebers can
open a proposal to share ideas or features they want to implement or
see implemented. The proposal has various phases and it is useful to
have a html page that summarize all of them. This tool generates
that list starting from the proposals repository.`,
}

func Execute(log *zap.Logger) {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
