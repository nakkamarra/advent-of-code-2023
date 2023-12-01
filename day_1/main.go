package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file: %s\n", filename)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	sum := 0
	for {
		bytes, err := reader.ReadBytes('\n')
		sum += processLine(bytes)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Encountered error reading bytes: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stdout, "Sum: %d\n", sum)
}

// processLine will return an integer representation of the first and last
// numbers encountered in the slice of bytes.
func processLine(input []byte) int {
	first := ""
	last := ""
	asString := string(input)
	for i := range asString {
		if asString[i] > 47 && asString[i] < 58 { // is a digit
			if first == "" {
				first = string(asString[i])
			} else {
				last = string(asString[i])
			}
		}
	}
	if first != "" && last == "" {
		last = first
	}
	parsed, _ := strconv.Atoi(first + last)
	return parsed
}
