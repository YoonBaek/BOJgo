// Solved By Github YoonBaek
// Q : https://www.acmicpc.net/problem/10757
// 미완성된 코드입니다.

package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var a, b string

	Fscanf(r, "%s %s\n", &a, &b)
	digitsA, digitsB := make([]int, len(a)), make([]int, len(b))
	for i, Rune := range a {
		digitsA[len(a)-1-i] = int(Rune - 48)
		if i > len(b)-1 {
			continue
		}
		digitsB[len(b)-1-i] = int(b[i] - 48)
	}
	sum := 0

	for i := 0; i < len(a); i++ {
		if i != 0 && i%18 != 0 {
			if i > len(b) {
				sth := (digitsA[i] + digitsB[i]) * int(math.Pow10(i))
				continue
			}
			sth := (digitsA[i] + digitsB[i]) * int(math.Pow10(i))
			sum += sth
			Println(sum)
		}
	}
	Fprintln(w, a, b, len(a), digitsA, digitsB)
}
