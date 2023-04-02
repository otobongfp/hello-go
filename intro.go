package main

import (
	"fmt"
	"runtime"
)

func main() {

	switch os := runtime.GOOS; os {
	case "linux", "unix":
		fmt.Printf("%s \n", os)
	case "windows":
		fmt.Printf("%s \n", os)
	default:
		fmt.Printf("Something is wrong")
	}

}
