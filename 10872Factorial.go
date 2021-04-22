// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10872
// 문제에서 재귀함수로 풀어달라고 했으니, 재귀함수로 풀어본다.

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	n := 0
	Fscanln(r, &n)
	Fprintln(w, factorial(n))
}
