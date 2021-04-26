// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/7568
package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, wr := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	n := 0
	Fscanln(r, &n)
	ws := make([]int, n)
	hs := make([]int, n)
	ranks := make([]int, n)

	for i := 0; i < n; i++ {
		w, h := 0, 0
		Fscanln(r, &w, &h)
		ws[i] = w
		hs[i] = h
	}

	for i := 0; i < n; i++ {
		rank := 1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if ws[i] < ws[j] && hs[i] < hs[j] {
				rank++
			}
		}
		ranks[i] = rank
		Fprintln(wr, rank)
	}
}
