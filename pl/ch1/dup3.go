package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main3() {
	counts := make(map[string]int)

	for _, filePath := range os.Args[1:] {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err)
		} else {
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}

	for key, value := range counts {
		fmt.Printf("%d:%s", value, key)
		fmt.Println()
	}
}
