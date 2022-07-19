package test8958

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test8958() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	n := strings.TrimSpace(str)
	
	nums, _ := strconv.Atoi(n)
	
	
	for i:=0; i<nums; i++{
		defer writer.Flush()
		var str2 string
		
		fmt.Fscanln(reader, &str2)
		quiz := strings.TrimSpace(str2)
		quizes := strings.Split(quiz,"")
		
		sum := 0
		score := 0

		for j:=0; j<len(quiz);j++{
			if quizes[j] == "O" {
				if score == 0 {
					score = 1
					sum += score
				} else {
					score +=1
					sum += score
				}
			} else {
				score = 0
			}
		}
		fmt.Fprintln(writer, sum)
	}
}