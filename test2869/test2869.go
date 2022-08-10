package test2869

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2869() {

	defer writer.Flush()

	var up int
	var down int
	var tree int

	fmt.Fscanln(reader, &up, &down, &tree)

	var day = 1

	if (tree-up)%(up-down) == 0 {
		day += (tree - up) / (up - down)
	} else {
		day += (tree-up)/(up-down) + 1
	}

	fmt.Print(day)
}
