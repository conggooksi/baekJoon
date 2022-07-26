package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

// 내가 한거
func main() {
	var n int
	
	fmt.Fscanln(reader, &n)
	
	var nums int
	var sum int
	
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums)

		sum += nums
	}
	fmt.Println(sum)
}

// 정답
func sum(a[]int) int {
	var r int
	for _, val := range a{
		r += val
	}
	return r
}