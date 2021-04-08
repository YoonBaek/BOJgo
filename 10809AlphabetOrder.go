// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10809
package main

import (
	. "fmt"
)

func main() {
	s := ""
	Scanln(&s)
	var order = make([]int, 26)
	for i := range order {
		order[i] = -1
	}
	for i, str := range s {
		if order[str-97] == -1 {
			order[str-97] = i
		}
	}
	for _, i := range order {
		Printf("%d ", i)
	}
	Println()
}
