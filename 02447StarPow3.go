// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2447

package main

import (
	. "fmt"
	"strings"
)

func makeSquare(pat string) {
	Println(pat + pat + pat + "\n" + pat + strings.Replace(pat, "*", " ", -1) + pat + pat + pat + pat)
}

func main() {
	makeSquare("***\n* *\n***")
}
