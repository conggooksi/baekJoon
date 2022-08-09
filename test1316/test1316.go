package test1316

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test1316() {
	var num int
	defer writer.Flush()

	fmt.Fscan(reader, &num)
	var text string
	var cnt int

	for i := 0; i < num; i++ {
		fmt.Fscan(reader, &text)

		var chk = make(map[string]int)

		for j := 0; j < len(text); j++ {
			if strings.Contains(text, string(text[j])) == true {
				chk[string(text[j])] += 1
			}
		}

		var groupWord int = 1
		for key, val := range chk {
			var chkGroup []int
			if val > 1 {

				for j := 0; j < len(text); j++ {
					if string(text[j]) == key {
						chkGroup = append(chkGroup, j)
					}
				}

				for k := 0; k < len(chkGroup)-1; k++ {
					if (chkGroup[k+1] - chkGroup[k]) == 1 {
						groupWord = 1
					} else {
						groupWord = 0
						break
					}
				}
			}

			if groupWord == 0 {
				break
			}

			if groupWord != 0 && val == 1 {
				groupWord = 1
			}
		}
		cnt += groupWord
	}
	fmt.Fprintln(writer, cnt)

}
