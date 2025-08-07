package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/iterator"
)

func main() {
	fmt.Println("Array based Iterator")
	library := &iterator.Library{}
	library.AddBook(iterator.Book{Title: "The Go Programming Language", Author: "Donovan"})
	library.AddBook(iterator.Book{Title: "Clean Code", Author: "Martin"})
	library.AddBook(iterator.Book{Title: "Design Patterns", Author: "Gamma"})
	iter := library.CreateIterator()
	for iter.HasNext() {
		book := iter.Next()
		fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
	}
	fmt.Println()
	fmt.Println("Linked list based Iterator")
	libraryNode := &iterator.LibraryNode{}
	libraryNode.AddBook(iterator.Book{Title: "The Go Programming Language", Author: "Donovan"})
	libraryNode.AddBook(iterator.Book{Title: "Clean Code", Author: "Martin"})
	libraryNode.AddBook(iterator.Book{Title: "Design Patterns", Author: "Gamma"})
	iter = libraryNode.CreateIterator()
	for iter.HasNext() {
		book := iter.Next()
		fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
	}
}
