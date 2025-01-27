package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Part2()
}

func Part1() {
	scanner := bufio.NewScanner(os.Stdin)

	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

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

func Part2() {
	scanner := bufio.NewScanner(os.Stdin)

	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

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
