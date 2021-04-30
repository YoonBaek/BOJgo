// Solved by Github YoonBaek
// Q : https: //www.acmicpc.net/problem/13297
// ec : estimated cost
package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	n, ec := 0, ""

	Fscanln(r, &n)
	for ; n > 0; n-- {
		Fscanln(r, &ec)
		Fprintln(w, len(ec))
	}
}
