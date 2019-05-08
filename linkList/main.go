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
	fmt.Printf("find index: %d, value: %+v\n", index, p.find(index))

	//push
	p.push(&Student{Name: "last", Age: 99, Score: 100})
	fmt.Printf("after push - count: %d\n", p.count(true)) //11

	//insert
	p.insert(0, &Student{Name: "insert", Age: 50, Score: 50})
	fmt.Printf("after insert - count: %d\n", p.count(true))

	//unshift
	p.unshift(&Student{Name: "new head", Age: 1, Score: 1})
	fmt.Printf("after unshift - count: %d\n", p.count(true))
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

func (s *Student) push(p *Student) {
	last := s.find(s.count(false) - 1)
	last.next = p
}

func (s *Student) find(index int) *Student {
	var n = s
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n
}

func (s *Student) insert(index int, p *Student) {
	switch {
	case index <= 0:
		s.unshift(p)
	case index > s.count(false)-1:
		s.push(p)
	default:
		prev := s.find(index - 1)
		p.next = prev.next
		prev.next = p
	}
}

func (s *Student) unshift(p *Student) {
	//var n = *s
	//*p = *s
	//*p.next = *s.next
	//return n

	//fmt.Printf("%v, %p\n", s, s)

	*s, *p = *p, *s
	s.next = p

	//fmt.Printf("%v, %p\n", s, s)
	//fmt.Printf("%v, %p\n", p, p)
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

func (s *Student) iterator() {
	s.count(true)
}
