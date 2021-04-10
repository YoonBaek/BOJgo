// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2941

package main

import (
	b "bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	r := b.NewReader(os.Stdin)
	var input string

	Fscanln(r, &input)
	input = strings.Replace(input, "dz=", "z=", -1)
	cnt := len(input)

	before := 'A'
	for _, val := range input {
		if val == 'j' {
			if before == 'l' || before == 'n' {
				cnt--
			}
		}
		if val == '=' {
			if before == 'c' || before == 's' || before == 'z' {
				cnt--
			}
		}
		if val == '-' {
			if before == 'c' || before == 'd' {
				cnt--
			}
		}
		before = val
	}
	Println(cnt)
}
