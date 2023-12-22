package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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
	lines := make([]string, 0, 100000)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't read line: %s\n", err)
		}
		lines = append(lines, string(line))
	}
	sum := processLines(lines)
	fmt.Fprintf(os.Stdout, "sum: %d\n", sum)
	os.Exit(0)
}

type Card struct {
	ID            string
	TargetNumbers map[int]bool
	OurNumbers    []int
}

func processLines(lines []string) int {
	sum := 0
	for i := range lines {
		line := lines[i]
		substrings := strings.Split(line, ":")
		if len(substrings) <= 1 {
			fmt.Fprintf(os.Stderr, "line didn't contain `:` %s", line)
			os.Exit(1)
		}
		strings.TrimSpace(substrings[1])
		game := strings.Split(substrings[1], " | ")
		if len(game) <= 1 {
			fmt.Fprintf(os.Stderr, "line didn't contain `|` %s", line)
			os.Exit(1)
		}
		card := new(Card)
		card.ID = substrings[0]
		numbers := strings.Split(game[1], " ")
		ourNumbers := make([]int, 0, len(numbers))
		for i := range numbers {
			num := strings.TrimSpace(numbers[i])
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Fprintf(os.Stderr, "couldn't parse as int %s\n", num)
			} else {
				ourNumbers = append(ourNumbers, val)
			}
		}
		card.OurNumbers = ourNumbers
		targets := make(map[int]bool, len(game[0]))
		winningNums := strings.Split(game[0], " ")
		for _, num := range winningNums {
			val, err := strconv.Atoi(num)
			if err == nil {
				targets[val] = true
			}
		}
		card.TargetNumbers = targets
		sum += calculatePoints(card)
	}
	return sum
}

func calculatePoints(c *Card) int {
	numberOfMatches := 0
	for i := range c.OurNumbers {
		num := c.OurNumbers[i]
		_, ok := c.TargetNumbers[num]
		if ok {
			numberOfMatches++
		}
	}
	if numberOfMatches == 0 {
		return 0
	}
	return int(math.Pow(2.0, float64(numberOfMatches)-1))
}
