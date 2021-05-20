// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/11047
// Greedy Algorithm 문제

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k, a, cnt int
	Fscanln(r, &n, &k)
	var coins = make([]int, n)

	for i := 0; i < n; i++ {
		Fscanf(r, "%d ", &a)
		coins[n-1-i] = a
	}
	for _, coin := range coins {
		if k >= coin {
			cnt += k / coin
			k %= coin
		}
	}
	Fprintln(w, cnt)
}
