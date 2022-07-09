package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DigitalRoot(493193))

}

func DigitalRoot(n int) int {

	words := strings.Split(strconv.Itoa(n), "")

	if len(words) == 1 {
		return n
	}

	for len(words) > 1 {
		sum := 0

		for _, value := range words {
			s, _ := strconv.Atoi(value)
			sum += s
		}

		words = strings.Split(strconv.Itoa(sum), "")
	}

	result, _ := strconv.Atoi(words[0])
	return result
}

// func sumSliceString(sl []string) int {
// 	sum := 0
// 	for _, value := range sl {
// 		s, err := strconv.Atoi(value)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		sum += s

// 	}

// 	return sum
// }
