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
	var n, m, cntB, cntW int
	min := 2500

	if sc.Scan() {
		input := sc.Text()
		inputs := strings.Split(input, " ")
		n, _ = strconv.Atoi(inputs[0])
		m, _ = strconv.Atoi(inputs[1])
	}

	Board := make([][]rune, n)

	for i := 0; i < n; i++ {
		if sc.Scan() {
			inputs := sc.Text()
			Board[i] = []rune(inputs)
		}
	}

	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			cntB, cntW = 0, 0
			for k := 0; k < 8; k++ {
				for l := 0; l < 8; l++ {
					if (k+l)%2 == 0 {
						if Board[i+k][j+l] == 'B' {
							cntW++
						} else {
							cntB++
						}
					} else {
						if Board[i+k][j+l] == 'B' {
							cntB++
						} else {
							cntW++
						}
					}
				}
			}
			if cntW < min {
				min = cntW
			}
			if cntB < min {
				min = cntB
			}
		}
	}
	fmt.Fprintln(wr, min)
}
