package test1065

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test1065() {
	var n int
	defer writer.Flush()

	fmt.Fscan(reader, &n)

	cnt := han(n)
	fmt.Fprintln(writer, cnt)

}

func han(n int) (cnt int) {
	if n < 100 {
		cnt = n
		return
	} else {
		for i := 100; i <= n; i++ {
			one := i % 10
			ten := i / 10 % 10
			hund := i / 100

			if (hund - ten) == (ten - one) {
				cnt += 1
			}
		}
	}
	cnt += 99
	return cnt
}
