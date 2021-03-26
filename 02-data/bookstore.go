package bookstore

import (
	"math/rand"
)

type Book struct {
	Title       string
	Author      string
	Description string
	PriceCents  float64
	ID          int
}

var Books = []Book{}

func GetAllBooks() []Book {
	book1 := Book{Title: "Book1"}
	book2 := Book{Title: "Book2"}
	return []Book{book1, book2}
}

func GetNewID() int {
	return rand.Int()
}
