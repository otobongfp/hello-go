package main

import "fmt"

func main() {

	num := []int{2, 45, 67, 34, 23, 12, 21, 34, 56}

	for i, n := range num {
		fmt.Println(i, n)
	}

}
