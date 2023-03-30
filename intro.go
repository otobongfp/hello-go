package main

import "fmt"

func main() {

	num := []int{2, 45, 67, 34, 23, 12, 21, 34, 56}

	for i, n := range num {
		fmt.Println(i, n)
	}

	magicWords := map[int]string{1: "please", 2: "excuse me", 3: "thank you", 4: "sorry"}

	fmt.Println("Magic Words")
	for k, v := range magicWords {
		fmt.Println(k, v)
	}
}
