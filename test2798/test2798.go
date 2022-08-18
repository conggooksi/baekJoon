package test2798

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2798() {
	var n, m int
	var card int
	var cards []int
	fmt.Fscan(reader, &n, &m)

	defer writer.Flush()

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &card)
		cards = append(cards, card)
	}

	plus(n, m, cards)

}

func plus(n, m int, cards []int) {
	var sum int
	var sums []int

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum = cards[i] + cards[j] + cards[k]
				if sum == m {
					fmt.Fprint(writer, sum)
					return
				}
				sums = append(sums, sum)
			}
		}
	}

	sort.Ints(sums)

	for i := len(sums) - 1; i > 0; i-- {
		if sums[i] < m {
			fmt.Fprint(writer, sums[i])
			return
		}
	}

}
