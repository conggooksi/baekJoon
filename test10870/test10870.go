package test10870

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Test10870() {
	var n int
	var fibo []int = []int{0, 1}
	var result int
	fmt.Fscan(reader, &n)

	for i := 2; i <= n; i++ {
		if n == 0 {
			result = fibo[0]
		} else if n == 1 {
			result = fibo[1]
		} else if n > 1 {
			fibo = append(fibo, (fibo[i-2] + fibo[i-1]))
		}
	}
	result = fibo[n]
	fmt.Print(result)
}
