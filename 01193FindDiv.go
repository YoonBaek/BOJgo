// Solved By Github YoonBaek
// Q : https://www.acmicpc.net/problem/1193

package main

import "fmt"

func getBoonmo(x int) int {
	v, limit := 0, 2
	check := -1
	for cnt := 1; cnt <= x; cnt++ {
		if check == -1 {
			v++
		} else if check == 0 {
			v = 1
			check--
		} else {
			v--
		}
		if v == limit {
			check = 1
			limit += 2
		} else if v == 1 && check != -1 {
			check = 0
		}
	}
	return v
}

func getBoonja(x int) int {
	v, limit := 0, 1
	check := -1
	for cnt := 1; cnt <= x; cnt++ {
		if check == -1 {
			v++
		} else if check == 0 {
			v = 1
			check = -1
		} else {
			v--
		}
		if cnt == 1 {
			check = 0
			limit = limit*2 + 1
		}
		if v == limit && cnt != 1 {
			check = 1
			limit = limit + 2
		} else if v == 1 && check == 1 {
			check = 0
		}
	}
	return v
}

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Printf("%d/%d\n", getBoonja(x), getBoonmo(x))
}
