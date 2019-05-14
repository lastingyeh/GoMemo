package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := Student{
		Name:  "jack",
		Age:   19,
		Score: 99,
	}
	test(stu)

	b := 1
	testInt(&b)
	fmt.Println("after setInt:", b)

	// testStruct
	testStruct(&stu)
	fmt.Println("after setString for Name: ", stu)
}

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

func (s Student) Print() {
	fmt.Printf("%+v\n", s)
}

func test(obj interface{}) {
	// TypeOf -> main.Student (defined-type)
	t := reflect.TypeOf(obj)
	fmt.Println(t)

	// ValueOf -> {jack 19 99}
	v := reflect.ValueOf(obj)
	fmt.Println(v)

	// ValueOf().Kind() -> struct (prototype)
	k := reflect.ValueOf(obj).Kind()
	fmt.Println(k)

	// ValueOf(obj).Interface() ->
	iv := reflect.ValueOf(obj).Interface()
	if stu, ok := iv.(Student); ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func testInt(obj interface{}) {
	// ref
	val := reflect.ValueOf(obj)
	// Elem().Int() -> 1
	fmt.Println("before SetInt:", val.Elem().Int())

	val.Elem().SetInt(99)

}

func testStruct(obj interface{}) {
	val := reflect.ValueOf(obj)

	k := val.Kind()
	if k != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("unexpected struct")
		return
	}

	fLength := val.Elem().NumField()
	mLength := val.Elem().NumMethod()
	fmt.Println("fLength: ", fLength)
	// public method number
	fmt.Println("mLength: ", mLength)

	// call method
	var params []reflect.Value
	val.Elem().Method(0).Call(params)

	// fields
	for i := 0; i < fLength; i++ {
		fmt.Println("tag = ", val.Elem().Type().Field(i).Tag.Get("json"))
		fmt.Printf("%v %v\n", val.Elem().Field(i), val.Elem().Field(i).Kind())
	}

	val.Elem().FieldByName("Name").SetString("update")
}
