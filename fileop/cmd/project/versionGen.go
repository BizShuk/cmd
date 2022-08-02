/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

/* bash
cd $(git rev-parse --show-toplevel), might failed if no git cmd or not git proejct
versions=($(git tag --points-at HEAD))
versions+=($(git log --pretty=format:'%h' -n 1))
echo -n "${versions[*]}" > version
*/
var versionGenCmd = &cobra.Command{
	Use:   "versionGen",
	Short: "Use git tags/commit as version to generate version file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// If git not found

		//
	},
}

func init() {
	projectCmd.AddCommand(versionGenCmd)
}
