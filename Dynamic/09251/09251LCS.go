// Solved by Github YoonBaek
// memory = DP

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 변수 선언
var (
	X, Y   string
	sc     = bufio.NewScanner(os.Stdin)
	wr     = bufio.NewWriter(os.Stdout)
	memory [][]int
)

// input 받는 부분
func scanln() string {
	sc.Scan()
	return sc.Text()
}

// buffer print
func println(content interface{}) {
	fmt.Fprintln(wr, content)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// LCS 구현
func LCS(i, j int) {
	if X[i] == Y[j] {
		memory[i+1][j+1] = memory[i][j] + 1
	} else {
		memory[i+1][j+1] = max(memory[i+1][j], memory[i][j+1])
	}
}

func main() {
	defer wr.Flush()

	X = scanln()
	Y = scanln()

	// DP
	memory = make([][]int, len(X)+1)

	for i := range memory {
		memory[i] = make([]int, len(Y)+1)
	}

	for i := range X {
		for j := range Y {
			LCS(i, j)
		}
	}

	res := memory[len(X)][len(Y)]
	println(res)
}
