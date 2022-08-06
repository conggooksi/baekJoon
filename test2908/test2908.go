package test2908

import (
	"bufio"
	"fmt"
	"os"
)

func Test2908() {
	var reader = bufio.NewReader(os.Stdin)

	var n1, n2 int

	fmt.Fscan(reader, &n1, &n2)

	changedNum1 := (n1 % 10 * 100) + (n1%100 - n1%10) + n1/100
	changedNum2 := (n2 % 10 * 100) + (n2%100 - n2%10) + n2/100

	if changedNum1 > changedNum2 {
		fmt.Print(changedNum1)
	} else {
		fmt.Print(changedNum2)
	}
}
