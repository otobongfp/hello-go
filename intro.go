package main

import "fmt"

// Author type with name field
type author struct {
	name string
}

// Book with title and author field
type book struct {
	title  string
	author author
}

// A library that maps books
type library map[string][]book

// add books
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

func main() {

	type cash float32
	type farm string

	var money cash = 78
	var crops farm = "Cassava"

	fmt.Println(money)
	fmt.Println(crops)

}
