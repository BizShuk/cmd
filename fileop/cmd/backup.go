/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CopyCount *int

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a file with versions",
	Long: `Backup a file with versions.
		ex: file.jpg
		=> file.jpg.bak
		=> file.jpg.bak.0
		=> file.jpg.bak.1
		=> file.jpg.bak.2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// check file exist
		fileUri := args[0]
		filepathAbs, err := filepath.Abs(fileUri)

		if err != nil {
			logrus.Fatalln("fails to interpret file paht:", fileUri)
		}

		cleanBackup(filepathAbs)

		if _, err := os.Stat(filepathAbs); os.IsNotExist(err) {
			logrus.Fatalln("file not exists:", fileUri, err)
		}

		backup(filepathAbs, -1)
	},
}

// from original file to file.bak
func backup(file string, idx int) bool {
	bakname := file + ".bak"

	_, err := os.Stat(bakname)
	if !os.IsNotExist(err) { // Warning: os.IsExist(err) not work for os.Stat
		backupSeq(bakname, bakname, idx+1)
	}

	return copyFile(file, bakname)
}

// from bak to bak.0 .1 .2
func backupSeq(ori string, file string, idx int) bool {
	bakname := ori + "." + strconv.Itoa(idx)

	_, err := os.Stat(bakname)
	if !os.IsNotExist(err) {
		backupSeq(ori, bakname, idx+1)
	}

	return copyFile(file, bakname)
}

func copyFile(source string, destination string) bool {
	input, err := os.ReadFile(source)
	if err != nil {
		logrus.Fatalln("Error read file:", source, err)
		return false
	}

	err = os.WriteFile(destination, input, 0644)
	if err != nil {
		logrus.Fatalln("Error creating", destination, err)

		return false
	}
	return true
}

func cleanBackup(file string) {
	dir, filename := filepath.Dir(file), filepath.Base(file)

	entries, err := os.ReadDir(dir)
	if err != nil {
		logrus.Fatalln("Read dir failed", err)
	}

	lastBackupFile := filename + ".bak." + strconv.Itoa(*CopyCount-1)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasPrefix(entry.Name(), filename) && entry.Name() > lastBackupFile {
			err = os.Remove(dir + "/" + entry.Name())
			if err != nil {
				logrus.Fatal(err)
			}
			logrus.Infoln("Remove backup file: ", entry.Name())
		}
	}
}

func init() {
	RootCmd.AddCommand(backupCmd)

	CopyCount = backupCmd.Flags().IntP("count", "c", 5, "Maximum number of copies")

}
