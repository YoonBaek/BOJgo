// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/5622

package main

import . "fmt"

func main() {
	var input string
	Scanln(&input)
	duration := 0
	for _, val := range input {
		if val >= 'A' {
			duration += 3
		}
		if val >= 'D' {
			duration += 1
		}
		if val >= 'G' {
			duration += 1
		}
		if val >= 'J' {
			duration += 1
		}
		if val >= 'M' {
			duration += 1
		}
		if val >= 'P' {
			duration += 1
		}
		if val >= 'T' {
			duration += 1
		}
		if val >= 'W' {
			duration += 1
		}
	}
	Println(duration)
}
