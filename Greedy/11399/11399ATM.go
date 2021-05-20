// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/11399
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

	var minute = make([]int, 1001)
	var n, p, min, ans int
	Fscanln(r, &n)
	for ; n > 0; n-- {
		Fscanf(r, "%d ", &p)
		minute[p]++
	}
	for i, value := range minute {
		for ; value > 0; value-- {
			min += i
			ans += min
		}
	}
	Fprintln(w, ans)
}
