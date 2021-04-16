// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2869
// 낮에 a meter 올라가고, 밤에 b meter 떨어지는 달팽이가
// v 만큼 올라가는 데 걸리는 시간을 구하는 문제
// n(a - b) + a >= v 일 때 n-1의 값을 구하면 될 듯
// 사칙 연산으로 만들어야 빨리 풀리나?

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a, b, v, result int
	Fscanf(r, "%d %d %d\n", &a, &b, &v)
	if (v-a)%(a-b) != 0 {
		result = (v-a)/(a-b) + 2
	} else {
		result = (v-a)/(a-b) + 1
	}
	Println(result)
}
