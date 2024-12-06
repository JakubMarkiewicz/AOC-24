package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func B(input []string) {
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

	for _, l := range lines {
		valid := true
		for i := 1; i < len(l); i++ {

			v := l[i]
			mV := m[v]

			for j := 0; j < i; j++ {
				if slices.Contains(mV, l[j]) {
					valid = false
					prev := l[j]
					l[j] = v
					l[i] = prev
					i = 1
					break
				}
			}
		}
		if !valid {
			fmt.Println(buildInput(l))
			m := l[len(l)/2]
			i, _ := strconv.Atoi(m)
			res += i
		}
	}

	fmt.Println(res)

}

func buildInput(v []string) string {
	return strings.Join(v, ",")
}
