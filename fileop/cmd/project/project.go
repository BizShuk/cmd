/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/bizshuk/cmd/fileop/cmd"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "All related to project",
	Long:  `All related to project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

func init() {
	cmd.RoodCmd.AddCommand(projectCmd)

}
