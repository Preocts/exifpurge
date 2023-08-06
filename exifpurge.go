package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

// TODO: Bring these in from a CLI interface
const TargetDirectory string = "/home/preocts"
const IgnoreDotFiles bool = true
const IgnoreFiles string = "foobar,.profile"

// Return a slice of files from the targetDirectory respecting defined exclusions and any error that occurred
func getDirectoryFiles(targetDirectory string, ignoreFiles []string, ignoreDotFiles bool) ([]string, error) {

	returnFiles := make([]string, 0)

	files, err := os.ReadDir(targetDirectory)

	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if ignoreDotFiles && strings.HasPrefix(file.Name(), ".") {
				continue
			}

			if slices.Contains(ignoreFiles, file.Name()) {
				continue
			}

			returnFiles = append(returnFiles, file.Name())
		}
	}

	return returnFiles, err
}

func main() {
	ignoreFiles := strings.Split(IgnoreFiles, ",")
	fmt.Println("Target Directory: ", TargetDirectory)

	files, err := getDirectoryFiles(TargetDirectory, ignoreFiles, IgnoreDotFiles)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
