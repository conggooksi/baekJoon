package test1157

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test1157() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	text := strings.Fields(str)

	sent := text[0]

	asc := []byte(sent)

	for i := 0; i < len(asc); i++ {
		if int(asc[i]) > 90 {
			asc[i] -= 32
		}
	}

	chk := make(map[int]int)

	for i := 0; i < len(asc); i++ {
		chk[int(asc[i])] += 1
	}

	maxVal := 0
	var maxKey string

	for key, val := range chk {
		if val > maxVal {
			maxVal = val
			maxKey = string(key)
		} else if val == maxVal {
			maxKey = "?"
		}
	}

	fmt.Println(maxKey)
}
