package test2751

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2751() {

	var n, m int
	var numbers []int
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &m)

		numbers = append(numbers, m)
	}

	sort.Ints(numbers)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, numbers[i])
	}
}
