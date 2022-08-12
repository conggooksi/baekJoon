package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// 내가 했던 건데 틀림 ---------------------------------------
// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
//
// func main() {
// 	defer writer.Flush()
// 	var num1, num2 uint
//
// 	fmt.Fscan(reader, &num1, &num2)
//
// 	result := num1 + num2
//
// 	var s string = strconv.FormatUint(uint64(result), 10)
//
// 	fmt.Fprint(writer, s)
//
// }

// 2번 방법 --------------------------------------------------------
//
// func factoral(n uint64) (r *big.Int) {
//
// 	one, bn := big.NewInt(1), new(big.Int).SetUint64(n)
//
// 	r = big.NewInt(1)
// 	if bn.Cmp(one) <= 0 {
// 		return
// 	}
// 	for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i, one) {
// 		r.Mul(r, i)
// 	}
// 	return
// }
//
// func add(number *big.Int) *big.Int {
// 	ten := big.NewInt(10)
// 	sum := big.NewInt(0)
// 	mod := big.NewInt(0)
// 	for ten.Cmp(number) < 0 {
// 		sum.Add(sum, mod.Mod(number, ten))
// 		number.Div(number, ten)
// 	}
// 	sum.Add(sum, number)
// 	return sum
// }
// func main() {
// 	fmt.Println(add(factoral(100)))
//
// }

// 3번 방법 안되는 듯(숫자는 크게 나오지만 정확하지 않다)-------
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a, b, c big.Int

	fmt.Fscan(reader, &a, &b)

	// add := new(big.Int)
	// add = add.Add(&a, &b)

	fmt.Fprint(writer, c.Add(*a, *b))
}
