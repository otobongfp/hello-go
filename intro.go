package main

import "fmt"

func parseOddsEvens(nums []int) (odds []int, evens []int) {

	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return

}

func main() {

	odds, evens := parseOddsEvens([]int{2, 2, 34, 23, 45, 65, 67, 89, 87, 65, 45, 11, 12, 21, 22, 44, 35, 76, 0, 5})

	fmt.Println(odds)
	fmt.Println(evens)

}
