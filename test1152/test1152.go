package test1152

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test1152() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	text := strings.Fields(str)

	// splitedText := strings.Split(text[1], " ")

	fmt.Println(len(text))
}
