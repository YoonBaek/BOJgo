// Solved by Github YoonBaek
// Q: https://www.acmicpc.net/problem/4153

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for true {
		var a, b, c, max, left int
		Fscanf(r, "%d %d %d\n", &a, &b, &c)

		if a >= max {
			max = a
			left = b*b + c*c
		}
		if b >= max {
			max = b
			left = a*a + c*c
		}
		if c >= max {
			max = c
			left = b*b + a*a
		}
		if a == 0 && b == 0 && c == 0 {
			break
		} else if max*max == left {
			Fprintln(w, "right")
		} else {
			Fprintln(w, "wrong")
		}
	}
}
