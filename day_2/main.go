package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	const filename = "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open filename: %s\n", filename)
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
			fmt.Fprintf(os.Stderr, "encountered error reading bytes: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stdout, "sum: %d\n", sum)
	fmt.Fprintf(os.Stdout, "closing...\n")
	os.Exit(0)
}

func processLine(input string) int {
	if input == "" {
		return 0
	}
	fmt.Fprintf(os.Stdout, "processing line: %s", input)
	result := strings.Split(input, ":")
	if len(result) < 2 {
		fmt.Fprintf(os.Stderr, "splitting string on : resulted in less than 2 parts. string: %s\n", input)
		return 0
	}
	// gameTitle := result[0]
	gameData := result[1]
	// if isGameValid(gameData) {
	//	return getGameID(gameTitle)
	// }

	powers := getPowers(gameData)
	fmt.Fprintf(os.Stdout, "got powers: %d, %d, %d\n", powers.Red, powers.Blue, powers.Green)
	return powers.Red * powers.Blue * powers.Green
}

type Bounds struct {
	Red   int
	Green int
	Blue  int
}

func isGameValid(gameData string) bool {
	fmt.Fprintf(os.Stdout, "checking game data: %s", gameData)
	pulls := strings.Split(gameData, ";")
	bounds := Bounds{Red: 12, Green: 13, Blue: 14}
	for _, pull := range pulls {
		p := strings.TrimSpace(pull)
		colors := strings.Split(p, ",")
		for _, color := range colors {
			c := strings.TrimSpace(color)
			parts := strings.Split(c, " ")
			fmt.Fprintf(os.Stdout, "got parts: %s\n", parts)
			name := parts[1]
			value, _ := strconv.Atoi(parts[0])
			if !isColorValid(name, value, bounds) {
				fmt.Fprintf(os.Stderr, "invalid color + count combo: %s %d\n", name, value)
				return false
			}
		}
	}
	return true
}

func isColorValid(name string, value int, upperBounds Bounds) bool {
	switch name {
	case "red":
		return value <= upperBounds.Red
	case "green":
		return value <= upperBounds.Green
	case "blue":
		return value <= upperBounds.Blue
	default:
		return false
	}
}

func getGameID(gameTitle string) int {
	fmt.Fprintf(os.Stdout, "getting game ID: %s\n", gameTitle)
	result := strings.Split(gameTitle, " ")
	if len(result) < 2 {
		fmt.Fprintf(os.Stderr, "splitting string on whitespace resulted in less than 2 parts. string: %s\n", gameTitle)
		return 0
	}
	id, err := strconv.Atoi(result[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't convert string to int: %s\n", result[1])
	}
	return id
}

func getPowers(gameData string) Bounds {
	fmt.Fprintf(os.Stdout, "getting powers for game data: %s", gameData)
	pulls := strings.Split(gameData, ";")
	bounds := Bounds{Red: 0, Green: 0, Blue: 0}
	for _, pull := range pulls {
		p := strings.TrimSpace(pull)
		colors := strings.Split(p, ",")
		for _, color := range colors {
			c := strings.TrimSpace(color)
			parts := strings.Split(c, " ")
			fmt.Fprintf(os.Stdout, "got parts: %s\n", parts)
			name := parts[1]
			value, _ := strconv.Atoi(parts[0])
			checkValue(name, value, &bounds)
		}
	}
	return bounds
}

func checkValue(name string, value int, upperBounds *Bounds) {
	switch name {
	case "red":
		if value > upperBounds.Red {
			upperBounds.Red = value
		}
	case "green":
		if value > upperBounds.Green {
			upperBounds.Green = value
		}
	case "blue":
		if value > upperBounds.Blue {
			upperBounds.Blue = value
		}
	default:
		return
	}
}
