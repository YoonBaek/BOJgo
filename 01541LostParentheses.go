// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/1541

package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var input string
	Fscanln(r, &input)

	var nums string
	var ans, num int
	var pm bool = true

	for _, R := range input {
		if R > 47 && R < 58 {
			nums += string(R)
		} else {
			num, _ = strconv.Atoi(nums)
			nums = ""
			if pm {
				ans += num
				if R != 43 {
					pm = false
				}
			} else if !pm {
				ans -= num
			}
		}
	}
	if !pm {
		num, _ = strconv.Atoi(nums)
		ans -= num
	} else {
		num, _ = strconv.Atoi(nums)
		ans += num
	}
	Fprintln(w, ans)
}
