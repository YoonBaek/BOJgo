package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func main() {
	defer wr.Flush()

	var A, B string
	if sc.Scan() {
		strs := strings.Fields(sc.Text())
		A = strs[0]
		B = strs[1]
	}

	numA := len(A)
	numB := len(B)

	if numB > numA || (numB == numA) {
		tmp := A
		A = B
		B = tmp

		tmp2 := numA
		numA = numB
		numB = tmp2
	}

	numDiff := numA - numB
	// 배열로 풀이
	var ans []string = make([]string, numA+1)

	// 십의 자리수 조정
	upper := 0
	// 바로 뒤 루프에 십의 자리수 연산 해줘야 하기때문에 역순 진행
	for i := numA - 1; i >= 0; i-- {
		sum := 0
		if i-numDiff >= 0 {
			sum = int(A[i]+B[i-numDiff]-'0'-'0') + upper
		} else {
			sum = int(A[i]-'0') + upper
		}

		if sum >= 10 {
			upper = 1
			ans[i+1] = strconv.Itoa(sum % 10)
		} else {
			upper = 0
			ans[i+1] = strconv.Itoa(sum)
		}
	}
	// 잔여 upper 처리
	if upper == 1 {
		ans[0] = "1"
	}
	for i := 0; i < numA+1; i++ {
		fmt.Fprint(wr, ans[i])
	}
	fmt.Fprintln(wr)
}
