package test10872

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test10872() {
	defer writer.Flush()
	var n int
	var fac int = 1
	fmt.Fscan(reader, &n)

	for i := 1; i <= n; i++ {
		fac *= i
	}
	fmt.Fprint(writer, fac)

}
