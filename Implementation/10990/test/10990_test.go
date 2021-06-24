// Solved by Github YoonBaek

package star15

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

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

func Test1(t *testing.T) {
	n := 3000

	wr := bufio.NewWriter(os.Stdout)
	grid := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]bool, n+i)
		grid[i][n-i-1] = true
		grid[i][n+i-1] = true
	}
	fmt.Fprintf(wr, string(renderGraph(grid)))
}
