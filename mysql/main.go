package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Location struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_user")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query("SELECT user_id, username, sex, email FROM person WHERE user_id = ?", 1)
	exec("INSERT INTO person(username, sex, email) VALUES (?, ?, ?)", "jack", "male", "jack@ex.com")
	exec("UPDATE person SET username = ? WHERE user_id = ?", 1)
	exec("DELETE FROM person WHERE user_id = ?", 1)
}

func exec(cmd string, params ...interface{}) {
	result, err := db.Exec(cmd, params...)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert successfully: %d\n", id)
}

func query(cmd string, params ...interface{}) {
	rows, err := db.Query(cmd, params...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := Person{}
		err := rows.Scan(&p.UserId, &p.Username, &p.Sex, &p.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("person: %+v", p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
