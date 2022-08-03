package main

import "fmt"

func main() {
	selfNums := selfNum(10000)

	for i := 1; i < len(selfNums); i++ {
		if selfNums[i] == true {
			fmt.Println(i)
		}
	}
}

func selfNum(n int) (chkNum map[int]bool) {

	chkNum = make(map[int]bool, n+1)
	for i := 0; i < n+1; i++ {
		chkNum[i] = true
	}

	for i := 0; i < n+1; i++ {
		var sum = i
		var num = i
		for j := num; j != 0; j /= 10 {
			sum += j % 10
		}
		if sum <= n {
			chkNum[sum] = false
		}
	}

	return chkNum
}
