package test2108

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2108() {
	defer writer.Flush()
	var n, sum, modeCnt, modeResult int
	var min = 4000
	var max = -4000
	var counts = make([]int, 8001)
	var modes []int

	fmt.Fscanln(reader, &n)
	var numList = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &numList[i])

		sum += numList[i]

		idx := numList[i] + 4000

		counts[idx]++

		if modeCnt < counts[idx] {
			modes = []int{numList[i]}
			modeCnt = counts[idx]
		} else if modeCnt == counts[idx] {
			modes = append(modes, numList[i])
		}

		if min > numList[i] {
			min = numList[i]
		}

		if max < numList[i] {
			max = numList[i]
		}

	}
	avgResult := int(math.Round(float64(sum) / float64(n)))

	sort.Ints(numList)
	sort.Ints(modes)

	centerResult := numList[int(math.Ceil(float64(n/2)))]

	if len(modes) > 1 {
		modeResult = modes[1]
	} else {
		modeResult = modes[0]
	}

	rangeResult := max - min

	fmt.Fprintf(writer, "%d\n%d\n%d\n%d\n", avgResult, centerResult, modeResult, rangeResult)
}
