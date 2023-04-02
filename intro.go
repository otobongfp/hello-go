package main

import "fmt"

func later(result string) {
	fmt.Println(result)
}

func main() {

	//defer

	defer later("first")
	defer later("second")
	defer later("third")

	fmt.Println("Lets Skip to the Good Part")
	fmt.Println("Yes!!!")

	//recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from Panic: ", err)
		}
	}()

	//panics
	panic("Something went wrong!!!")

}
