/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "embed"

	"github.com/bizshuk/cmd/fileop/cmd"
	_ "github.com/bizshuk/cmd/fileop/cmd/project"
)

//go:embed version
var Version string

func main() {
	cmd.Execute(Version)
}
