// Solved by Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	var word, prev string
	var length = make([][]string, 50)
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}

	for i := 0; i < n; i++ {
		if sc.Scan() {
			word = sc.Text()
		}
		length[len(word)-1] = append(length[len(word)-1], word)
	}
	for _, words := range length {
		for i := 0; i < len(words)-1; i++ {
			for j := i + 1; j < len(words); j++ {
				if words[i] > words[j] {
					tmp := words[i]
					words[i] = words[j]
					words[j] = tmp
				}
			}
		}
		for _, word := range words {
			if prev != word {
				fmt.Fprintln(wr, word)
			}
			fmt.Fprintln(wr, word)
			prev = word
		}
	}
}
