package test5597

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test5597() {
	var n int
	check := make(map[int]bool)

	defer writer.Flush()

	for i := 1; i <= 30; i++ {
		check[i] = false
	}

	for i := 0; i < 28; i++ {
		fmt.Fscanln(reader, &n)
		check[n] = true
	}

	for i := 1; i <= 30; i++ {
		if check[i] == false {
			fmt.Fprintln(writer, i)
		}
	}

}
