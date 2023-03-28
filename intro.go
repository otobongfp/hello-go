package main

import "fmt"

func main() {
	var ages = make([]int, 0)

	fmt.Println(ages)

	ages = append(ages, 5)
	ages = append(ages, 4)
	ages = append(ages, 12)
	ages = append(ages, 8)

	fmt.Println(ages)
	fmt.Println(ages[1:6])
}
