package test10989

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test10989() {
	defer writer.Flush()
	var n, m int

	fmt.Fscanln(reader, &n)

	var counts = make([]int, 10001)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &m)
		counts[m]++ // 카운팅정렬
	}

	for i := 1; i < 10001; i++ {
		for j := 0; j < counts[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}
