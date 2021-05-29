// solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dp = [500][500]int{}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, max int

	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	for d := 0; d < n; d++ {
		if sc.Scan() {
			strs := strings.Split(sc.Text(), " ")
			inputs := make([]int, len(strs))
			for i := 0; i < len(strs); i++ {
				inputs[i], _ = strconv.Atoi(strs[i])
			}
			if d == 0 {
				dp[d][0] = inputs[0]
				continue
			}
			for b := 0; b < len(inputs); b++ {
				if b == 0 {
					dp[d][b] = dp[d-1][b] + inputs[b]
				} else if b == d {
					dp[d][b] = dp[d-1][b-1] + inputs[b]
				} else {
					dp[d][b] = Max(dp[d-1][b-1], dp[d-1][b]) + inputs[b]
				}
				if d == n-1 {
					max = Max(max, dp[d][b])
				}
			}
		}
	}
	fmt.Fprintln(wr, max)
}
