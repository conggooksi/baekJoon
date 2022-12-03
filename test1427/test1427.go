package test1427

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func test1427() {
	defer writer.Flush()
	var n, result string

	fmt.Fscanln(reader, &n)

	splitedStr := strings.Split(n, "")
	sort.Sort(sort.Reverse(sort.StringSlice(splitedStr)))

	for i := 0; i < len(n); i++ {
		result += splitedStr[i]
	}

	fmt.Fprintln(writer, result)

}
