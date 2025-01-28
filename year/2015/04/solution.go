package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
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

func part1() {
	i := 0
	for {
		i++
		key := input + fmt.Sprint(i)
		hashBytes := md5.Sum([]byte(key))
		hashStr := hex.EncodeToString(hashBytes[:])
		// fmt.Println(i, hashStr)
		if hashStr[0:5] == "00000" {
			break
		}
	}
	fmt.Println("part 1:", i)
}

func part2() {
	i := 0
	for {
		i++
		key := input + fmt.Sprint(i)
		hashBytes := md5.Sum([]byte(key))
		hashStr := hex.EncodeToString(hashBytes[:])
		// if hashStr[0:5] == "00000" {
		// 	fmt.Println(i, hashStr)
		// }
		if hashStr[0:6] == "000000" {
			break
		}
	}
	fmt.Println("part 2:", i)
}
