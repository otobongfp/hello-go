package main

import "fmt"

type author struct {
	name, book string
}

func main() {
	me := author{
		"Otobong",
		"Book of Life",
	}

	fmt.Printf("%#v \n", me)
}
