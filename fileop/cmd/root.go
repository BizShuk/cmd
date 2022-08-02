/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Verbose bool = false
)

// RoodCmd represents the base command when called without any subcommands
var RoodCmd = &cobra.Command{
	Use:   "fileop",
	Short: "All file-related op",
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RoodCmd.
func Execute() {
	err := RoodCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
