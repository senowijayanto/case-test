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
	fmt.Printf("Case 3: %d\n")
}

func sumArray(num []int) int {
	var sum int

	for i := 0; i < len(num); i++ {
		sum += num[i]
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

func factorial([]int) int {}
