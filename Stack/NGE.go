// solved by Github YoonBaek
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	NGE []int
	n   int
	sc  *bufio.Scanner
	wr  *bufio.Writer
)

type Stack struct {
	stack []int
	sIdx  []int
	size  int
}

func (s *Stack) Push(i int, n int) {
	s.stack = append(s.stack, n)
	s.sIdx = append(s.sIdx, i)
	s.size++
}

func (s *Stack) Pop() (int, int) {
	pop := s.stack[s.size-1]
	popI := s.sIdx[s.size-1]
	s.stack = s.stack[:s.size-1]
	s.sIdx = s.sIdx[:s.size-1]
	s.size--
	return popI, pop
}

func (s *Stack) Top() int {
	top := s.stack[s.size-1]
	return top
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	NGE = make([]int, n)
	elems := make([]int, n)
	if sc.Scan() {
		fields := strings.Fields(sc.Text())
		for i := 0; i < n; i++ {
			elems[i], _ = strconv.Atoi(fields[i])
		}
	}

	s := Stack{}

	s.Push(0, elems[0])

	for i := 1; i < n; i++ {
		for s.size != 0 && s.Top() < elems[i] {
			idx, _ := s.Pop()
			NGE[idx] = elems[i]
		}
		s.Push(i, elems[i])
	}
	for _, nge := range NGE {
		if nge == 0 {
			fmt.Fprintf(wr, "%d ", -1)
		} else {
			fmt.Fprintf(wr, "%d ", nge)
		}
	}
	fmt.Fprintln(wr)
}
