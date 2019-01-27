package main

import (
	"fmt"
	"os"

	"github.com/joshbohde/lab"
	"github.com/spf13/cobra"
)

var mergeRequestService = lab.MergeRequestService{
	Git:    gitService,
	Gitlab: gitlabService,
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

	mergeRequestCmd.Flags().StringVar(&createMergeRequestOptions.Title, "title", "", "The title of the merge request.")
	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.Description, "description", "d", "", "The description of the merge request. If not provided, will open an editor.")
	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.SourceBranch, "source", "s", "", "The source branch. If not provided, will use your local branch.")
	mergeRequestCmd.Flags().StringVarP(&createMergeRequestOptions.TargetBranch, "target", "t", "", "The target branch. If not provided, will use the project default.")

	mergeRequestCmd.Flags().BoolVarP(&createMergeRequestOptions.KeepSource, "keep-source", "k", false, "Keep source branch after merging.")
}
