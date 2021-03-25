package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      "jim Bob",
		Description: "some book ill never read",
		PriceCents:  1199,
	}
}
