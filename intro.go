package main

import "fmt"

type author struct {
	first string
	last  string
}

type person struct {
	author
	penName string
}

func (value author) fullName() string {
	return value.first + " " + value.last
}

func (value person) fullName() string {
	return fmt.Sprintf("%s , (%s)", value.author.fullName(), value.penName)
}

func main() {

	me := person{
		author{
			"Otobong",
			"Peter",
		},
		"Bolt",
	}

	fmt.Println(me.fullName())
	fmt.Printf("%#v \n", me)

}
