/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var DoRecur bool = false

// prefixtagCmd represents the prefixtag command
var prefixtagCmd = &cobra.Command{
	Use:   "prefixtag",
	Short: "",
	Long:  ``,
	Run:   cleanFileName,
}

func cleanFileName(cmd *cobra.Command, args []string) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current folder")
	}

	cleanFiles(path)
	log.Info("Rename all files successfully")
}

func cleanFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Failed to read this folder")
	}

	for _, file := range files {
		fullFileName := fmt.Sprintf("%s/%s", path, file.Name())

		if file.IsDir() && !DoRecur {
			continue
		}

		if file.IsDir() {
			cleanFiles(fullFileName)
		}

		newFileName := removePrefixTags(file.Name())
		newFullFileName := fmt.Sprintf("%s/%s", path, newFileName)

		if file.Name() == newFileName {
			continue
		}

		err := os.Rename(fullFileName, newFullFileName)
		log.Info("Remove prefix tags Successully on ", fullFileName)

		if err != nil {
			log.Fatal("rename failed on ", file.Name(), " to ", newFileName, " error: ", err)
		}
	}
}

func removePrefixTags(s string) string {
	cn := GetRegexPattern().ReplaceAllString(s, "")
	return cn
}

func init() {
	renameCmd.AddCommand(prefixtagCmd)

	// logrus format
	log.SetFormatter(&log.TextFormatter{DisableLevelTruncation: true})
	// flags

	prefixtagCmd.Flags().BoolVarP(&DoRecur, "recursive", "r", false, "Do clean up recursively")

	prefixtagCmd.Flags().BoolP("confirm", "y", false, "Confirm for execution")
	_ = prefixtagCmd.MarkFlagRequired("confirm")
}

func GetRegexPattern() *regexp.Regexp {
	return regexp.MustCompile(`^([\[\(【].*?[\]\)】])*[ \t]*`)
}
