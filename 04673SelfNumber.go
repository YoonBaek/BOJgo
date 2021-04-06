// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/4673
package main

import (
	. "fmt"
	"math"
)

func selfNumber(n int) int {
	sN := n
	for i := 0; i < 5; i++ {
		sN += n / int(math.Pow10(i)) % 10
	}
	return sN
}

func sNSlices(emptySlice []int) []int {
	for i := 1; i <= 10000; i++ {
		sN := selfNumber(i)
		if sN <= 10000 {
			emptySlice[sN-1] = 1
		}
	}
	return emptySlice
}

func main() {
	a := make([]int, 10000)
	a = sNSlices(a)
	for i, v := range a {
		if v == 0 {
			Println(i + 1)
		}
	}
}
