// Solved by Github YoonBae

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func main() {
	defer wr.Flush()
	sc.Scan()
	strs := strings.Split(sc.Text(), " ")
	a, b := strs[0], strs[1]
	overDigit := false
	maxLen, minLen := 0, 0
	ansStr := ""
	if len(a) > len(b) {
		maxLen = len(a)
		minLen = len(b)
		ansStr = a[0 : maxLen-minLen]
	} else if len(a) == len(b) {
		maxLen = len(a)
		minLen = len(b)
		ansStr = ""
	} else {
		maxLen = len(b)
		minLen = len(a)
		ansStr = b[0 : maxLen-minLen]
	}

	ansStrs := make([]string, minLen)

	for i := 0; i < minLen; i++ {
		digitSum := int(a[len(a)-1-i]-'0') + int(b[len(b)-1-i]-'0')
		if overDigit {
			digitSum++
		}
		if digitSum >= 10 {
			overDigit = true
			digitSum %= 10
		} else {
			overDigit = false
		}
		ansStrs[minLen-i-1] = strconv.Itoa(digitSum)
	}
	ansRest := ""
	for _, str := range ansStrs {
		ansRest += str
	}
	frontStrs := []string{}
	if overDigit {
		for i := len(ansStr) - 1; i >= 0; i-- {
			digitSum := int(ansStr[i]-'0') + 1
			if digitSum >= 10 {
				digitSum %= 10
			} else {
				overDigit = false
			}
			frontStrs = append(frontStrs, strconv.Itoa(digitSum))
			if !overDigit {
				break
			}
		}
	}
	ans := ansStr[:len(ansStr)-len(frontStrs)]
	tmp := ""

	// reverse
	for i := len(frontStrs) - 1; i >= 0; i-- {
		tmp += frontStrs[i]
	}
	ans += (tmp + ansRest)
	if overDigit {
		ans = "1" + ans
	}
	fmt.Fprintln(wr, ans)
}
