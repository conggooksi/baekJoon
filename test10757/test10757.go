package test10757

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func Test10757() {
	defer writer.Flush()
	var n, m big.Int

	fmt.Fscan(reader, &n, &m)

	sum := new(big.Int)
	sum = sum.Add(&n, &m)
	fmt.Println(sum)
}
