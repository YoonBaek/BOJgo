// Solved by Github YoonBaek

package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(r)
	defer w.Flush()

	var n, elem int
	Fscanln(r, &n)
	cnts := make([]int, 10001)

	for i := 0; i < n; i++ {
		sc.Scan()
		elem, _ = strconv.Atoi(sc.Text())
		cnts[elem]++
	}
	for idx, cnt := range cnts {
		s := strconv.Itoa(idx) + "\n"
		if cnt > 0 {
			Fprint(w, strings.Repeat(s, cnt))
		}
	}
}
