// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/9012

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := b.NewReader(os.Stdin), b.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		T, stack int
		P, isVPS string
	)

	Fscanln(r, &T)
	for ; T > 0; T-- {
		Fscanln(r, &P)
		stack = 0
		for _, p := range P {
			if p == 40 {
				stack++
			} else if p == 41 {
				stack--
			}
			if stack < 0 {
				break
			}
		}
		if stack == 0 {
			isVPS = "YES"
		} else {
			isVPS = "NO"
		}
		Fprintln(w, isVPS)
	}
}
