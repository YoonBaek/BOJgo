// Solved By Github YoonBaek
// Q : https://www.acmicpc.net/problem/10757

package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()cv
	var a, b string
	Fscanf(r, "%s %s\n", &a, &b)
	v, v0, t := 0, 0, 0
	var sth, sth2, result string
	if len(a) >= len(b) {
		t = len(b)
	} else {
		t = len(a)
	}
	for ; t > 0; t-- {
		if (t-1)%18 == 0 {
			if v/int(math.Pow10(len(sth))) == 1 {
				v0++
			}
			sth = ""
			sth2 = ""
			result += string(v0)
		}
		sth += string(a[len(a)-t])
		sth2 += string(b[len(b)-t])
		r1, _ := strconv.Atoi(sth)
		r2, _ := strconv.Atoi(sth2)
		v = r1 + r2
		v0 = v
		Println(len(a), len(b), t, v/int(math.Pow10(len(sth))), r1, r2, result)
	}
	Fprintln(w, a, b, len(a))
}
