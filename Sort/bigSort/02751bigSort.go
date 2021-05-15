// Solved by Github YoonBaek

package main

import (
	"bufio"
	. "fmt"
	"os"
)

const max = 2000001

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, elem int

	Fscanln(r, &n)
	arr := make([]bool, max)
	for ; n > 0; n-- {
		Fscanln(r, &elem)
		arr[elem+1000000] = true
	}
	for key, val := range arr {
		if val {
			Fprintln(w, key-1000000)
		}
	}
}
