package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"

	var result string
	numberInput := []int32{}

	stroka := strings.Split(input, " ")
	var max int32 = 0
	var min int32 = 0

	for _, valueInput := range stroka {
		num, err := strconv.Atoi(valueInput)
		if err == nil {
			numberInput = append(numberInput, int32(num))
		}
	}

	for _, value := range numberInput {
		if max == 0 {
			max = value
		}

		if value > max {
			max = value
		}

		if min == 0 {
			min = value
		}
		if value < min {
			min = value
		}

	}

	if max == min {
		result = strconv.Itoa(int(max))
	} else {
		result = strconv.Itoa(int(max))
		result += " "
		result += strconv.Itoa(int(min))

	}

	fmt.Println(result)

}
