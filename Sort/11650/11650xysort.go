// Solved by Github Yoonbaek
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Memory [200001][]int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, x, y int

	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	m := Memory{}
	temp := []int{}

	for i := 0; i < n; i++ {
		if sc.Scan() {
			fields := strings.Fields(sc.Text())
			x, _ = strconv.Atoi(fields[0])
			y, _ = strconv.Atoi(fields[1])
		}
		x += 100001
		temp = append(temp, 0)
		fmt.Fprintln(wr, y, "temp :", temp)
		remark := 0
		for i := 0; i < len(m[x]); i++ {
			if y < m[x][i] {
				fmt.Fprintln(wr, "Before", i, temp, m[x])
				temp[i+1] = m[x][i] + 1 - 1
				fmt.Fprintln(wr, "After", i, temp, m[x])
			} else {
				remark++
			}
		}
		temp[remark] = y
		m[x] = temp
	}
	for i, x := range m {
		for _, y := range x {
			fmt.Fprintln(wr, i-100001, y)
		}
	}
}
