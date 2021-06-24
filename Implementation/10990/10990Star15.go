// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func iScanf() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func print(pr string) {
	fmt.Fprint(wr, pr)
}

func renderGraph(memory [][]bool) []byte {
	graph := []byte{}
	for _, row := range memory {
		for _, elem := range row {
			if !elem {
				graph = append(graph, ' ')
			} else {
				graph = append(graph, '*')
			}
		}
		graph = append(graph, '\n')
	}
	return graph
}

func main() {
	n := 0
	defer wr.Flush()

	n = iScanf()

	memory := make([][]bool, n)

	for i := range memory {
		memory[i] = make([]bool, n+i)
		memory[i][n+i-1] = true
		memory[i][n-i-1] = true
	}
	print(string(renderGraph(memory)))
}
