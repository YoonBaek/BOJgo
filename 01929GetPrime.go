// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/1929

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	w, r := bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdin)
	defer w.Flush()
	var m, n int
	Fscanln(r, &m, &n)

	primes := make([]bool, n+1)

	primes[1] = true

	for factor := 2; factor*factor <= n; factor++ {
		if !primes[factor] {
			for i := factor * 2; i <= n; i += factor {
				primes[i] = true
			}
		}
	}
	// Println(primes)
	for i := m; i <= n; i++ {
		if !primes[i] {
			Fprintln(w, i)
		}
	}
}
