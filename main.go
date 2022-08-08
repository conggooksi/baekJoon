package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var text string
	fmt.Fscan(reader, &text)

	fmt.Println(text)

	fmt.Println(strings.Contains(text, "lj"))
}
