// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/16395
// 재귀로 푸는 것 보단 배열로 푸는게 빠를 거 같아서 배열로 품

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func makeTable() [][]int {
	tbl := make([][]int, 30)

	for i := 0; i < 30; i++ {
		row := make([]int, 30)
		row[0], row[i] = 1, 1
		for n := 1; n < i; n++ {
			row[n] = tbl[i-1][n] + tbl[i-1][n-1]
		}
		tbl[i] = row
	}
	return tbl
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k int
	Fscanln(r, &n, &k)
	tbl := makeTable()
	Fprintln(w, tbl[n-1][k-1])
}
