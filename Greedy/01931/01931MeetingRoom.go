// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Meeting struct {
	Start int
	End   int
}
type Table []Meeting

func (t Table) Len() int {
	return len(t)
}

func (t Table) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Table) Less(i, j int) bool {
	if t[i].End == t[j].End {
		return t[i].Start < t[j].End
	} else {
		return t[i].End < t[j].End
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var N int
	var meeting Meeting

	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}

	tables := make(Table, N)

	for i := 0; i < N; i++ {
		if sc.Scan() {
			field := strings.Fields(sc.Text())
			meeting.Start, _ = strconv.Atoi(field[0])
			meeting.End, _ = strconv.Atoi(field[1])
			tables[i] = meeting
		}
	}
	sort.Sort(tables)

	answer, prevEnd := 0, 0
	for i := 0; i < N; i++ {
		if tables[i].Start >= prevEnd {
			answer++
			prevEnd = tables[i].End
		}
	}
	fmt.Fprintln(wr, answer)
}
