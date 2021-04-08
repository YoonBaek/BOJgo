// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10809

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	reader := b.NewReader(os.Stdin)
	var (
		t, r      int
		s, result string
	)
	Scanln(&t)
	for i := 0; i < t; i++ {
		Fscanf(reader, "%d %s\n", &r, &s)
		result = ""
		for _, str := range s {
			for k := 0; k < r; k++ {
				result += string(str)
			}
		}
		Println(result)
	}
}
