package main

import (
	"bufio"
	"fmt"
	"os"
)

// Modifique dup2 para que exiba os nomes de todos os arquivos em que cada linha duplicada ocorre.

func main() {
	files := os.Args[1:]
	fileLines := make(map[string]map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, fileLines, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileLines, arg)
			f.Close()
		}

		for line, fileMap := range fileLines {
			totalCount := 0
			for _, count := range fileMap {
				totalCount += count
			}
			if totalCount > 1 {
				fmt.Printf("%d\t%s\n", totalCount, line)
				for file, count := range fileMap {
					fmt.Printf("\t%d\t%s\n", count, file)
				}
			}
		}
	}
}

func countLines(f *os.File, fileLines map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if fileLines[line] == nil {
			fileLines[line] = make(map[string]int)
		}
		fileLines[line][fileName]++
	}
}
