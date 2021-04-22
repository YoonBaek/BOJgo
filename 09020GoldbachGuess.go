// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/9020
// 골드바흐의 추측을 구하라고 함
// 골드바흐의 추측이란 4이상의 모든 짝수는 소수의 합으로 표현할 수 있는데,
// 예를 들어 8 = 3 + 5
// 한 짝수에 여러 골드바흐의 수가 들어있을 수도 있음
// 이 경우에는 두 소수간의 차이가 가장 적은 골드바흐 수를 구하라는 문제임

// 풀이 내꺼
package main

import (
	"bufio"
	. "fmt"
	"os"
)

// 이거 소수 구하는 함수
// 인풋으로 넣은 골드바흐 수 값이 소수가 맞는지 확인
// cnt = 0이면 소수 (원래는 2가 소수이나 편의상 1과 자신의 크기 두개는 빼서 0)
func checker(gb int) (cnt int) {
	factor, cnt := 2, 0
	for factor*factor <= gb {
		if gb%factor == 0 {
			gb /= factor
			cnt++
		} else {
			factor++
		}
	}
	return cnt
}

// 이게 메인 파트 고랭은 메인함수에 넣어야 작동함
func main() {
	// 버퍼
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	// 코드 끝나면 버퍼 메모리 관리해줌
	defer w.Flush()
	// 여기서부터 시작
	// case 갯수 t가 주어짐. t만큼 반복
	var t int
	Fscanln(r, &t)
	for ; t > 0; t-- {
		var n int
		Fscanln(r, &n)
		gb1, gb2 := n/2, n/2
		cnt1, cnt2 := checker(gb1), checker(gb2)
		for cnt1 != 0 || cnt2 != 0 {
			gb1--
			gb2++
			cnt1 = checker(gb1)
			cnt2 = checker(gb2)
		}
		Fprintln(w, gb1, gb2)
	}
}
