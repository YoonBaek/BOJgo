// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/9020
// 골드바흐의 추측을 구하라고 함
// 골드바흐의 추측이란 4이상의 모든 짝수는 소수의 합으로 표현할 수 있는데,
// 예를 들어 8 = 3 + 5
// 한 짝수에 여러 골드바흐의 수가 들어있을 수도 있음
// 이 경우에는 두 소수간의 차이가 가장 적은 골드바흐 수를 구하라는 문제임

// Solve1
package main

import (
	"bufio"
	. "fmt"
	"os"
)

// 소수 구하는 함수
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
	var primes [10001]int
	Fscanln(r, &t)
	for i := 2; i <= 10000; i++ {
		if primes[i] == 1 || primes[i] == -1 {
			continue
		}
		if checker(i) == 0 {
			primes[i] = 1
			n := 2
			for n*i <= 10000 {
				primes[n*i] = -1
				n++
			}
		}
	}
	for ; t > 0; t-- {
		//n의 골드바흐 수를 구해야 함
		var n int
		Fscanln(r, &n)
		// 일단 가장 차이가 적어야 하므로 2분의 1을 각각 할당
		gb1, gb2 := n/2, n/2
		// 소수인지 확인
		// 둘 중 하나라도 소수가 아니면 계속 반복문 돌아감
		// golang은 while이 없이 for에 조건식을 주면 while처럼 작동
		// 둘 다 소수일때까지만 반복 돌림
		for primes[gb1] != 1 || primes[gb2] != 1 {
			gb1--
			gb2++
		}
		Fprintln(w, gb1, gb2)
	}
}
