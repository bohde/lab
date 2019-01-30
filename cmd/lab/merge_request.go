package main

import (
	"fmt"
	"os"

	"github.com/joshbohde/lab"
	"github.com/spf13/cobra"
)

var mergeRequestService = lab.MergeRequestService{
	Git:     gitService,
	Gitlab:  gitlabService,
	Message: messageService,
	Writer:  os.Stdout,
}

var createMergeRequestOptions = lab.CreateMergeRequestOptions{}

// mergeRequestCmd represents the mergeRequest command
var mergeRequestCmd = &cobra.Command{
	Use:   "merge-request",
	Short: "Open a merge request on Gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := mergeRequestService.Create(&createMergeRequestOptions)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(mergeRequestCmd)

	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.Message, "message", "m", "", "The message of the merge request. The first line is the title, the rest is the description.")
	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.File, "file", "F", "", "Read the merge request title and description from this file.")
	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.SourceBranch, "source", "s", "", "The source branch. If not provided, will use your local branch.")
	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.TargetBranch, "target", "t", "", "The target branch. If not provided, will use the project default.")

	mergeRequestCmd.Flags().BoolVarP(&createMergeRequestOptions.KeepSource, "keep-source", "k", false, "Keep source branch after merging.")

	mergeRequestCmd.Flags().BoolVarP(&createMergeRequestOptions.Edit, "edit", "e", false, "Edit provided message.")

}
