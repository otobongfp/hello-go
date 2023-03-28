package main

import "fmt"

type person struct {
	name string
}

func main() {

	persons := map[string]person{
		"OP": {name: "Otobong Peter"},
		"DB": {name: "Dan Brown"},
	}

	fmt.Printf("%#v \n", persons)

}
