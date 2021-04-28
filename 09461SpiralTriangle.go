// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/9461

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func P(n int) []int {
	lengths := make([]int, n+1)
	lengths[1] = 1
	lengths[2] = 1
	lengths[3] = 1
	lengths[4] = 2
	lengths[5] = 2
	for k := 6; k <= n; k++ {
		lengths[k] = lengths[k-1] + lengths[k-5]
	}
	return lengths
}

func main() {
	r, w := b.NewReader(os.Stdin), b.NewWriter(os.Stdout)
	defer w.Flush()

	var N, T int
	Fscanln(r, &T)
	lengths := P(100)

	for ; T > 0; T-- {
		Fscanln(r, &N)
		Fprintln(w, lengths[N])
	}
}
