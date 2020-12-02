package AoC2020

import (
	"strconv"
	"strings"
)

func Day2_1() (valid int) {
	lines, err := ReadAllLines("../inputs/02.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		pwd, requiredSymbol, minCount, maxCount := ParseLine(line)
		if CheckValidity1(pwd, requiredSymbol, minCount, maxCount){
			valid++
		}
	}
	return
}

func Day2_2() (valid int) {
	lines, err := ReadAllLines("../inputs/02.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		pwd, requiredSymbol, pos1, pos2 := ParseLine(line)
		if CheckValidity2(pwd, requiredSymbol, pos1, pos2) {
			valid++
		}
	}
	return
}



func ParseLine(line string) (pwd, symbol string, num1, num2 int) {
	splits := strings.Split(line, " ")
	pwd = splits[2]
	symbol = strings.Trim(splits[1], ":")
	counts := strings.Split(splits[0], "-")
	num1, err := strconv.Atoi(counts[0])
	if err != nil {
		panic(err)
	}
	num2, err = strconv.Atoi(counts[1])
	if err != nil {
		panic(err)
	}
	return
}

func CheckValidity1(password, requiredSymbol string, minCount, maxCount int) bool {
	realCount := strings.Count(password, requiredSymbol)
	return realCount >= minCount && realCount <= maxCount
}

func CheckValidity2(pwd string, symbol string, pos1 int, pos2 int) bool {
	if pwd[pos1-1] == symbol[0] {
		return pwd[pos2-1] != symbol[0]
	}
	return pwd[pos2-1] == symbol[0]
}