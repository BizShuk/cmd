/*
Copyright © 2022 Shuk

*/
package main

import "github.com/bizshuk/cmd/cleanfilename/cmd"

func main() {
	cmd.Execute()

	// filepath.Walk("//User/shuk/Downloads/HAV-default/", func(path string, info fs.FileInfo, err error) error {
	// 	fmt.Println(path)
	// 	return nil
	// })
}
