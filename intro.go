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

	fmt.Printf("%v \n", persons)
	fmt.Printf("DB = %v \n", persons["DB"])

	//change value
	persons["DB"] = person{name: "Joseph Smith"}

	value, ok := persons["DB"]
	fmt.Printf("DB = %v, ok = %v \n", value, ok)

	delete(persons, "DB")
	fmt.Printf("DB = %v, ok = %v \n", value, ok)
	fmt.Printf("DB = %v \n", persons["DB"])
	fmt.Printf("%v \n", persons)

}
