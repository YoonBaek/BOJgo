// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func iScan() (first, second int) {
	fmt.Fscanf(rd, "%d %d\n", &first, &second)
	return first, second
}

func print(content string) {
	fmt.Fprintln(wr, content)
}

func main() {
	defer wr.Flush()
	for true {
		first, second := iScan()
		if first == 0 && second == 0 {
			break
		}

		if first%second == 0 {
			print("multiple")
		} else if second%first == 0 {
			print("factor")
		} else {
			print("neither")
		}
	}
}
