package main

import "fmt"

func main() {

	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}

	var result []int
	var check bool

	for _, valueArr := range arr {
		if result == nil {
			result = append(result, valueArr)
		}

		for _, resultValue := range result {
			if valueArr == resultValue {
				check = false
				break

			} else {
				check = true

			}
		}
		if check {
			result = append(result, valueArr)
		}
	}

	fmt.Println(result)

}
