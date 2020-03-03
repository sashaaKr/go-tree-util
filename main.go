package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	delimeter            = "───"
	tabCharacter         = "\t"
	postfixForNonLast    = "|"
	prefixNameForNonLast = "├"
	prefixNameForLast    = "└"
)

func createName(isLast bool, file os.FileInfo, level int) string {
	line := []string{}
	if isLast {
		line = append(line, prefixNameForLast)
	} else {
		line = append(line, prefixNameForNonLast)
	}

	line = append(line, delimeter)
	line = append(line, file.Name())
	if !file.IsDir() {
		size := strconv.FormatInt(file.Size(), 10)
		line = append(line, " ("+size+"b)")
	}

	return strings.Join(line, "")
}

func getFiles(dir string, showFiles bool) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	n := 0
	for _, file := range files {
		if showFiles {
			files[n] = file
			n++
		} else if file.IsDir() {
			files[n] = file
			n++
		}
	}
	files = files[:n]

	return files
}

func printTree(dir string, level int, linePrefix string, showFiles bool) {
	files := getFiles(dir, showFiles)
	filesCount := len(files) - 1

	for i := 0; i <= filesCount; i++ {
		file := files[i]
		isLast := i == filesCount

		isDir := file.IsDir()

		fmt.Println(linePrefix + createName(isLast, file, level))

		if isDir {
			var nextLinePrefix string
			if isLast {
				nextLinePrefix = linePrefix
			} else {
				nextLinePrefix = linePrefix + "|"
			}
			nextLinePrefix = nextLinePrefix + tabCharacter
			printTree(filepath.Join(dir, file.Name()), level+1, nextLinePrefix, showFiles)
		}
	}
}

func main() {
	dir := os.Args[1:2][0]
	showFiles := false
	if len(os.Args) > 2 {
		showFiles = os.Args[2:3][0] == "-f"
	}

	printTree(dir, 0, "", showFiles)
}
