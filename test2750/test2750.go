package test2750

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2750() {
	var n, m int
	var a []int

	defer writer.Flush()

	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &m)
		a = append(a, m)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	for i := range a {
		fmt.Fprintln(writer, a[i])
	}
}
