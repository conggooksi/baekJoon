package test2941

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Test2941() {
	var text string
	fmt.Fscan(reader, &text)

	text = strings.Replace(text, "c=", "!", -1)
	text = strings.Replace(text, "c-", "@", -1)
	text = strings.Replace(text, "dz=", "#", -1)
	text = strings.Replace(text, "d-", "$", -1)
	text = strings.Replace(text, "lj", "%", -1)
	text = strings.Replace(text, "nj", "^", -1)
	text = strings.Replace(text, "s=", "&", -1)
	text = strings.Replace(text, "z=", "*", -1)

	fmt.Println(len(text))
}
