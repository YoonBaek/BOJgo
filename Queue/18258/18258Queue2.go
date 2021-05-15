// Solved by Github YoonBaek

package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

type Queue struct {
	head *Elem
	tail *Elem
	size int
}

type Elem struct {
	value int
	next  *Elem
}

func (q *Queue) Front() {
	if q.head == nil {
		Fprintf(w, "%d\n", -1)
	} else {
		Fprintf(w, "%d\n", q.head.value)
	}
}

func (q *Queue) Back() {
	if q.head == nil {
		Fprintf(w, "%d\n", -1)
	} else {
		Fprintf(w, "%d\n", q.tail.value)
	}
}

func (q *Queue) Push(n int) {
	if q.head == nil {
		q.tail = &Elem{n, nil}
		q.head = q.tail
	} else {
		q.tail.next = &Elem{n, nil}
		q.tail = q.tail.next
	}
	q.size++
}

func (q *Queue) Pop() {
	q.Front()
	if q.head != nil {
		q.head = q.head.next
		q.size--
	}

}

func (q *Queue) Size() {
	Fprintf(w, "%d\n", q.size)
}

func (q *Queue) Empty() {
	if q.size == 0 {
		Fprintf(w, "%d\n", 1)
	} else {
		Fprintf(w, "%d\n", 0)
	}
}

func main() {
	defer w.Flush()
	var n, v int
	q := Queue{}
	cmd := ""
	Fscanln(r, &n)
	for ; n > 0; n-- {
		Fscanln(r, &cmd, &v)
		switch cmd {
		case "push":
			q.Push(v)
		case "pop":
			q.Pop()
		case "front":
			q.Front()
		case "back":
			q.Back()
		case "empty":
			q.Empty()
		case "size":
			q.Size()
		}
	}
}
