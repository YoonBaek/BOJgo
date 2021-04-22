// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/11653
// 제곱의 경우의 수를 줄여줌으로서 연산 시간을 줄여줄 수 있었다.

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
	factor := 2
	for factor <= n {
		if n%factor == 0 {
			n /= factor
			Fprintln(w, factor)
		} else {
			factor++
			if factor*factor > n {
				Fprintln(w, n)
				break
			}
		}
	}
}
