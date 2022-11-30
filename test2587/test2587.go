package test2587

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2587() {

	var numbers []int
	var n, sum int
	defer writer.Flush()

	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &n)

		numbers = append(numbers, n)
	}

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	sort.Ints(numbers)

	fmt.Fprintln(writer, sum/5, numbers[2])
}
