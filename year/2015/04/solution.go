package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	fmt.Println("parallel part 1:", findHashParallel(5))
	fmt.Println("parallel part 2:", findHashParallel(6))
	fmt.Println("parallel part X:", findHashParallel(7))
	fmt.Println("sync part 1:", findHash(5))
	fmt.Println("sync part 2:", findHash(6))
}

func findHash(zeroCount int) int {
	zeroes := strings.Repeat("0", zeroCount)
	for i := 0; ; i++ {
		key := input + fmt.Sprint(i)
		hashBytes := md5.Sum([]byte(key))
		hashStr := hex.EncodeToString(hashBytes[:])
		if hashStr[0:zeroCount] == zeroes {
			return i
		}
	}
}

func findHashParallel(zeroCount int) int {
	zeroes := strings.Repeat("0", zeroCount)

	numCPU := runtime.NumCPU()
	result := make(chan int, numCPU)

	for workerId := 0; workerId < numCPU; workerId++ {
		go func(id int) {
			for i := id; ; i += numCPU {
				key := input + fmt.Sprint(i)
				hashBytes := md5.Sum([]byte(key))
				hashStr := hex.EncodeToString(hashBytes[:])
				if hashStr[0:zeroCount] == zeroes {
					result <- i
					close(result)
					return
				}

				select {
				case <-result:
					return
				default:
				}
			}
		}(workerId)
	}

	answer := <-result
	return answer
}
