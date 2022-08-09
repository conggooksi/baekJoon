package test25304

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test25304() {
	defer writer.Flush()
	var total int
	var num int

	fmt.Fscanln(reader, &total)
	fmt.Fscanln(reader, &num)

	var sum int
	for i := 0; i < num; i++ {
		var price, n int
		fmt.Fscan(reader, &price, &n)

		sum += (price * n)
	}

	if sum == total {
		fmt.Fprint(writer, "Yes")
	} else {
		fmt.Fprint(writer, "No")
	}
}
