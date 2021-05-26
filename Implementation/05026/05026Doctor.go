// Solved by Github YoonBaek
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	var a, b, n int
	command := []string{}

	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	for i := 0; i < n; i++ {
		if sc.Scan() {
			if sc.Text()[0] >= '0' && sc.Text()[0] <= '9' {
				command = strings.Split(sc.Text(), "+")
				a, _ = strconv.Atoi(command[0])
				b, _ = strconv.Atoi(command[1])
				fmt.Fprintln(wr, a+b)
			} else {
				fmt.Fprintln(wr, "skipped")
			}
		}
	}
}
