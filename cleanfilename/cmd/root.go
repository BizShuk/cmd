/*
Copyright © 2022 Shuk

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

var (
	DoRecur bool
)

var rootCmd = &cobra.Command{
	Use:   "cleanfilename",
	Short: "Clean prefix tags on filename. ex: [xxx](yyy)filename.ext => filename.ext",
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

		if !DoRecur {
			continue
		}

		if file.IsDir() {
			cleanFiles(fullFileName)
		}

		newFileName := removePrefixTags(file.Name())
		newFullFileName := fmt.Sprintf("%s/%s", path, newFileName)
		fmt.Println(newFullFileName)
		err := os.Rename(fullFileName, newFullFileName)
		if err != nil {
			log.Fatal("rename failed on ", file.Name(), " to ", newFileName, " error: ", err)
		}
	}
}

var r *regexp.Regexp

func removePrefixTags(s string) string {
	cn := r.ReplaceAllString(s, "")
	return cn
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// regex
	r = regexp.MustCompile(`^([\[\(【].*?[\]\)】])*[ \t]*`)
	// logrus format
	log.SetFormatter(&log.TextFormatter{DisableLevelTruncation: true})
	// flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().BoolVarP(&DoRecur, "recursive", "r", false, "Do clean up recursively")

	rootCmd.Flags().BoolP("yes", "y", false, "Confirm for execution")
	_ = rootCmd.MarkFlagRequired("yes")
}
