// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10870
// 역시 재귀함수로 풀어보란다...
// 0번째, 1번째만 조심하면 된다.

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func fibo(x int) int {
	if x <= 1 {
		return x
	} else {
		return fibo(x-1) + fibo(x-2)
	}
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	n := 0
	Fscanln(r, &n)
	Fprintln(w, fibo(n))
}
