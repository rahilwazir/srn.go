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
		fmt.Println("String replace and new file")
		fmt.Println("ERROR: Required parameters are missing")
		fmt.Println("Usage: srn <file> <old_str> <new_str>")
		os.Exit(1)
	}

	fileSrc := args[0]
	old_str := args[1]
	new_str := args[2]

	data, err := ioutil.ReadFile(fileSrc)

	errorHandler(err)

	fileContents := string(data)
	finalString := strings.Replace(fileContents, old_str, new_str, -1)
	basename := filepath.Dir(fileSrc) + string(os.PathSeparator)
	file := basename + new_str + filepath.Ext(fileSrc)

	_err := ioutil.WriteFile(file, []byte(finalString), 0644)

	errorHandler(_err)

	fmt.Printf("New file created at: %s\n", file)
}

func errorHandler(e error) {
	if e != nil {
		panic(e)
	}
}
