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
	fmt.Println("part 1:", findHash(5))
	fmt.Println("part 2:", findHash(6))
}

func findHash(zeroCount int) int {
	zeroes := strings.Repeat("0", zeroCount)
	i := 0
	for {
		i++
		key := input + fmt.Sprint(i)
		hashBytes := md5.Sum([]byte(key))
		hashStr := hex.EncodeToString(hashBytes[:])
		if hashStr[0:zeroCount] == zeroes {
			break
		}
	}
	return i
}
