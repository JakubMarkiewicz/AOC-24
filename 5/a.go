package main

import (
	"fmt"
	"strconv"
	"strings"
)

func A(input []string) {
	m := make(map[string][]string)
	lines := [][]string{}

	isDef := true

	for _, v := range input {
		if v == "" {
			isDef = false
			continue
		}

		if isDef {
			d := strings.Split(v, "|")
			mD, prs := m[d[0]]
			if prs {
				m[d[0]] = append(mD, d[1])
			} else {
				m[d[0]] = []string{d[1]}
			}
		} else {
			line := strings.Split(v, ",")
			lines = append(lines, line)
		}
	}

	res := 0

	for i, l := range lines {
		visited := make(map[string]bool)
		valid := true
		for _, lV := range l {
			if !valid {
				break
			}
			// check
			for _, c := range m[lV] {
				_, prs := visited[c]
				if prs {
					valid = false
					break
				}
			}
			visited[lV] = true
		}
		if valid {
			m := l[len(l)/2]
			i, _ := strconv.Atoi(m)
			res += i
		} else {
			fmt.Printf("%d is invalid!!\ndata:%v\n", i, l)
		}
	}

	fmt.Println(res)
}
