package main

// The IterableCollection interface defines the createIterator
// function, which returns an iterator object
type IterableCollection interface{
	createIterator() iterator
}

// The iterator interface contains the hasNext and next functions
// which allow the collection to return items as nedded
type iterator interface {
	hasNext() bool
	next() *Book
}

// TODO: BookIterator is a concrete iterator for a Book collection
type BookIterator struct{
	current int
	books []Book
}

func (b *BookIterator) hasNext() bool{
	// TODO: implement hasNext()
	if b.current < len(b.books){
		return true
	}
	return false
}

func (b *BookIterator) next() *Book{
	// TODO: implement next()
	if b.hasNext(){
		bk := b.books[b.current]
		b.current++
		return &bk
	}
	return nil
}