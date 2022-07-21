package test4344

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test4344() {
	var tc int

	fmt.Fscanln(reader, &tc)
	defer writer.Flush()

	for i := 0; i < tc; i++ {
		var num int
		fmt.Fscan(reader, &num)
		
		var scores = make([]float64, num)
		var sum, avg float64

		for j := 0; j < num; j++{
			fmt.Fscan(reader, &scores[j])
			sum += scores[j]
			avg = sum / float64(num)
		}
		
		var higherCnt float64
		for j := 0; j < num; j++ {
			if scores[j] > avg {
				higherCnt += 1
			}
		}
		result := higherCnt / float64(num) * 100
		fmt.Fprintf(writer, "%.3f%s\n", result, "%")
	}
}