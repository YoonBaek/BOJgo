// Solved by Github YoonBaek

package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

var (
	n, a, max int
	ints, dp  []int
	sc        = bufio.NewScanner(os.Stdin)
	wr        = bufio.NewWriter(os.Stdout)
)

func LISN2() {
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if ints[i] > ints[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	ints = make([]int, n)
	dp = make([]int, n)

	sc.Scan()
	strs := strings.Split(sc.Text(), " ")

	for i := 0; i < n; i++ {
		a, _ = strconv.Atoi(strs[i])
		ints[i] = a
	}
	LISN2()

	Println(max, dp)
}
