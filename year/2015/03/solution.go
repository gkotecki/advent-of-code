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
	part1()
	part2()
}

type Coord struct {
	x, y int
}

func part1() {
	x, y := 0, 0
	visited := make(map[Coord]bool)

	for _, c := range input {
		switch string(c) {
		case "^":
			y++
		case "v":
			y--
		case ">":
			x++
		case "<":
			x--
		}
		visited[Coord{x, y}] = true
	}

	fmt.Println("part 1:", len(visited))
}

func part2() {
	var santa, roboSanta Coord
	visited := make(map[Coord]bool)

	for i, c := range input {
		var coord *Coord
		if i%2 == 0 {
			coord = &santa
		} else {
			coord = &roboSanta
		}

		switch string(c) {
		case "^":
			coord.y++
		case "v":
			coord.y--
		case ">":
			coord.x++
		case "<":
			coord.x--
		}
		visited[*coord] = true
	}

	fmt.Println("part 2:", len(visited))
}
