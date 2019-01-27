package main

import (
	"fmt"
	"os"

	"github.com/joshbohde/lab"
	"github.com/spf13/cobra"
)

var issueService = lab.IssueService{
	Git:     gitService,
	Gitlab:  gitlabService,
	Message: messageService,
}

var createIssueOptions = lab.CreateIssueOptions{}

// mergeRequestCmd represents the mergeRequest command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Open an issue on Gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := issueService.Create(&createIssueOptions)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)

	issueCmd.Flags().StringVarP(&createIssueOptions.Message, "message", "m", "", "The message of the merge request. The first line is the title, the rest is the description.")
	issueCmd.Flags().StringVarP(&createIssueOptions.File, "file", "F", "", "Read the merge request title and description from this file.")
	issueCmd.Flags().BoolVarP(&createIssueOptions.Edit, "edit", "e", false, "Edit provided message.")

}
