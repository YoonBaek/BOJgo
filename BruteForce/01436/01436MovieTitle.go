// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	memory = []string{}
	sc     = bufio.NewScanner(os.Stdin)
	wr     = bufio.NewWriter(os.Stdout)
)

func solve1(n int) {
	// slow....
	i := 666
	for true {
		m := strconv.Itoa(i)
		for i := 2; i < len(m); i++ {
			if m[i-2] == '6' && m[i-1] == '6' && m[i] == '6' {
				memory = append(memory, m)
				break
			}
		}
		i++
		if len(memory) >= n {
			fmt.Fprintln(wr, memory[n-1])
			break
		}
	}
}

func solve2(n int) {
	// fast
	i := 666
	for true {
		cur := i
		for cur >= 666 {
			if cur%1000 == 666 {
				n--
				break
			}
			cur /= 10
		}
		if n <= 0 {
			fmt.Fprintln(wr, i)
			break
		}
		i++
	}
}

func main() {
	defer wr.Flush()

	var n int
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	solve2(n)
}
