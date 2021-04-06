// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/11720
package main

import (
	. "fmt"
)

func main() {
	var (
		n, sum int
		num    string
	)
	Scanf("%d\n%s\n", &n, &num)
	for _, i := range num {
		Println(i - '0')
		// sum += int(i - '0')
	}
	Println(sum)
}
