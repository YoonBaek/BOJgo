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
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var input string
	intSort := make([]int, 10)
	fmt.Fscanln(rd, &input)

	for _, s := range input {
		intSort[int(s)-48]++
	}

	ans := ""
	for i := 9; i >= 0; i-- {
		ans += strings.Repeat(strconv.Itoa(i), intSort[i])
	}
	fmt.Fprintln(wr, ans)
}
