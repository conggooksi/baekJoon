package test1978

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test1978() {

	var n, m int
	fmt.Fscan(reader, &n)

	defer writer.Flush()

	check := make(map[int]bool)
	var result []int

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m)

		check[m] = true

		for j := 2; j < m; j++ {
			if m%j == 0 {
				check[m] = false
			}
		}

		if check[m] == true && m > 1 {
			result = append(result, m)
		}
	}

	fmt.Fprintln(writer, len(result))

}
