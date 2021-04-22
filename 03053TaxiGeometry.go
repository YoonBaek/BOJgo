// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/3053
// 맨해튼 거리로 원 넓이 풀으라는 문제 같음.
// 맨해튼 거리에서는 거리가 같은 점들이 마름모 꼴로 형성
// 즉 원의 넓이와 마름모의 넓이를 각각 구하면 되는 문제.

package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var radius float64
	Fscanln(r, &radius)
	Fprintf(w, "%.6f\n", radius*radius*math.Pi)
	Fprintf(w, "%.6f\n", radius*radius*2)
}
