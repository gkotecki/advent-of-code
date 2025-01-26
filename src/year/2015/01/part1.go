package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Your solution code here
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
