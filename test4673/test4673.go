package test4673

import "fmt"

func Test4673() {
	nums := num(10000)

	for i := 0; i < len(nums); i++ {
		if nums[i] == true {
			fmt.Println(i)
		}
	}
}

func num(n int) (chkNum map[int]bool) {
	chkNum = make(map[int]bool, n+1)

	for i := 0; i < n+1; i++ {
		chkNum[i] = true
	}

	for i := 0; i < n+1; i++ {
		var selfNum = i
		var number = i

		for j := number; j != 0; j /= 10 {
			selfNum += j % 10
		}

		if selfNum <= n {
			chkNum[selfNum] = false
		}
	}

	return chkNum
}
