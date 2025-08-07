package iterator

// 1. Array based iterator

// Book represents an item in the collection
type Book struct {
	Title  string
	Author string
}

type Library struct {
	Books []Book
}

type BookIterator interface {
	HasNext() bool
	Next() *Book
}

type libraryIterator struct {
	books []Book
	index int
}

func (l *libraryIterator) HasNext() bool {
	return l.index < len(l.books)
}

func (l *libraryIterator) Next() *Book {
	if l.HasNext() {
		book := &l.books[l.index]
		l.index++
		return book
	}
	return nil
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) CreateIterator() BookIterator {
	return NewBookIterator(l)
}

func NewBookIterator(library *Library) BookIterator {
	return &libraryIterator{
		books: library.Books,
		index: 0,
	}
}

// 2. Linked list based iterator

type BookNode struct {
	Book *Book
	Next *BookNode
}

type LibraryNode struct {
	Head *BookNode
}

type libraryLinkedListIterator struct {
	current *BookNode
}

func (l *libraryLinkedListIterator) HasNext() bool {
	return l.current != nil
}

func (l *libraryLinkedListIterator) Next() *Book {
	if !l.HasNext() {
		return nil
	}
	book := l.current.Book
	l.current = l.current.Next
	return book
}

func (l *LibraryNode) AddBook(book Book) {
	newNode := &BookNode{Book: &book}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LibraryNode) CreateIterator() BookIterator {
	return NewLibraryLinkedListIterator(l)
}
func NewLibraryLinkedListIterator(library *LibraryNode) BookIterator {
	return &libraryLinkedListIterator{
		current: library.Head,
	}
}
