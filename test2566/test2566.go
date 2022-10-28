package test2566

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2566() {

	defer writer.Flush()

	a := make([][]int, 9)

	for i := range a {
		a[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var num int
			fmt.Fscan(reader, &num)
			a[i][j] = num
		}
	}

	max := 0
	maxI := 0
	maxJ := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if max < a[i][j] {
				max = a[i][j]
				maxI = i
				maxJ = j
			}
		}
	}

	fmt.Fprintf(writer, "%d\n%d %d", max, maxI+1, maxJ+1)
}
