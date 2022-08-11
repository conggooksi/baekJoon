package test10250

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test10250() {
	defer writer.Flush()
	var t int

	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var room []int
		var h, w, n int

		fmt.Fscan(reader, &h, &w, &n)

		for i := 1; i <= w; i++ {
			for j := 1; j <= h; j++ {
				room = append(room, (j*100)+i)
			}
		}
		fmt.Fprintln(writer, room[n-1])
	}

}
