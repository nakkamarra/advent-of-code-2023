package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
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
	lines := make([][]byte, 0, 100000)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't read line: %s", err)
		}
		lines = append(lines, line)
	}
	sum := processLines(lines)
	fmt.Fprintf(os.Stdout, "sum: %d\n", sum)
	os.Exit(0)
}

func processLines(lines [][]byte) int {
	sum := 0
	for i, line := range lines {
		for j, b := range line {
			if isSymbol(b) {
				sum += checkNeighbor(lines, i-1, j-1)
				sum += checkNeighbor(lines, i-1, j)
				sum += checkNeighbor(lines, i-1, j+1)
				sum += checkNeighbor(lines, i, j-1)
				sum += checkNeighbor(lines, i, j+1)
				sum += checkNeighbor(lines, i+1, j-1)
				sum += checkNeighbor(lines, i+1, j)
				sum += checkNeighbor(lines, i+1, j+1)
			}
		}
	}
	return sum
}

func checkNeighbor(lines [][]byte, i int, j int) int {
	if i < 0 || i > len(lines)-1 {
		return 0
	}
	if j < 0 || j > len(lines[i])-1 {
		return 0
	}
	char := rune(lines[i][j])
	if unicode.IsDigit(char) {
		num := composeNumber(lines, i, j)
		converted, err := strconv.Atoi(num)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't convert string to number: %s\n", num)
			return 0
		}
		return converted
	}
	return 0
}

func isSymbol(b byte) bool {
	char := rune(b)
	return b != '.' && !unicode.IsDigit(char) && !unicode.IsSpace(char)
}

func composeNumber(lines [][]byte, i int, j int) string {
	fmt.Fprintf(os.Stdout, "composing number from %s", lines[i])

	number := ""
	byte := lines[i][j]
	number += string(byte)
	lines[i][j] = '.'
	for k := j + 1; k < len(lines[i])-1; k++ {
		char := lines[i][k]
		if unicode.IsDigit(rune(char)) {
			number += string(char)
			lines[i][k] = '.'
		} else {
			break
		}
	}
	for k := j - 1; k >= 0; k-- {
		char := lines[i][k]
		if unicode.IsDigit(rune(char)) {
			number = string(char) + number
			lines[i][k] = '.'
		} else {
			break
		}
	}

	fmt.Fprintf(os.Stdout, "composed number: %s\n", number)
	return number
}
