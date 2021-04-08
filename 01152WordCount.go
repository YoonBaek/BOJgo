// Solved by Github Yoonbaek
// Q : https://www.acmicpc.net/problem/1152

package main

import (
	b "bufio"
	. "fmt"
	"os"
	str "strings"
)

func trim(s string) string {
	s = str.Trim(s, " ")
	s = str.TrimSpace(s)
	return s
}

func main() {
	r := b.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = trim(s)
	cnt := 0
	for _, a := range str.Split(s, " ") {
		if a != "" {
			cnt++
		}
	}
	Println(cnt)
}
