package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mergeRequestCmd represents the mergeRequest command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display lab version information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\nbuildTime: %s\nbuilder: %s\ngoversion: %s\n",
			version, buildTime, builder, goversion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
