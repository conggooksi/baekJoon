package test2292

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Test2292() {
	var num int

	fmt.Fscan(reader, &num)

	var end int = 1
	i := 0
	for {
		end += 6 * i

		if end >= num {
			fmt.Println(i + 1)
			break
		} else {
			i += 1
		}
	}

}
