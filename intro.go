package main

import "fmt"

func parseOddsEvens(num []int) (odds []int, evens []int) {
	for _, v := range num {
		if even := v%2 == 0; even {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return
}

func main() {

	odds, evens := parseOddsEvens([]int{11, 25, 67, 78, 98, 22, 46, 21, 12, 2, 44, 43, 66, 44, 19})

	fmt.Println(odds)
	fmt.Println(evens)

}
