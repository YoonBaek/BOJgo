// Solved by Github YoonBaek
// https://www.acmicpc.net/problem/2908
package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

// You can manage largenumber by modifyng p values here.
// if p = 5, you can do 654321 to 123456
func reverse(a int, p int) int {
	val := 0
	for i := p; i >= 0; i-- {
		val += a / int(math.Pow10(i)) % 10 * int(math.Pow10(p-i))
	}
	return val
}

func main() {
	a, b, c := 0, 0, 0
	r := bufio.NewReader(os.Stdin)
	Fscanf(r, "%d %d\n", &a, &b)
	a = reverse(a, 5)
	b = reverse(b, 5)
	c = b
	if a > b {
		c = a
	}
	Println(c)
}
