package test

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Test2480() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))

	diceNum := strings.TrimSpace(str)

	blank := strings.Fields(diceNum)

	for i := 0; i < 3; i++ {
		countDice := strings.Count(diceNum, blank[i])
		if countDice == 3 {
			num, _ := strconv.Atoi(blank[i])
			result := 10000 + num*1000
			fmt.Println(result)
			return
		} else if countDice == 2 {
			num, _ := strconv.Atoi(blank[i])
			result := 1000 + num*100
			fmt.Println(result)
			return
		} else {
			for j := 0; j < 3; j++ {
				num1, _ := strconv.Atoi(blank[i])
				num2, _ := strconv.Atoi(blank[j])
				num3, _ := strconv.Atoi(blank[2])
				num4, _ := strconv.Atoi(blank[1])

				if num3 == num4 {
					continue
				} else {

					if num1 < num2 {
						num1, num2 = num2, num1
						if num1 < num3 {
							continue
						} else {
							result := num1 * 100
							fmt.Println(result)
							return
						}
					}
				}
			}
		}
	}
}

func Test2438() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func Test2439() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		for k := t - i; k > 0; k-- {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func Test10871() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	blank := strings.Fields(str)
	blank2 := strings.Fields(str2)

	var blink1 = []int{}
	var a = []int{}

	for i := 0; i <= len(blank)-1; i++ {
		b1, _ := strconv.Atoi(blank[i])
		blink1 = append(blink1, b1)
	}

	for i := 0; i <= len(blank2)-1; i++ {
		b2, _ := strconv.Atoi(blank2[i])
		a = append(a, b2)
	}

	n := blink1[0]
	x := blink1[1]

	for i := 0; i <= n-1; i++ {
		if x > a[i] {
			fmt.Printf("%d ", a[i])
		}
	}
}

func Test11022() {
	var t int
	fmt.Scan(&t)

	var num1 int
	var num2 int

	for i := 1; i <= t; i++ {
		fmt.Scan(&num1, &num2)
		result := num1 + num2
		fmt.Printf("Case #%d: %d + %d = %d\n", i, num1, num2, result)
	}

}

func Test10818() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	nString := strings.TrimSpace(str)
	nList := strings.TrimSpace(str2)

	blank := strings.Fields(nList)

	n, _ := strconv.Atoi(nString)

	var intList []int
	var num int

	// sort 패키지 활용해서 최대 최소 구하기
	for i := 0; i < n; i++ {
		num, _ = strconv.Atoi(blank[i])
		intList = append(intList, num)
	}
	sort.Ints(intList)
	fmt.Print(intList[0], intList[n-1])
	// 최대 최소 for문으로 직접 만든거
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < i; j++ {
	// 		if blank[i] < blank[j] {
	// 			blank[j], blank[i] = blank[i], blank[j]
	// 		}
	// 	}
	// }
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test2577() {

	defer writer.Flush()

	var a int
	var b int
	var c int

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	multiple := a * b * c

	n := strconv.Itoa(multiple)

	slice := strings.Split(n, "")

	for i:=0; i< 10; i++ {
		var numCnt int
		for j:=0; j< len(slice); j++{
			if slice[j] == strconv.Itoa(i) {
				numCnt += 1
			}
		}
		fmt.Println(numCnt)
	}
}


func Test3052() {

	defer writer.Flush()

	var a int
	var nList []int
	var remainList []int

	for true{
		num, _ := fmt.Fscanln(reader, &a)
		if num != 1 {
			break
		} else {
			nList = append(nList, a)
		}
	}

	for i:=0; i< len(nList); i++ {
		remainList = append(remainList, nList[i]%42)
	}


keys := make(map[int]bool)
var result [] int

for _, value := range remainList{
	if _, saveValue := keys[value]; !saveValue{ 
		keys[value] = true
		result = append(result,value)
	}
} 

	fmt.Fprintln(writer, len(result))
}


func Test1546() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))

	n := strings.TrimSpace(str)
	scores := strings.TrimSpace(str2)

	blank := strings.Fields(scores)

	nums, _ := strconv.Atoi(n)

	var floatList[] float64

	for i:=0; i<nums;i++ {
		floats, _ := strconv.ParseFloat(blank[i], 64)
		floatList = append(floatList, floats)
	}
	for i := 0; i < nums; i++ {
		for j := 0; j < i; j++ {
			if floatList[i] < floatList[j] {
				floatList[j], floatList[i] = floatList[i], floatList[j]
			}
		}
	}

	max := floatList[nums-1]
	var newFloats [] float64
	var result float64
	var mean float64

	for i:=0; i<nums;i++{
		newFloat := (floatList[i] / max) * 100
		newFloats = append(newFloats, newFloat)
		result = result + newFloats[i]
		floatNums, _ := strconv.ParseFloat(n, 64)
		mean = result / floatNums
	}
	fmt.Printf("%f", mean)	
}

