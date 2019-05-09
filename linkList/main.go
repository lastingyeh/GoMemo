package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

var p Student

func init() {
	//init data
	p = Student{Name: "head", Age: 20, Score: 99}
	listInit(&p)
}

func main() {
	//count
	fmt.Printf("count: %d\n", p.count(true)) //10

	//find
	index := 2
	fmt.Printf("find index: %d, value: %+v\n", index, p.get(index))

	//push
	p.append(&Student{Name: "last", Age: 99, Score: 100})
	fmt.Printf("after push - count: %d\n", p.count(true)) //11

	//insert by index
	p.insert(index, &Student{Name: "insert", Age: 50, Score: 50})
	fmt.Printf("after insert - count: %d\n", p.count(true))

	//unshift
	p.prepend(&Student{Name: "new head", Age: 1, Score: 1})
	fmt.Printf("after unshift - count: %d\n", p.count(true))

	//remove by index
	fmt.Printf("remove: %+v\n", p.remove(index))
	fmt.Printf("after remove index: %d - count: %d\n", index, p.count(true))
}

func listInit(p *Student) {
	for i := 1; i < 10; i++ {
		stu := &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		p.next = stu
		p = stu
	}
}

func (s *Student) append(p *Student) {
	last := s.get(s.count(false) - 1)
	last.next = p
}

func (s *Student) get(index int) *Student {
	if index < 0 || index > s.count(false)-1 {
		panic("index out of bounds")
	}

	var n = s
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n
}

func (s *Student) insert(index int, p *Student) {
	switch {
	case index <= 0:
		s.prepend(p)
	case index > s.count(false)-1:
		s.append(p)
	default:
		prev := s.get(index - 1)
		p.next = prev.next
		prev.next = p
	}
}

func (s *Student) prepend(p *Student) {
	*s, *p = *p, *s
	s.next = p
}

func (s *Student) count(show bool) int {
	var c int
	var n = s

	for n != nil {
		if show {
			fmt.Printf("%+v\n", n)
		}
		n = n.next
		c++
	}
	return c
}

func (s *Student) toString() {
	s.count(true)
}

func (s *Student) remove(index int) Student {
	c := s.count(false)
	cur := *s.get(index)

	switch {
	case index == 0:
		next := s.get(index + 1)
		*next, *s = *s, *next
	case index == c-1:
		last := s.get(c - 2)
		last.next = nil
	case index > 0 || index < c-1:
		prev := s.get(index - 1)
		prev.next = cur.next
	default:
		panic("index out of bounds")
	}
	return cur
}
