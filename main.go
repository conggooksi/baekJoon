package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	eraseLine := strings.Fields(str)

	a := []byte(eraseLine)

	// splitedText := strings.Split(eraseLine[0], "")

}
