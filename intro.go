package main

import "fmt"

var val = "global"

func changeVal() {
	fmt.Println("global val = ", val)
}

func updateVal() {
	val = "updated global"
	fmt.Println("global val = ", val)
}

func main() {

	val := 42
	//a = *val
	fmt.Printf("%T, local val = %v \n", val, val)
	changeVal()
	updateVal()
	fmt.Printf("%T, local val = %v \n", &val, &val)
}
