package test2447

import (
	"bufio"
	"fmt"
	"os"
)

func Test2447() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	var n int

	defer writer.Flush()

	fmt.Fscan(reader, &n)

	starMap := make([][]bool, n)

	for i := 0; i < n; i++ {
		starMap[i] = make([]bool, n)
	}

	star(n, 0, 0, starMap)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if starMap[i][j] {
				fmt.Fprint(writer, "*")
			} else {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}

func star(n, x, y int, starMap [][]bool) {
	if n == 3 {
		blank := 0
		for i := x; i < x+n; i++ {
			for j := y; j < y+n; j++ {
				blank++
				if blank == 5 {
					starMap[i][j] = false
				} else {
					starMap[i][j] = true
				}
			}
		}
		return
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			} else {
				star(n/3, x+i*n/3, y+j*n/3, starMap)
			}
		}
	}

}
