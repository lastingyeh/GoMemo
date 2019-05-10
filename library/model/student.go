package model

import (
	"bytes"
	"fmt"
	"time"
)

type Student struct {
	Name, Grade, Id, Sex string
	Books []*BookItem
}

type BookItem struct {
	Book   *Book
	Amount int
}

var books []*Book

func init() {
	books = []*Book{
		CreateBook("Todo", 5, "Rose", time.Now()),
		CreateBook("Optimistic", 3, "Jack", time.Date(2018, time.November, 10, 0, 0, 0, 0, time.UTC)),
	}
}

func CreateStudent(name, grade, id, sex string) *Student {
	return &Student{Name: name, Grade: grade, Id: id, Sex: sex}
}

func (s *Student) AddBook(b *BookItem) error {
	for idx, book := range books {
		if book.Title == b.Book.Title {
			_, err := books[idx].Borrow(b.Amount)
			if err != nil {
				return err
			}
			s.Books = append(s.Books, b)
			return nil
		}
	}
	return ErrorNotFoundBook
}

func (s *Student) RemoveBook(b *BookItem) error {
	for idx, item := range s.Books {
		if item.Book.Title == b.Book.Title {
			item.Amount -= b.Amount
			books[idx].Quantity += b.Amount
			if item.Amount == 0 {
				s.Books = append(s.Books[:idx], s.Books[idx+1:]...)
			}
			return nil
		}
	}
	return ErrorNotFoundBook
}

func (s *Student) ListBooks() []*BookItem {
	return s.Books
}

func (s *Student) String() string {
	var books bytes.Buffer

	for _, v := range s.Books {
		books.WriteString(fmt.Sprintf("%s(%d)\n", v.Book.Title, v.Amount))
	}
	return fmt.Sprintf("%s | %s | %s | %s => %s\n", s.Name, s.Grade, s.Id, s.Sex,
		books.String())
}

func ShowBooks() {
	for _, v := range books {
		fmt.Println(v)
	}
	fmt.Println()
}
