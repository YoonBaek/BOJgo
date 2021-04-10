// Solved By Github YoonBaek
// Q : https://www.acmicpc.net/problem/1316

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	n, c, cnt := 0, 0, 0
	input := ""
	Fscanln(r, &n)
	for n > 0 {
		Fscanln(r, &input)
		var checker = make([]int, 26)
		c = 1
		v := 'A'
		for _, val := range input {
			if val == v {
				continue
			} else if checker[val-'a'] == 0 {
				checker[val-'a'] += 1
			} else {
				c = 0
				break
			}
			v = val
		}
		cnt += c
		n--
	}
	Println(cnt)
}
