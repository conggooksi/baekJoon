package test7568

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test7568() {
	var n int
	fmt.Fscan(reader, &n)

	defer writer.Flush()
	var x, y int
	var nameList = make([][]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x, &y)
		nameList[i] = make([]int, 2)

		nameList[i][0] = x
		nameList[i][1] = y
	}

	var cnt int
	for i := 0; i < n; i++ {
		cnt = 1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if (nameList[i][0] < nameList[j][0]) && (nameList[i][1] < nameList[j][1]) {
				cnt++
			} else {
				continue
			}
		}
		fmt.Fprintf(writer, "%d ", cnt)
	}
}
