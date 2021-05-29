// Solved by Github YoonBaek
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m, ans int
	min := 300000

	if sc.Scan() {
		strs := strings.Fields(sc.Text())
		n, _ = strconv.Atoi(strs[0])
		m, _ = strconv.Atoi(strs[1])
	}

	cards := make([]int, n)

	if sc.Scan() {
		strs := strings.Fields(sc.Text())
		for i := 0; i < n; i++ {
			cards[i], _ = strconv.Atoi(strs[i])
		}
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := cards[i] + cards[j] + cards[k]
				if sum <= m {
					if m-sum < min {
						ans = sum
						min = m - sum
						if min == 0 {
							break
						}
					}
				}
			}
		}
	}
	fmt.Fprintln(wr, ans)
}
