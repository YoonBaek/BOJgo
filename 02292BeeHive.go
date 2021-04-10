// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2292

package main

import . "fmt"

func main() {
	var n, i, sum int
	Scanln(&n)
	for true {
		sum += i
		i++
		if n > 6*sum+1 {
			continue
		} else {
			Println(i)
			break
		}
	}
}
