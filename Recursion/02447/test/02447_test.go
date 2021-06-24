package star10

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var (
	wr = bufio.NewWriter(os.Stdout)
)

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

func Test1(t *testing.T) {
	n := 243 * 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			render(i, j, n)
		}
		print("\n")
	}
}
