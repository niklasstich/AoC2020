package AoC2020

import (
	"sort"
	"strings"
)

func Day6_1() (answers int) {
	lines, err := ReadLinesSeperatedByEmptyLines("../inputs/06.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		u := []rune(strings.Replace(line, " ", "", -1))
		sort.Sort(sortRunes(u))
		u = DistinctSlice(u)
		answers += len(u)
	}
	return
}
