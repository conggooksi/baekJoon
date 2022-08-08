package test5622

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Test5622() {
	var text string

	fmt.Fscan(reader, &text)

	var time = 0

	for i := 0; i < len(text); i++ {
		if text[i] >= 65 && text[i] <= 67 {
			time += 3
		} else if text[i] >= 68 && text[i] <= 70 {
			time += 4
		} else if text[i] >= 71 && text[i] <= 73 {
			time += 5
		} else if text[i] >= 74 && text[i] <= 76 {
			time += 6
		} else if text[i] >= 77 && text[i] <= 79 {
			time += 7
		} else if text[i] >= 80 && text[i] <= 83 {
			time += 8
		} else if text[i] >= 84 && text[i] <= 86 {
			time += 9
		} else if text[i] >= 87 && text[i] <= 90 {
			time += 10
		} else {
			time += 11
		}
	}
	fmt.Print(time)
}
