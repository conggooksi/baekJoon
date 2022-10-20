package test11653

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test11653() {
	var n int
	fmt.Fscan(reader, &n)

	defer writer.Flush()

	var divideNum int
	for true {
		if n == 0 || n == 1 {
			break
		}
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				fmt.Fprintln(writer, i)
				divideNum = i
				break
			}

		}
		n = n / divideNum
	}
}
