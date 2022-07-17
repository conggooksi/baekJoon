package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a int
	var firstList []int
	var secondList []int
	cnt := 0

	for i:=0; i<9; i++ {
		fmt.Fscanln(reader, &a)

		firstList = append(firstList, a)
		secondList = append(secondList, a)

		sort.Ints(firstList)

	}

	for i :=0; i < 9; i++{
		if secondList[i] != firstList[8] {
			cnt += 1
		} else {
			break
		}
	}
	
	fmt.Fprintln(writer, firstList[len(firstList)-1])
	fmt.Fprintln(writer, cnt+1)
}