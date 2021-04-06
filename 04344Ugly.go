// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/4344
package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	r := b.NewReader(os.Stdin)
	var (
		c, n          int
		sum, avg, cnt float64
	)
	Fscanln(r, &c)
	for i := 0; i < c; i++ {
		Fscan(r, &n)
		scores := make([]float64, n)
		sum, cnt = 0, 0
		for i := 0; i < n; i++ {
			Fscanf(r, " %f", &scores[i])
			sum += scores[i]
		}
		avg = sum / float64(n)
		for _, score := range scores {
			if score > avg {
				cnt += 100
			}
		}
		Printf("%.3f%%\n", cnt/float64(n))
	}
}
