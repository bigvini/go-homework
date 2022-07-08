package main

import "fmt"

func main() {

	a := []int{1, 2, 1, 3, 1}
	b := []int{1, 2}
	d := []int{}

	d = ArrayDiff(a, b)

	fmt.Println(d)

}

func ArrayDiff(a, b []int) []int {

	aa := a
	bb := b

	for _, bValue := range bb {
		d := deleteOneNumberSlice(aa, bValue)
		aa = d
	}

	return aa

}

func deleteOneNumberSlice(list []int, num int) []int {
	d := []int{}
	for _, value := range list {
		if num != value {
			d = append(d, value)
		}
	}
	return d
}
