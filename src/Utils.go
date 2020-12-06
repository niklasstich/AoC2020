package AoC2020

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

//bitwise complement
const MaxUint = ^uint(0)
const MinUint = 0

//2's complement
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func ReadFileAsInts(filename string) ([]int, error) {
	var nums []int
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	for err != io.EOF {
		line, err := reader.ReadString(0xa)
		if err != nil && err != io.EOF {
			return nil, err
		}
		line = strings.TrimSpace(line)
		//skip empty lines
		if len(line) == 0 {
			break
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func ReadAllLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var sarr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//i don't think this is needed
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		sarr = append(sarr, line)
	}
	return sarr, nil
}

func ReadLinesSeperatedByEmptyLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var sarr []string
	var buffer string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// if line is an empty line, push buffer into sarr and clear it
		if line == "" {
			sarr = append(sarr, buffer)
			buffer = ""
			continue
		}
		// else just append line to buffer
		if buffer != "" {
			buffer += " "
		}
		buffer += line
	}
	if buffer != "" {
		sarr = append(sarr, buffer)
	}
	return sarr, nil
}

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func SumInts(slice []int) (retval int, overflow bool) {
	for _, i := range slice {
		if retval+i < retval {
			overflow = true
		}
		retval += i
	}
	return
}

//Gets distinct subset of slice with minimal allocations (1 map with keys only)
func DistinctSlice(s []rune) (distinct []rune) {
	seen := make(map[rune]struct{}, len(s)) //map for uniqueness
	j := 0                                  //insertion index
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue //value already in map which means its also in s
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
