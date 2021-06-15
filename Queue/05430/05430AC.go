// Solved by Github YoonBaek
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DQ struct {
	item []string
}

func (dq *DQ) spop() {
	dq.item = dq.item[:len(dq.item)-1]
}

func (dq *DQ) qpop() {
	dq.item = dq.item[1:]
}

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	num := 0
	text, _ := rd.ReadString('\n')
	text = strings.TrimSpace(text)
	num, _ = strconv.Atoi(text)
	return num
}

func main() {
	defer wr.Flush()
	var (
		T                 int
		p                 string
		dq                DQ
		Switch, errSwitch bool
	)
	T = scanInt()
	for i := 0; i < T; i++ {
		Switch, errSwitch = false, false
		text, _ := rd.ReadString('\n')
		p = strings.TrimSpace(text)

		scanInt()

		arrString, _ := rd.ReadString('\n')
		arrString = strings.TrimSpace(arrString)
		arrString = arrString[1 : len(arrString)-1]
		if len(arrString) != 0 {
			dq.item = strings.Split(arrString, ",")
		} else {
			dq.item = []string{}
		}

		for _, Func := range p {
			if Func == 'R' {
				Switch = !Switch
				continue
			}
			if Switch {
				if len(dq.item) > 0 {
					dq.spop()
				} else {
					errSwitch = true
					break
				}
			} else {
				if len(dq.item) > 0 {
					dq.qpop()
				} else {
					errSwitch = true
					break
				}
			}
		}
		if errSwitch {
			fmt.Fprintln(wr, "error")
			continue
		}
		fmt.Fprint(wr, "[")
		if Switch {
			for i := len(dq.item) - 1; i >= 0; i-- {
				if i != 0 {
					fmt.Fprint(wr, dq.item[i]+",")
				} else {
					fmt.Fprint(wr, dq.item[i])
				}
			}
		} else {
			fmt.Fprint(wr, strings.Join(dq.item, ","))
		}
		fmt.Fprint(wr, "]\n")
	}
}
