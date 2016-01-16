package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	counts       = make(map[string]int)
	containFiles = make(map[string]map[string]bool)
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin)
		return
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for name, _ := range containFiles[line] {
				fmt.Printf("%s\t", name)
			}

			fmt.Println("")
		}
	}
}

func countLines(f *os.File) {
	input := bufio.NewScanner(f)
	name := f.Name()

	for input.Scan() {
		line := input.Text()
		counts[line]++

		if containFiles[line] == nil {
			containFiles[line] = make(map[string]bool)
		}

		if !containFiles[line][name] {
			containFiles[line][name] = true
		}
	}
}
