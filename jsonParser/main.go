package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	UserName string    `json:"username"`
	NickName string    `json:"nickname"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
	Sex      string    `json:"sex"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

func main() {
	test(_struct(), &User{})

	var m map[string]interface{}
	test(_map("user02", 19, "female"), &m)

	var s []map[string]interface{}
	test(_slice(), &s)
}

func _struct() *User {
	return &User{
		UserName: "user01",
		NickName: "picky",
		Age:      18,
		Birthday: time.Date(2018, 8, 8, 0, 0, 0, 0, time.Local),
		Sex:      "男",
		Email:    "user01@cc.com",
		Phone:    "012345678",
	}
}

func _map(name string, age int, sex string) map[string]interface{} {
	m := make(map[string]interface{})
	m["username"] = name
	m["age"] = age
	m["sex"] = sex

	return m
}

func _slice() []map[string]interface{} {
	s := make([]map[string]interface{}, 2)
	s[0] = _map("user03", 20, "male")
	s[1] = _map("user04", 21, "female")

	return s
}

func serialize(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

func deserialize(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

func test(v interface{}, target interface{}) {
	if str, err := serialize(v); err != nil {
		fmt.Println("serialize error: ", err)
	} else {
		fmt.Println("serialize: ", str)
		if err = deserialize(str, target); err != nil {
			fmt.Println("deserialize error: ", err)
		}
		fmt.Printf("deserialize: %+v\n", target)
	}
}
