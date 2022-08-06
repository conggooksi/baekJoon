package test10809

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test10809() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	erasedLine := strings.Fields(str)

	text := erasedLine[0]

	chk := make(map[int]int)
	for i := 0; i < 26; i++ {
		chk[i+97] = -1
	}

	for i := 0; i < len(text); i++ {
		converted := int(text[i])
		if chk[converted] == -1 {
			chk[converted] = i
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%d ", chk[i+97])
	}
	fmt.Printf("\n")
}
