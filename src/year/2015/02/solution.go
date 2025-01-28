package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	part1()
	part2()
}

func part1() {
	wrapping := 0
	for _, line := range strings.Split(input, "\n") {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)

		abc := []int{l, w, h}
		slices.Sort(abc)

		wrapping += 2*l*w + 2*w*h + 2*h*l + abc[0]*abc[1]
	}
	fmt.Println(wrapping)
}

func part2() {
	ribbon := 0
	for _, line := range strings.Split(input, "\n") {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)

		abc := []int{l, w, h}
		slices.Sort(abc)

		ribbon += abc[0]*2 + abc[1]*2 + l*w*h
	}
	fmt.Println(ribbon)
}
