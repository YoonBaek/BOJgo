// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10817

package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c int
	Fscanf(r, "%d %d %d\n", &a, &b, &c)
	ints := []int{a, b, c}
	sort.Ints(ints)
	Fprintln(w, ints[1])
}
