package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("ERROR: Required parameters are missing")
		fmt.Println("Usage: srn <fileSrc> <oldStr> <newStr>")
		os.Exit(1)
	}

	fileSrc := args[0]
	oldStr := args[1]
	newStr := args[2]

	data, err := ioutil.ReadFile(fileSrc)

	errorHandler(err)

	// Searching for the filename (without ext)
	if oldStr == "_" {
		oldStr = strings.Replace(filepath.Base(fileSrc), filepath.Ext(fileSrc), "", -1)
	}

	fileContents := string(data)
	finalString := strings.Replace(fileContents, oldStr, newStr, -1)
	basename := filepath.Dir(fileSrc) + string(os.PathSeparator)
	file := basename + newStr + filepath.Ext(fileSrc)

	_err := ioutil.WriteFile(file, []byte(finalString), 0644)

	errorHandler(_err)

	fmt.Printf("New file created at: %s\n", file)
}

func errorHandler(e error) {
	if e != nil {
		panic(e)
	}
}
