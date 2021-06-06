// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 큐 타입 정의
type Queue struct {
	item []int
	size int
}

// 큐 메서드 정의
func (q *Queue) Push(n int) {
	// 동적 추가
	q.item = append(q.item, n)
	// 사이즈 체크
	q.size++
}

// 큐 메서드 정의
func (q *Queue) Pop() int {
	pop := q.item[0]
	// 큐에서 제거
	q.item = q.item[1:]
	// 사이즈 체크
	q.size--
	return pop
}

var (
	n, k int
	sc   *bufio.Scanner
	wr   *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	if sc.Scan() {
		fields := strings.Fields(sc.Text())
		n, _ = strconv.Atoi(fields[0])
		k, _ = strconv.Atoi(fields[1])
	}
	// 배열 선언
	O := make([]int, n)
	for i := 0; i < n; i++ {
		O[i] = i + 1
	}
	// 큐 선언
	q := Queue{item: O}
	q.size = n

	// 문제 format 맞춰주기
	fmt.Fprint(wr, "<")
	// 요세푸스 코드
	// q가 비어있게 되면 종료
	for q.size > 0 {
		// k 전까지는 pop - push 반복
		// q.size 유지
		for i := 0; i < k-1; i++ {
			q.Push(q.Pop())
		}
		// k 번째에서 pop하고 출력
		// q.size 감소
		if q.size == 1 {
			fmt.Fprintf(wr, "%d", q.Pop())
		} else {
			fmt.Fprintf(wr, "%d, ", q.Pop())
		}
	}
	// 문제 format 맞춰주기
	fmt.Fprint(wr, ">\n")
}
