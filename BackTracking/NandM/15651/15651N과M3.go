// Solved by Github YoonBaek

package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	n, m    int
	visited []bool
	arr     []int
	rd      = bufio.NewReader(os.Stdin)
	wr      = bufio.NewWriter(os.Stdout)
)

func dfs(cnt int) {
	if cnt == m {
		for i := 0; i < m; i++ {
			Fprintf(wr, "%d ", arr[i])
		}
		Fprintln(wr)
	}
	for i := 0; i < n; i++ {
		visited[i] = true
		arr[cnt] = i + 1
		if cnt < m {
			dfs(cnt + 1)
		}
		visited[i] = false
	}
}

func main() {
	defer wr.Flush()
	Fscanf(rd, "%d %d\n", &n, &m)
	visited = make([]bool, n)
	arr = make([]int, n)
	dfs(0)
}
