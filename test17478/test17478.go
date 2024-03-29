package test17478

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test17478() {
	var n int

	fmt.Fscan(reader, &n)

	text1 := "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다."

	fmt.Println(text1)

	paragraph2(n)

	paragraph(n, n)

}

func paragraph2(n int) (result string) {
	prefix := "____"

	text2 := `"재귀함수가 뭔가요?"`

	text3 := `"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.`
	text3_1 := `마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.`
	text3_2 := `그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어."`

	text4 := `"재귀함수는 자기 자신을 호출하는 함수라네"`
	text5 := "라고 답변하였지."

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(prefix)
		}
		fmt.Println(text2)
		for j := 0; j < i; j++ {
			fmt.Print(prefix)
		}
		fmt.Println(text3)
		for j := 0; j < i; j++ {
			fmt.Print(prefix)
		}
		fmt.Println(text3_1)
		for j := 0; j < i; j++ {
			fmt.Print(prefix)
		}
		fmt.Println(text3_2)
	}

	for j := n; j > 0; j-- {
		fmt.Print(prefix)
	}
	fmt.Println(text2)

	for j := n; j > 0; j-- {
		fmt.Print(prefix)
	}
	fmt.Println(text4)

	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(prefix)
		}
		fmt.Println(text5)

	}

	fmt.Println(text5)

	return result
}

func paragraph(n int, depth int) {
	headText := "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다."
	middleText := "\"재귀함수가 뭔가요?\""
	bodyText := []string{"\"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.",
		"마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.",
		"그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어.\""}
	bottomText := "\"재귀함수는 자기 자신을 호출하는 함수라네\""
	endText := "라고 답변하였지."
	prefix := "____"
	prefixes := ""

	if n == depth {
		fmt.Println(headText)
	}

	for i := 0; i < n-depth; i++ {
		prefixes += prefix
	}

	fmt.Printf("%s%s\n", prefixes, middleText)

	if depth == 0 {
		fmt.Printf("%s%s\n", prefixes, bottomText)
		fmt.Printf("%s%s\n", prefixes, endText)
		return
	}

	for _, t := range bodyText {
		fmt.Printf("%s%s\n", prefixes, t)
	}

	paragraph(n, depth-1)
	fmt.Printf("%s%s\n", prefixes, endText)
}
