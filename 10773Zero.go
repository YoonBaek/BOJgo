// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10773

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

type Stack struct {
	stack [100000]int
	size  int
}

func push(s *Stack, d int) {
	s.stack[s.size] = d
	s.size++
	// Println(s.size)
}

func pop(s *Stack) int {
	var data int
	if s.size > 0 {
		s.size--
		data = s.stack[s.size]
	}
	return data
}

func main() {
	r, w := b.NewReader(os.Stdin), b.NewWriter(os.Stdout)
	defer w.Flush()
	stack := &Stack{}
	var k, n, ans int
	Fscanln(r, &k)

	for i := 0; i < k; i++ {
		Fscanln(r, &n)

		if n == 0 {
			pop(stack)
		} else {
			push(stack, n)
		}
	}
	for stack.size > 0 {
		ans += pop(stack)
	}
	Fprintln(w, ans)
}
