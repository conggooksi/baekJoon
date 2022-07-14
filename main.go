package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	nString := strings.TrimSpace(str)
	nList := strings.TrimSpace(str2)

	blank := strings.Fields(nList)

	n, _ := strconv.Atoi(nString)

	var intList []int
	var num int

	// sort 패키지 활용해서 최대 최소 구하기
	for i := 0; i < n; i++ {
		num, _ = strconv.Atoi(blank[i])
		intList = append(intList, num)
	}
	sort.Ints(intList)
	fmt.Print(intList[0], intList[n-1])
	// 최대 최소 for문으로 직접 만든거
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < i; j++ {
	// 		if blank[i] < blank[j] {
	// 			blank[j], blank[i] = blank[i], blank[j]
	// 		}
	// 	}
	// }

}
