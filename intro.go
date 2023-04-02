package main

import (
	"fmt"
	"runtime"
)

func main() {

	os := runtime.GOOS

	if os == "Linux" || os == "Unix" {
		fmt.Printf("%v \n", os)
	} else if os == "Windows" {
		fmt.Printf("%v \n", os)
	} else {
		fmt.Printf("%s \n", os)
	}

}
