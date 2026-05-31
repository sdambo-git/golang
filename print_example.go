package main

import "fmt"

func main() {
	fmt.Println("Books in stock:")
	var book = "'Master and Commander' , by Patrick O'Brain"
	fmt.Println(book)
	book = "'The Count of Monte Cristo' , by Alexander Dumas"
	fmt.Println(book)

	var author = "Stephen King"
	var title = "'mr Mercedes 📚'"
	printBook(title, author)

	author = "J.K. Rowling"
	title = "'Harry Potter and the Philosopher's Stone 📖'"
	printBook(title, author)

	const name, age = "kim", 22
	fmt.Printf("%s is %d years old. and the name type is %T and age type is %T \n\n", name, age, name, age)
	fmt.Println(`the big bear is done
	by wood and the big bear is a book `)
}

func printBook(title, author string) {
	fmt.Println(title, "by", author)
}
