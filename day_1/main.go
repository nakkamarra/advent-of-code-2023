package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		sum += processLine(string(bytes))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Encountered error reading bytes: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stdout, "Sum: %d\n", sum)
}

var words []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}

// processLine will return an integer representation of the first and last
// numbers encountered in the string.
func processLine(input string) int {
	if input == "" {
		fmt.Fprintf(os.Stderr, "processing line: got empty string, skipping...\n")
		return 0
	}
	first := findFirstNumber(input)
	last := findLastNumber(input)
	fmt.Fprintf(os.Stdout, "processing line: %s got first %s last %s\n", strings.Trim(input, "\n"), first, last)
	parsed, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse string to int: %s\n", first+last)
		return 0
	}
	return parsed
}

func findFirstNumber(input string) string {
	word := ""
	for _, char := range input {
		if unicode.IsDigit(char) {
			return string(char)
		} else {
			word += string(char)
			for _, suf := range words {
				if strings.HasSuffix(word, suf) {
					return wordToInt(suf)
				}
			}
		}
	}
	return ""
}

func findLastNumber(input string) string {
	word := ""
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(input[i])) {
			return string(input[i])
		} else {
			temp := string(input[i])
			temp = temp + word
			word = temp
			for _, pre := range words {
				if strings.HasPrefix(word, pre) {
					return wordToInt(pre)
				}
			}
		}
	}
	return ""
}

func wordToInt(input string) string {
	switch input {
	case "zero":
		return "0"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"
	}
}
