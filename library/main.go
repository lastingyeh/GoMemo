package main

import (
	"fmt"
	"go_dev/library/model"
)

func main() {
	// create Student
	stu := model.CreateStudent("Student01", "class1", "001", "m")
	// Add bookItem from 'student01'
	addBookItem := &model.BookItem{
		Book:   &model.Book{Title: "Todo"},
		Amount: 2,
	}
	fmt.Println("---------- Original Books ----------")
	model.ShowBooks()

	fmt.Println("++++++++++ Student01 checkout Todo * 2 ++++++++++")
	if err := stu.AddBook(addBookItem); err != nil {
		fmt.Println("stu.AddBook error: ", err)
		return
	}
	fmt.Printf("%s", stu)

	fmt.Println("---------- After Check Out ----------")
	model.ShowBooks()

	// Remove bookItem from 'student01'
	removeBookItem := &model.BookItem{
		Book:   &model.Book{Title: "Todo"},
		Amount: 1,
	}

	fmt.Println("++++++++++ Student01 checkin Todo * 1 ++++++++++")
	if err := stu.RemoveBook(removeBookItem);err != nil{
		fmt.Println("stu.RemoveBook error: ", err)
		return
	}
	fmt.Printf("%s", stu)

	fmt.Println("---------- After Check In ----------")
	model.ShowBooks()
}
