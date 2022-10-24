package test10807

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test10807() {
	var n, v, a int
	var check []int

	fmt.Fscan(reader, &n)

	defer writer.Flush()

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		check = append(check, a)
	}

	fmt.Fscan(reader, &v)

	var cnt int
	for i := 0; i < len(check); i++ {
		if check[i] == v {
			cnt += 1
		}
	}

	fmt.Fprint(writer, cnt)
}
