// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/3009
// 3가지 꼭지점 좌표를 참고해서 마지막 꼭지점을 찾는 문제
// XOR 연산을 이용해서 풀었다.
// ex) 30 = 11110
//     10 = 01010
//    xor = 10100
//     10 = 01010
//    xor = 11110
// xor을 두번 하면 하나만 있는 값을 반환
package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	xs := make([]int, 3)
	ys := make([]int, 3)
	for i := 0; i < 3; i++ {
		var x, y int
		Fscanf(r, "%d %d\n", &x, &y)
		xs[i] = x
		ys[i] = y
	}
	Fprintln(w, xs[0]^xs[1]^xs[2], ys[0]^ys[1]^ys[2])
}
