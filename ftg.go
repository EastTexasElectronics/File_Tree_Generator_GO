package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	excludePatterns = map[string]bool{}
	outputLocation  string
	inputDirectory  string
	version         = "1.0.1"
	author          = "https://github.com/easttexaselectronics"
	repository      = "https://github.com/EastTexasElectronics/File-Tree-Generator-Multiverse/tree/main/Go"
	donation        = "https://www.buymeacoffee.com/rmhavelaar"
)

func init() {
	log.SetFlags(0)
}

func showUsage() {
	fmt.Println(`Usage: ftg [-e pattern1,pattern2,...] [-o output_location] [-d input_directory] [-i] [-c] [-h] [-v]
Options:
  -e, --exclude      Exclude directories or files (comma-separated)(.git,node_modules,.vscode)
  -o, --output       Specify an output location; default output is in the pwd
  -d, --directory    Specify an input directory; default is the pwd
  -i, --interactive  Interactive mode to select items to exclude
  -c, --clear        Clear the exclusion list
  -h, --help         Show this help message and exit
  -v, --version      Show version information and exit`)
	os.Exit(1)
}

func showVersion() {
	fmt.Printf("File Tree Generator version: %s\nLeave us a star at %s\nAuthor: %s\n", version, repository, author)
	fmt.Printf("Buy me a coffee: %s\n", donation)
	os.Exit(0)
}

func errorExit(message string) {
	log.Fatalf("Error: %s\n", message)
}

func shouldExclude(name string) bool {
	return excludePatterns[name]
}

func getEntries(path string) ([]fs.DirEntry, error) {
	return os.ReadDir(path)
}

func printEntry(writer io.Writer, name, entryType, prefix string, isLast bool) {
	var connector string
	if isLast {
		connector = "└──"
	} else {
		connector = "├──"
	}
	if _, err := fmt.Fprintf(writer, "%s%s [%s] %s\n", prefix, connector, entryType, name); err != nil {
		log.Printf("Error writing entry: %v", err)
	}
}

func getEntryType(entry fs.DirEntry) string {
	if entry.IsDir() {
		return "D"
	}
	return "F"
}

func generateTree(writer io.Writer, path string, prefix string, entries []fs.DirEntry) {
	for i, entry := range entries {
		name := entry.Name()
		if shouldExclude(name) {
			continue
		}

		isLast := i == len(entries)-1
		entryType := getEntryType(entry)
		printEntry(writer, name, entryType, prefix, isLast)

		if entryType == "D" {
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			subEntries, err := getEntries(filepath.Join(path, name))
			if err != nil {
				log.Printf("Cannot read directory %s: %v", filepath.Join(path, name), err)
				continue
			}
			generateTree(writer, filepath.Join(path, name), newPrefix, subEntries)
		}
	}
}

func interactiveMode() {
	fmt.Println("Interactive mode not implemented.")
	os.Exit(1)
}

func main() {
	var exclude string
	var interactive, clearExclusions, help, versionFlag bool

	flag.StringVar(&exclude, "e", "", "Exclude directories or files (comma-separated)")
	flag.StringVar(&outputLocation, "o", "", "Specify an output location")
	flag.StringVar(&inputDirectory, "d", "", "Specify an input directory")
	flag.BoolVar(&interactive, "i", false, "Interactive visual mode to select items to exclude")
	flag.BoolVar(&clearExclusions, "c", false, "Clear the exclusion list")
	flag.BoolVar(&help, "h", false, "Show this help message and exit")
	flag.BoolVar(&versionFlag, "v", false, "Show version information and exit")

	flag.Parse()

	switch {
	case help:
		showUsage()
	case versionFlag:
		showVersion()
	case clearExclusions:
		excludePatterns = map[string]bool{}
	}

	if exclude != "" {
		for _, pattern := range strings.Split(exclude, ",") {
			excludePatterns[pattern] = true
		}
	}

	commonExcludes := []string{"node_modules", ".next", ".vscode", ".idea", ".git", "target", "Cargo.lock"}
	for _, pattern := range commonExcludes {
		excludePatterns[pattern] = true
	}

	if interactive {
		interactiveMode()
	}

	if outputLocation == "" {
		currentTime := time.Now().Format("15-04-05")
		outputLocation = fmt.Sprintf("file_tree_%s.md", currentTime)
	}

	if inputDirectory == "" {
		var err error
		inputDirectory, err = os.Getwd()
		if err != nil {
			errorExit("Failed to get current directory")
		}
	}

	fmt.Printf("Generating your file tree for %s, while you wait... \nGive the project a star at %s\n", inputDirectory, repository)

	outputFile, err := os.Create(outputLocation)
	if err != nil {
		errorExit(fmt.Sprintf("Cannot write to output location %s", outputLocation))
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	if _, err := fmt.Fprintf(outputFile, "# File Tree for %s\n\n## Give the project a star at %s\n```sh\n", inputDirectory, repository); err != nil {
		errorExit(fmt.Sprintf("Error writing to output file: %v", err))
	}

	entries, err := getEntries(inputDirectory)
	if err != nil {
		errorExit("Cannot read the input directory")
	}
	generateTree(outputFile, inputDirectory, "", entries)

	if _, err := fmt.Fprintln(outputFile, "```"); err != nil {
		log.Printf("Error writing to output file: %v", err)
	}

	fmt.Printf("File tree has been written to %s\n", outputLocation)
}
