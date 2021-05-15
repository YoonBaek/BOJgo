// Solved by Github YoonBaek

package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var (
	n, value, sum, mid, maxCnt, freq, freqCnt, midCnt int
	max, min                                          = -4000, 4000
	toggle                                            bool
)

func main() {
	sc, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	cnts := make([]int, 8001)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		value, _ = strconv.Atoi(sc.Text())
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
		cnts[value+4000]++
	}
	for idx, cnt := range cnts {
		realIdx := idx - 4000
		sum += cnt * (realIdx)
		midCnt += cnt

		if cnt > maxCnt {
			freqCnt = 0
			freq = realIdx
			maxCnt = cnt
		} else if cnt == maxCnt {
			freqCnt++
			if freqCnt == 1 {
				freq = realIdx
			}
		}
		if midCnt > (n-1)/2 {
			if !toggle {
				mid = realIdx
				toggle = true
			}
		}
	}
	Fprintf(w, "%.0f\n", float64(sum)/float64(n))
	Fprintln(w, mid)
	Fprintln(w, freq)
	Fprintln(w, max-min)
}
