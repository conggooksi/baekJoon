package test2581

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2581() {
	var m, n int
	fmt.Fscan(reader, &m, &n)

	defer writer.Flush()

	check := make(map[int]bool)
	var result []int

	for i := m; i <= n; i++ {
		check[i] = true

		if i == 1 {
			check[i] = false
		} else if i == 2 {
			check[i] = true
		}

		for j := 2; j < i; j++ {
			if i%j == 0 {
				check[i] = false
			}
		}

		if check[i] == true {
			result = append(result, i)
		}
	}
	var total int
	for i := range result {
		total += result[i]
	}

	if len(result) == 0 {
		fmt.Fprint(writer, -1)
	} else {
		fmt.Fprintf(writer, "%d\n%d", total, result[0])
	}
}
