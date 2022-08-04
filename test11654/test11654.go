package test11654

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Test11654() {
	var n string
	fmt.Fscan(reader, &n)

	a := []byte(n)

	fmt.Println(a[0])
}

// func main() {
// 	var input byte
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ = reader.ReadByte()

// 	fmt.Printf("%d\n", input)
// }
