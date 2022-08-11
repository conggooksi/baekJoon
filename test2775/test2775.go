package test2775

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2775() {
	defer writer.Flush()
	var t int

	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var k, n int

		fmt.Fscan(reader, &k)
		fmt.Fscan(reader, &n)

		fmt.Fprintln(writer, people(k, n))
	}
}

func people(k, n int) (person int) {
	if k == 1 {
		for i := 1; i <= n; i++ {
			person += i
		}
		return person
	}
	for i := 1; i <= n; i++ {
		person += people(k-1, i)

	}

	return person
}

// 5층 1명 1+6명(7) 1+6+21명(28) 1+6+21+56명(84)
// 4층 1명 1+5명(6) 1+5+15명(21) 1+5+15+35명(56)
// 3층 1명 1+4명(5) 1+4+10명(15) 1+4+10+20명(35)
// 2층 1명 1+3명(4) 1+3+6명(10)  1+3+6+10명 (20)
// 1층 1명 1+2명(3) 1+2+3명(6)   1+2+3+4명  (10)
// 0층 1명 2명      3명          4
