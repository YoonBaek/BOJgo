// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/8958
package main

import (
	b "bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	r := b.NewReader(os.Stdin)
	n, sum, cnt := 0, 0, 0
	ox := ""
	Fscanln(r, &n)
	for i := 0; i < n; i++ {
		Fscanln(r, &ox)
		Os := strings.Split(ox, "X")
		sum = 0
		for _, O := range Os {
			cnt = 0
			for j := 1; j <= len(O); j++ {
				cnt += j
			}
			sum += cnt
		}
		Println(sum)
	}
}

// You can solve this by just using rune
