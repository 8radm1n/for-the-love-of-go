package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      "jim Bob",
		Description: "some book ill never read",
		PriceCents:  1199,
		ID:          1234567890,
	}
}

func TestGetAllBooks(t *testing.T) {
	book1 := bookstore.Book{Title: "Book1"}
	book2 := bookstore.Book{Title: "Book2"}
	bookstore.Books = []bookstore.Book{book1, book2}

	want := bookstore.GetAllBooks()
	got := bookstore.Books

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetNewID(t *testing.T) {
	want := 12345678910
	got := bookstore.GetNewID()
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
