package test9020

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test9020() {
	var a [10001]bool
	var t, n int

	fmt.Fscanln(reader, &t)

	defer writer.Flush()

	for k := 0; k < t; k++ {
		var result []int
		fmt.Fscanln(reader, &n)

		for i := 2; i <= n; i++ {
			a[i] = true
		}

		for i := 2; i <= n; i++ {
			if a[i] == false {
				continue
			}

			for j := i + i; j <= n; j += i {
				a[j] = false
			}

			if a[i] == true {
				result = append(result, i)
			}
		}

		//var minMinus []int
		var resultI, resultJ int
		for i := 0; i < len(result)-1; i++ {
			for j := 0; j < len(result); j++ {
				if result[i]+result[j] == n {
					if result[i] == result[j] {
						resultI = result[i]
						resultJ = result[j]
						break
					}

					if result[j] > result[i] {
						resultI = result[i]
						resultJ = result[j]
						break
					}
				}
			}
		}

		fmt.Fprintln(writer, resultI, resultJ)
	}
}
