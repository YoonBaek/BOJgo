package main

import (
	"bufio"
	. "fmt"
	"os"
)

type Queue struct {
	item []int
	size int
}

func (q *Queue) Push(val int) {
	q.item = append(q.item, val)
	q.size++
}

func (q *Queue) Pop() int {
	ret := -1
	if q.size > 0 {
		ret = q.item[0]
		q.item = q.item[1:]
	}
	if q.size != 0 {
		q.size--
	}
	return ret
}

func (q *Queue) Empty() int {
	if q.size == 0 {
		return 1
	} else {
		return 0
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Front() int {
	ret := -1
	if q.size > 0 {
		ret = q.item[0]
	}
	return ret
}

func (q *Queue) Back() int {
	ret := -1
	if q.size > 0 {
		ret = q.item[q.size-1]
	}
	return ret
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	queue := Queue{}
	var n int
	Fscanln(r, &n)
	var cmd string
	for ; n > 0; n-- {
		elem := 0
		Fscanf(r, "%s %d\n", &cmd, &elem)

		switch cmd {
		case "push":
			queue.Push(elem)
		case "pop":
			pop := queue.Pop()
			Fprintln(w, pop)
		case "front":
			Fprintln(w, queue.Front())
		case "back":
			Fprintln(w, queue.Back())
		case "size":
			Fprintln(w, queue.Size())
		case "empty":
			Fprintln(w, queue.Empty())
		}
	}
}
