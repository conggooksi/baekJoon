package test2738

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2738() {
	var n, m int

	defer writer.Flush()

	fmt.Fscan(reader, &n, &m)

	a := make([][]int, n)
	b := make([][]int, n)

	for i := range a {
		a[i] = make([]int, m)
	}

	for i := range a {
		b[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var num int
			fmt.Fscan(reader, &num)
			a[i][j] = num
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var num int
			fmt.Fscan(reader, &num)
			b[i][j] = num
		}
	}

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}

	for i := range result {
		for j := range result[i] {
			fmt.Fprintf(writer, "%d ", result[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}
