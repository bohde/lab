package main

import (
	"fmt"
	"os"

	"github.com/joshbohde/lab"
	"github.com/spf13/cobra"
)

var authService = lab.AuthService{
	Git: gitService,
}

// mergeRequestCmd represents the mergeRequest command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate on Gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		remote, err := authService.RemoteProject()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("Authenticated with %s\n", remote.Host)
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
