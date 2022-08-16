package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var n int

	fmt.Fscan(reader, &n)

	star(n)
}

func star(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == n/3 && j == n/3 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}

// 2 5 8 11 14 17 20 23

// 1013 16 30 31
