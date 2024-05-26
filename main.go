package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	// Define command-line flags
	inputDir := flag.String("input", "", "Input directory")
	outputFile := flag.String("output", "output.txt", "Output file")
	fileExtension := flag.String("ext", ".txt", "File extension to merge")
	recursive := flag.Bool("recursive", false, "Recursively include files in subdirectories")

	// Parse command-line flags
	flag.Parse()

	// Check if input directory is provided
	if *inputDir == "" {
		fmt.Println("Usage: mergefiles -input <input_directory> -output <output_file> -ext <file_extension> [-recursive]")
		return
	}

	var files []string

	// Walk through the input directory and collect all file paths with the specified extension
	err := filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), *fileExtension) {
			files = append(files, path)
		}
		if !*recursive && info.IsDir() && path != *inputDir {
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
		return
	}

	// Sort the files by name
	sort.Strings(files)

	var outputContent string

	// Read each file and append its content to the output content
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		outputContent += fmt.Sprintf("File: %s\n\n%s\n\n", file, string(content))
	}

	// Write the combined content to the output file
	err = ioutil.WriteFile(*outputFile, []byte(outputContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Println("Files have been successfully merged into", *outputFile)
}
