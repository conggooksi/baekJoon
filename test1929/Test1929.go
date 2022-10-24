package test1929

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test1929() {
	var m, n int

	fmt.Fscan(reader, &m, &n)

	defer writer.Flush()

	primeNumberSieve(m, n)

}

func primeNumberSieve(m int, n int) {
	var a [1000001]bool

	// 1. 배열을 생성하여 초기화한다.
	for i := 2; i <= n; i++ {
		a[i] = true
	}

	// 2. 2부터 시작해서 특정 수의 배수에 해당하는 수를 모두 지운다.
	// (지울 때 자기자신은 지우지 않고, 이미 지워진 수는 건너뛴다.)
	for i := 2; i <= n; i++ {
		if a[i] == false {
			continue
		} // 이미 지워진 수라면 건너뛰기

		// 이미 지워진 숫자가 아니라면, 그 배수부터 출발하여, 가능한 모든 숫자 지우기
		for j := i + i; j <= n; j += i {
			a[j] = false
		}
	}

	// 3. 2부터 시작하여 남아있는 수를 모두 출력한다.
	for i := 2; i <= n; i++ {
		if i < m {
			a[i] = false
		}

		if a[i] == true {
			fmt.Fprintf(writer, "%d\n", i)
		}
	}
}
