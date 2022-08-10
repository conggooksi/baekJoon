package test1712

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test1712() {
	defer writer.Flush()

	var fixedCost int
	var variableCost int
	var price int

	fmt.Fscan(reader, &fixedCost, &variableCost, &price)

	if variableCost >= price {
		fmt.Fprint(writer, -1)
	} else {
		// var totalCost int
		// var revenue int

		fmt.Print((fixedCost / (price - variableCost)) + 1)
	}

}
