// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/15596
package main

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
