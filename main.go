package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(d(i))
	}


	var a = []int64{1,2,3,4,5}
	var b []int64

	b = append(a[:2], a[2+1:]...),

	// for i := 1; i <= 10; i++ {
	// 	a = append(a, i)
	// }

	// for i := 0; i <= len(a); i++ {
	// 	for j := 1; j <= 10; j++ {
	// 		if i == d(j) {
	// 			b = append(a[:i], a[i+1:]...)
	// 		}
	// 	}

	// }
	fmt.Printf("%v\n", b)

}

func d(n int) int {
	var num int	
	num = n
	for true{
		remain := n%10
	
		num = num + remain
		n = n/10

		if n == 0 {
			break
		}
	}

	return num
}