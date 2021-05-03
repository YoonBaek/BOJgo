// Solved by github YoonBaek
// Q : https://www.acmicpc.net/problem/1260

package main

import (
	"bufio"
	. "fmt"
	"os"
)

const max = 1001

var (
	n, m, v int
	inter   [max][max]int
	visited [max]int
)

func dfs(w *bufio.Writer, v int) {
	Fprintf(w, "%d ", v)
	visited[v] = 1
	for i := 1; i <= n; i++ {
		if visited[i] == 1 || inter[v][i] == 0 {
			continue
		}
		dfs(w, i)
	}
}

func bfs(w *bufio.Writer, v int) {
	visited[v] = 0
	queue := []int{v}

	for len(queue) > 0 {
		front := queue[0]
		Fprintf(w, "%d ", front)
		queue = queue[1:]
		for i := 0; i < len(inter[v]); i++ {
			if inter[front][i] == 0 || visited[i] == 0 {
				continue
			}
			visited[i] = 0
			queue = append(queue, i)
		}
	}
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// n : 정점의 갯수
	// m : 간선의 갯수
	// v : 정점의 넘버
	var x, y int
	Fscanln(r, &n, &m, &v)
	for ; m > 0; m-- {
		Fscanln(r, &x, &y)
		inter[x][y] = 1
		inter[y][x] = inter[x][y]
	}
	dfs(w, v)
	Fprintln(w)
	bfs(w, v)
	Fprintln(w)
}
