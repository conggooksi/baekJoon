package test11729

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test11729() {
	var n int
	fmt.Fscan(reader, &n)

	defer writer.Flush()

	fmt.Fprintln(writer, int(math.Pow(2, float64(n))-1))
	move(n, 1, 3)
}

func move(n, from, to int) {
	if n == 1 {
		fmt.Fprintf(writer, "%d %d\n", from, to)
		return
	}
	empty := 6 - from - to
	move(n-1, from, empty)
	fmt.Fprintf(writer, "%d %d\n", from, to)
	move(n-1, empty, to)
}
