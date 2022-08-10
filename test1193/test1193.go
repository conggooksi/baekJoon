package test1193

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Test1193() {
	var x int

	fmt.Fscanln(reader, &x)

	var step = 1
	var num = 0

	for {

		if num > x {
			break
		}
		step++

		for i := 1; i < step; i++ {
			num++
			if num == x {
				if step%2 == 1 {
					fmt.Printf("%d/%d", i, step-i)
				} else {
					fmt.Printf("%d/%d", step-i, i)
				}
			}
		}
	}
}
