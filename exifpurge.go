package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type CLIConfig struct {
	targetDirectory string
	ignoreFiles     string
	ignoreDotFiles  bool
}

func parseArgs() CLIConfig {
	cliArgs := CLIConfig{
		targetDirectory: "/home/preocts",
		ignoreFiles:     ".",
		ignoreDotFiles:  false,
	}
	flag.StringVar(&cliArgs.targetDirectory, "target", ".", "Target directory to process (default: current directly).")
	flag.StringVar(&cliArgs.ignoreFiles, "ignore-files", "", "Comma separated list of files to ignore. (default: empty)")
	flag.BoolVar(&cliArgs.ignoreDotFiles, "ignore-dot", false, "When true, dot files are ignored. (default: false)")

	flag.Parse()

	return cliArgs
}

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
	cliArgs := parseArgs()

	ignoreFiles := strings.Split(cliArgs.ignoreFiles, ",")
	for i := 0; i < len(ignoreFiles); i++ {
		ignoreFiles[i] = strings.TrimSpace(ignoreFiles[i])
	}

	fmt.Println("Target Directory: ", cliArgs.targetDirectory)

	files, err := getDirectoryFiles(cliArgs.targetDirectory, ignoreFiles, cliArgs.ignoreDotFiles)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
