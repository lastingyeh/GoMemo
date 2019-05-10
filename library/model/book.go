package model

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrorInsufficient = errors.New("insufficient quantity")
	ErrorNotFoundBook = errors.New("bookItem not found")
)

type Book struct {
	Title     string
	Quantity  int
	Author    string
	CreatedAt time.Time
}

func CreateBook(title string, quantity int, author string, createdAt time.Time) *Book {
	return &Book{title, quantity, author, createdAt}
}

func (b *Book) Borrow(q int) (int, error) {
	if b.Quantity < q {
		return 0, ErrorInsufficient
	}
	b.Quantity -= q
	return q, nil
}

func (b *Book) Return(q int) {
	b.Quantity += q
}

func (b *Book) String() string {
	y, m, d := b.CreatedAt.Date()
	dateStr := fmt.Sprintf("%d/%d/%d", y, m, d)
	return fmt.Sprintf("%s(%d) | %s | %s", b.Title, b.Quantity, b.Author, dateStr)
}
