/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show command version",
	Long:  `Show commnad version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("fileop version: %v\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
