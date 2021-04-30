// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2442

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	n := 0

	Fscanln(r, &n)
	for i := 1; i <= n; i++ {
		content := ""
		j := n - i
		k := 2*i - 1
		for ; j > 0; j-- {
			content += " "
		}
		for ; k > 0; k-- {
			content += "*"
		}
		Fprintln(w, content)
	}
}
