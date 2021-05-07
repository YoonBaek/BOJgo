// Solved by Github Yoon Baek

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, split, nCopy, diff int
	Fscanln(r, &n)
	nCopy = n
	for nCopy > 0 {
		nCopy /= 10
		diff++
	}
	nCopy = 0
	for i := n - 9*diff; i < n; i++ {
		iCopy := i
		sum := i
		for iCopy != 0 {
			split = iCopy % 10
			iCopy /= 10
			sum += split
		}
		if sum == n {
			nCopy = i
			break
		}
	}
	Fprintln(w, nCopy)
}
