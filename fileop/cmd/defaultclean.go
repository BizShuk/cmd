/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	model "github.com/bizshuk/cmd/fileop/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var root *model.Node = &model.Node{Char: rune(32), Next: map[rune]*model.Node{}}

var defaultcleanCmd = &cobra.Command{
	Use:   "defaultclean",
	Short: "Clean unnecessary files",
	Long: `
File list will be removed:
    - .DS_Store
    `,
	Run: func(cmd *cobra.Command, args []string) {
		GenerateTries()
		absPath, err := os.Getwd()
		if err != nil {
			log.Fatal("Failed to get current folder")
		}

		err = filepath.Walk(absPath, func(absPath string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if err != nil {
				return err
			}

			found := FindFileInTries(path.Base(absPath))
			if !found {
				return nil
			}

			err = os.Remove(absPath)
			if err != nil {
				log.Fatal(err)
			}
			log.Info("Remove file:", absPath, " Successfully")

			log.Info("File Name:", info.Name())        // Base name of the file
			log.Info("Size:", info.Size())             // Length in bytes for regular files
			log.Info("Permissions:", info.Mode())      // File mode bits
			log.Info("Last Modified:", info.ModTime()) // Last modification time
			log.Info("Is Directory: ", info.IsDir())   // Abbreviation for Mode().IsDir()

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(defaultcleanCmd)
}

func GetFileList() []string {
	return []string{
		".DS_Store",
	}
}

func GenerateTries() {
	files := GetFileList()

	for _, filename := range files {
		curr := root
		for _, char := range filename {
			if _, ok := curr.Next[char]; !ok {
				curr.Next[char] = &model.Node{Char: char, Next: map[rune]*model.Node{}}
			}
			curr = curr.Next[char]
		}
		curr.End = true
	}
}

func FindFileInTries(filename string) bool {
	curr := root
	for _, char := range filename {

		if _, ok := curr.Next[char]; ok {
			curr = curr.Next[char]
			continue
		}

		return false
	}
	return curr.End
}
