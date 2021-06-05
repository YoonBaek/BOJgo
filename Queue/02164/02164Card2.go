// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue struct {
	item chan int
}

func (q Queue) Push(n int) {
	q.item <- n
}

func (q Queue) Pop() int {
	return <-q.item
}

var (
	n  int
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	q := Queue{item: make(chan int, n)}
	for i := 0; i < n; i++ {
		q.item <- i + 1
	}
	for i := 0; i < n-1; i++ {
		q.Pop()
		q.item <- q.Pop()
	}
	fmt.Fprintln(wr, q.Pop())
}
