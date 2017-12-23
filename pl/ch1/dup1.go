package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "$" {
			break
		}
		counts[input.Text()]++
	}

	for key, value := range counts {
		fmt.Printf("%d:%s", value, key)
		fmt.Println()
	}
}
