package test2231

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2231() {
	var n int
	fmt.Fscan(reader, &n)

	defer writer.Flush()
	natural(n)
}

func natural(n int) {
	var ans int

	for i := 1; i <= n; i++ {
		ans = calc(i)
		ans += i

		if ans == n {
			fmt.Fprint(writer, i)
			return
		}

	}
	if ans != n {
		fmt.Fprint(writer, 0)
	}
}

func calc(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}
