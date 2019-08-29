package main

import (
	"fmt"
	"strings"
)

func main() {
	// Case 1
	// ar = [1, 2, 3] => 6
	arr := []int{1, 2, 3}
	fmt.Printf("Case 1: %d\n", sumArray(arr))

	// Case 2
	// Tue 16 Jul 2019 into 2019-07-16
	strDate := "Tue 16 Jul 2019"
	fmt.Printf("Case 2: %s\n", dateConverter(strDate))

	// Case 3
	// if n = 5, we calculate 5 x 4 x 3 x 2 x 1 and get 120
	fmt.Printf("Case 3: %d\n", factorial(5))

	// Case 4
	/**
	5:00 -> five o'clock
	5:01 -> one minute past five
	5:10 -> ten minutes past five
	5:15 -> quarter past five
	5:30 -> half past five
	5:40 -> twenty minutes to six
	5:45 -> quarter to six
	5:47 -> thirteen minutes to six
	5:28 -> twenty eight minutes past five
	**/
	fmt.Printf("Case 4: %s\n", timeConverter(5, 0))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 01))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 10))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 15))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 30))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 40))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 45))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 47))
	fmt.Printf("Case 4: %s\n", timeConverter(5, 28))

	// Case 5
	arrList := []int{3, 3, 2, 1, 3}
	fmt.Printf("Case 5: %d\n", arrayReducer(arrList))
}

func sumArray(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func dateConverter(input string) string {
	var sliceMonth = map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Des": "12",
	}

	fullDate := strings.Split(input, " ")
	strMonth := sliceMonth[fullDate[2]]
	return fullDate[3] + "-" + strMonth + "-" + fullDate[1]
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func timeConverter(h, m int) string {
	var numbers = map[int]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		21: "twenty one",
		22: "twenty two",
		23: "twenty three",
		24: "twenty four",
		25: "twenty five",
		26: "twenty six",
		27: "twenty seven",
		28: "twenty eight",
		29: "twenty nine",
	}

	if m == 0 {
		return fmt.Sprintf("%s o'clock", numbers[h])
	} else if m == 1 {
		return fmt.Sprintf("one minute past %s", numbers[h])
	} else if m == 59 {
		return fmt.Sprintf("one minute to %s", numbers[(h%12)+1])
	} else if m == 15 {
		return fmt.Sprintf("quarter past %s", numbers[h])
	} else if m == 30 {
		return fmt.Sprintf("half past %s", numbers[h])
	} else if m == 45 {
		return fmt.Sprintf("quarter to %s", numbers[(h%12)+1])
	} else if m <= 30 {
		return fmt.Sprintf("%s minutes past %s", numbers[m], numbers[h])
	} else if m > 30 {
		return fmt.Sprintf("%s minutes to %s", numbers[60-m], numbers[(h%12)+1])
	}
	return ""
}

func arrayReducer(arr []int) int {
	var count int
	// var G []int
	for _, i := range arr {
		for _, j := range arr {
			if i == j {
				count = count + 1
			}
		}
		count = 0
	}
	return len(arr)
}
