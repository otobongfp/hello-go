package main

type author struct {
	name string
}

type book struct {
	title string
	author
}

type library struct {
	book
}

func (a book) addBook() {
	a.title = title
	a.author = author
}

func (value library) addBook() {
	value.book.title = title
	value.book.author = author
}

func (keyword library) lookupByAuthorName() string {
	return
}

func main() {

	books = book{
		"Harry Potter",
		author{
			"JK Rowling",
		},
	}

	fmt.println()

}
