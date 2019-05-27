package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type User struct {
	Title   string
	Name    string
	Age     int
	Courses []map[string]interface{}
}

const form = `
	<html>
		<body>
			<form action="#" method="post">
				<input type="text" name="in"/>
				<input type="text" name="in"/>
				<input type="submit" value="submit"/>
			</form>
		</body>
	</html>
`

// simple get
func simpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello Test</h1>")
	panic("simple panic")
}

// form get | post
func postServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}
}

// template render get
func templateServer(w http.ResponseWriter, r *http.Request) {
	// locate file at $GOPATH path
	temp, err := template.ParseFiles("./$GOPATH/xxx/index.html")
	if err != nil {
		panic("template render error: " + err.Error())
	}
	// age -> private field (invisible)
	p1 := make(map[string]interface{})
	p1["id"] = 1
	p1["course"] = "javascript"
	p1["date"] = "2019/1/1 - 2019/4/18"

	p2 := make(map[string]interface{})
	p2["id"] = 2
	p2["course"] = "golang"
	p2["date"] = "2019/3/1 - 2019/6/30"

	courses := []map[string]interface{}{p1, p2}
	p := User{Title: "MyBlog", Name: "Jason", Age: 25, Courses: courses}

	err = temp.Execute(w, p)
	if err != nil {
		panic("template execute: " + err.Error())
	}
}

// middleware for recover panic
func logPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("catch handle error: %v", err)
			}
		}()
		handle(w, r)
	}
}

// server init
func main() {
	http.HandleFunc("/simple", logPanic(simpleServer))
	http.HandleFunc("/post", logPanic(postServer))
	http.HandleFunc("/temp", logPanic(templateServer))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("listen server error: ", err)
	}
}
