package test2675

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2675() {

	var testNum int

	fmt.Fscanln(reader, &testNum)
	defer writer.Flush()

	for i := 0; i < testNum; i++ {
		var repeat int

		fmt.Fscan(reader, &repeat)

		var text string
		fmt.Fscanln(reader, &text)

		splitText := strings.Split(text, "")

		for i := 0; i < len(splitText); i++ {
			for j := 0; j < repeat; j++ {
				fmt.Fprint(writer, splitText[i])
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
