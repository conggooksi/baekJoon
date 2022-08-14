package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var n int

	fmt.Fscan(reader, &n)

	text1 := "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다."

	fmt.Println(text1)
	for i :=0; i < n; i++ {
		paragraph2(n)
	}

	
}

func paragraph1(n int) (result string) {
	prefix := "____"
	text2 := `"재귀함수가 뭔가요?"`
	text4 := `"재귀함수는 자기 자신을 호출하는 함수라네"`
	text5 := "라고 답변하였지."

	for i := n; i > 0; i-- {
		fmt.Print(prefix)
	}
	fmt.Println(text2)
	for i := n; i > 0; i-- {
		fmt.Print(prefix)
	}
	fmt.Println(text4)
	for i := n; i > 0; i-- {
		fmt.Print(prefix)
	}
	fmt.Println(text5)

	return result
}

func paragraph2(n int) (result string) {
	prefix := "____"

	text2 := `"재귀함수가 뭔가요?"`

	text3 := `"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.`
	text3_1 := `마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.`
	text3_2 := `그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어."`

	text5 := "라고 답변하였지."

	fmt.Println(text2)
	fmt.Println(text3)
	fmt.Println(text3_1)
	fmt.Println(text3_2)

	for i := 1; i < n; i++ {
		fmt.Print(prefix)
	}
	fmt.Println(text2)
	for i := 1; i < n; i++ {
		fmt.Print(prefix)
	}
	fmt.Println(text3)
	for i := 1; i < n; i++ {
		fmt.Print(prefix)
	}
	fmt.Println(text3_1)
	for i := 1; i < n; i++ {
		fmt.Print(prefix)
	}
	fmt.Println(text3_2)
	paragraph1(n)
	for i := n; i < n; i++ {
		fmt.Print(prefix)
	}
	fmt.Println(text5)

	return result
}
