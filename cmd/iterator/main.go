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

	// USE CASE ONE USING ARRAY BASED ITERATOR
	fmt.Println()
	arrayPlaylist := &iterator.PlayList{}
	arrayPlaylist.AddSong(iterator.Song{Name: "This is the life", Artist: "Amy MacDonald"})
	arrayPlaylist.AddSong(iterator.Song{Name: "Untouchable", Artist: "Anathema"})
	arrayPlaylist.AddSong(iterator.Song{Name: "The blind lead the blind", Artist: "Antimatter"})

	fmt.Println("Array-Based Playlist:")
	arrayIterator := arrayPlaylist.CreateIterator()
	for arrayIterator.HasNext() {
		song := arrayIterator.Next()
		fmt.Printf("Song: %s by %s\n", song.Name, song.Artist)
	}
}
