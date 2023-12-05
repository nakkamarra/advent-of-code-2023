package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const filename = "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		break
	}
	os.Exit(0)
}
