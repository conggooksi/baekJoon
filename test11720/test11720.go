package test11720

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test11720() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')
	str2, _ := br.ReadString('\n')

	eraseLine := strings.Fields(str)
	eraseLine2 := strings.Fields(str2)

	intN, _ := strconv.Atoi(eraseLine[0])

	var toSum int
	var sum int
	for i := 0; i < intN; i++ {
		splitN := strings.Split(eraseLine2[0], "")

		toSum, _ = strconv.Atoi(splitN[i])
		sum += toSum
	}
	fmt.Println(sum)

}
