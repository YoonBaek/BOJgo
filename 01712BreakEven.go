// Solved By Github YoonBaek
// Q : https://www.acmicpc.net/problem/1712
package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a, b, c, n int
	Fscanf(r, "%d %d %d\n", &a, &b, &c)
	if b >= c {
		n = -1
	} else {
		n = a/(c-b) + 1
	}
	Println(n)
}
