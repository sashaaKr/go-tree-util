package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	delimeter            = "───"
	tabCharacter         = "\t"
	postfixForNonLast    = "│"
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
		line = append(line, " ")
		sizeStr := "(" + size + "b)"
		if size == "0" {
			sizeStr = "(empty)"
		}
		line = append(line, sizeStr)
	}

	return strings.Join(line, "")
}

func getFiles(dir string, showFiles bool) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
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

	return files, nil
}

func printTree(
	dir string,
	level int,
	linePrefix string,
	showFiles bool,
	out io.Writer) error {
	files, err := getFiles(dir, showFiles)
	if err != nil {
		return err
	}

	filesCount := len(files) - 1

	for i := 0; i <= filesCount; i++ {
		file := files[i]
		isLast := i == filesCount

		isDir := file.IsDir()

		fmt.Fprintln(out, linePrefix+createName(isLast, file, level))

		if isDir {
			var nextLinePrefix string
			if isLast {
				nextLinePrefix = linePrefix
			} else {
				nextLinePrefix = linePrefix + postfixForNonLast
			}
			nextLinePrefix = nextLinePrefix + tabCharacter
			printTree(
				filepath.Join(dir, file.Name()),
				level+1,
				nextLinePrefix,
				showFiles,
				out)
		}
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err := printTree(path, 0, "", printFiles, out)
	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
