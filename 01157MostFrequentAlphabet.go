// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10809

package main

import (
	b "bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	r := b.NewReader(os.Stdin)
	s, result := "", ""
	max := 0
	// i guess that map costs some time... 28s in BOJ
	// strs := map[string]int{}
	Fscanln(r, &s)
	s = strings.ToUpper(s)
	strs := make([]int, 26)
	for _, str := range s {
		strs[str-65] += 1
	}
	for k, v := range strs {
		if v > max {
			result = string(k + 65)
			max = v
		} else if v == max {
			result = "?"
		}
	}
	Println(result)
}
