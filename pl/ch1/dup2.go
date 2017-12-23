package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	counts := make(map[string]int)
	filePaths := os.Args[1:]
	if len(filePaths) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filePath := range filePaths {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v \n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for key, value := range counts {
		fmt.Printf("%d:%s", value, key)
		fmt.Println()
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if f == os.Stdin && input.Text() == "$" {
			break
		}
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(f.Name())
		}
	}
}
