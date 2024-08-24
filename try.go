package main

import (
	"fmt"
	"sort"
	"strings"
)

type Try map[int]*Permutations

func (t Try) String() string {
	res := []string{}
	for i := range t {
		res = append(res, fmt.Sprintf("%d: %d\n", i, len(*t[i])))
	}
	sort.Strings(res)
	return fmt.Sprintln(strings.Join(res[:], ""))
}
