package test2839

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2839() {
	defer writer.Flush()
	var n int
	var result int = -1

	fmt.Fscan(reader, &n)

	// if n%5 == 0 {
	// 	result = n / 5
	// } else if n%5 != 0 && (n%5)%3 == 0 {
	// 	result = (n/5 + (n%5)/3)
	// } else if n%5 != 0 && n%3 == 0 {
	// 	result = (n / 3)
	// } else {
	// 	for i := 1; i*3 < n; i++ {
	// 		for j := 1; j*5 < n; j++ {
	// 			if (i*3)+(j*5) == n {
	// 				result = i + j
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i*5)+(j*3) == n {
				result = i + j
				break
			}
		}
	}
	fmt.Print(result)

}
