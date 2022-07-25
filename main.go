package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var n int

	fmt.Fscanln(reader, &n)

	// var sum int
	
	var nums = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums)
	}
		// sum += nums[i]
	fmt.Println(nums[1])

	
	// fmt.Fprint(writer, sum)
}