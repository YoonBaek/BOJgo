// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/10250
// 호텔 방 assign 문제
// 엘레베이터에서 가까운 방대로 채우고, 층 수는 아래층을 우선 선호
//
package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, wr := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var t, h, w, n int
	Fscanln(r, &t)
	for i := 0; i < t; i++ {
		Fscanf(r, "%d %d %d\n", &h, &w, &n)
		if n%h == 0 {
			Fprintf(wr, "%d%02d\n", h, n/h)
		} else {
			Fprintf(wr, "%d%02d\n", n%h, n/h+1)
		}
		wr.Flush()
	}
}
