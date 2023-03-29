package main

type author struct {
	name string
}

type book struct {
	title string
	author
}

type library map[string][]book

func (a book) addBook() {
	a.title = title
	a.author = author
}

func (lib library) addBook(b book) {
	lib[b.author.name] = append(lib[b.author.name], b)
}

func (lib library) lookupByAuthorName(name string) []book {
	return lib[name]
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
