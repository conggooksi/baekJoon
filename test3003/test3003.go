package test3003

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test3003() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	white := strings.Fields(str)

	var whiteArray []int

	for i := 0; i < len(white); i++ {
		whiteInt, _ := strconv.Atoi(white[i])

		whiteArray = append(whiteArray, whiteInt)
	}

	fmt.Printf("%d ", 1-int(whiteArray[0]))
	fmt.Printf("%d ", 1-int(whiteArray[1]))
	fmt.Printf("%d ", 2-int(whiteArray[2]))
	fmt.Printf("%d ", 2-int(whiteArray[3]))
	fmt.Printf("%d ", 2-int(whiteArray[4]))
	fmt.Printf("%d ", 8-int(whiteArray[5]))
}
