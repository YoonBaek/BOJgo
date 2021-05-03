// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/14889
// 축구 팀 시너지차이 최솟값 찾는 문제

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func getTable(r *bufio.Reader) [][]int {
	table := make([][]int, n)

	for i := 0; i < n; i++ {
		row := make([]int, n)

		for j := 0; j < n; j++ {
			Fscanf(r, "%d ", &row[j])
		}
		table[i] = row
	}
	return table
}

// func solve(table [][]int) {
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if i > j {
// 				synergy1 := table[i][j] + table[j][i]
// 				if
// 				Println(synergy1, synergy2)
// 			}
// 		}
// 	}
// }

func dfs(visited []bool, s int, k int) {
	if k > n/2 {
		Println(visited)
		Println("done")
	} else {
		for i := s; i < n; i++ {
			if !visited[i] {
				visited[i] = true
				dfs(visited, i+1, k+1)
			}
		}
	}
}

var (
	n int
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	Fscanln(r, &n)
	visited := make([]bool, n)
	table := getTable(r)
	start := make([]int, n/2)
	Fprintln(w, table, len(start))
	dfs(visited, 0, 0)
	// solve(n, table)
}
