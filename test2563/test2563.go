package test2563

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2563() {
	var n int

	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	drawingPaper := make([][]bool, 101)

	for i := range drawingPaper {
		drawingPaper[i] = make([]bool, 101)
	}

	// confetti : 색종이

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)

		for j := x; j < x+10; j++ {
			for k := y; k < y+10; k++ {
				drawingPaper[j][k] = true

				if drawingPaper[j][k] == true {
					continue
				}
			}
		}
	}

	cnt := 0
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			if drawingPaper[i][j] == true {
				cnt += 1
			}
		}
	}
	fmt.Fprintln(writer, cnt)
}
