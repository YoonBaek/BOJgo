// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2447

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func iScan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func print(content string) {
	fmt.Fprint(wr, content)
}

func render(i int, j int, n int) {
	if i/n%3 == 1 && j/n%3 == 1 {
		print(" ")
	} else {
		if n/3 == 0 {
			print("*")
		} else {
			render(i, j, n/3)
		}
	}
}

func main() {
	defer wr.Flush()
	n := iScan()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			render(i, j, n)
		}
		print("\n")
	}
}
