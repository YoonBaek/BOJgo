// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/1065https://www.acmicpc.net/problem/1065
package main

import (
	. "fmt"
	"math"
)

func getDigits(n int) []int {
	digits := []int{}
	for i := 3; i >= 0; i-- {
		digit := n / int(math.Pow10(i)) % 10
		if digit == 0 && len(digits) == 0 {
			continue
		}
		digits = append(digits, digit)
	}
	return digits
}

func main() {
	n, cnt := 0, 0
	Scanln(&n)
	for i := 1; i <= n; i++ {
		digits := getDigits(i)
		if len(digits) < 3 {
			cnt++
		} else if len(digits) != 4 {
			// in case that digits > 3, this condition should be fixed.
			if digits[1]-digits[0] == digits[2]-digits[1] {
				cnt++
			}
		}
	}
	Println(cnt)
}
