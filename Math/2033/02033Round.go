// Solved by Github YoonBaek
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	var n string

	if sc.Scan() {
		n = sc.Text()
	}

	var add byte
	round := make([]string, len(n)+1)

	for i := len(n) - 1; i >= 0; i-- {
		round[i+1] = string(n[i] + add)
		if round[i+1] >= "5" {
			add = 1
		} else {
			add = 0
		}
		if i != 0 {
			round[i+1] = "0"
		} else {
			if round[i+1] == ":" {
				round[i+1] = "10"
			}
		}
	}
	fmt.Fprintf(wr, "%s\n", strings.Join(round, ""))
}
