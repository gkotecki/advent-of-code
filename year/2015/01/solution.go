package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	part2()
}

func part1() {
	floor := 0
	for _, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println(floor)
}

func part2() {
	floor := 0
	for i, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
