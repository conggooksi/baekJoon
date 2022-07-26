package main

import "fmt"

func main() {

	fmt.Println(d(39))
}

func d(n int) int {
	num := n + n/10 + n%10

	return num
}