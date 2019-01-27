package main

import (
	"fmt"
	"os"

	git "github.com/joshbohde/lab/git"
	"github.com/joshbohde/lab/gitlab"
	"github.com/joshbohde/lab/message"
	"github.com/spf13/cobra"
)

var gitService = git.New()
var gitlabService = gitlab.New()
var messageService = message.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lab",
	Short: "Command line interface for Gitlab",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
