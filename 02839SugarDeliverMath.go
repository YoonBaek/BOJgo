// Solved By Github YoonBaek
// Q : https://www.acmicpc.net/problem/2839

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func check(x int) int {
	big := x / 5
	left := x % 5
	for left%3 != 0 {
		big--
		left += 5
	}
	if big < 0 {
		return -1
	}
	return big + left/3
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := 0
	Fscanln(r, &n)
	Fprintln(w, check(n))
}
