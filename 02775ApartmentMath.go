// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2775
// 아파트 문제
// k 층의 n호에 살고 있는 사람 구하기
// 0층 1 2 3 4 5 6 ... i
//  1 1 1 1 1
// 1층 1 3 6 10 15 21 ... i
//  2 3 4 5 6
// 2층 1 4 10 20 35 56
//  3 6 10 15 21

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func selfCount(k int, n int) (result int) {
	if k == 1 {
		for j := 1; j <= n; j++ {
			result += j
		}
	} else {
		for i := 1; i <= n; i++ {
			result += selfCount(k-1, i)
		}
	}
	return result
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var testCase int
	Fscanln(r, &testCase)
	defer w.Flush()

	for i := 0; i < testCase; i++ {
		kFloor, nRoomNumber := 0, 0
		Fscanf(r, "%d\n%d\n", &kFloor, &nRoomNumber)
		Fprintln(w, selfCount(kFloor, nRoomNumber))
	}
}
