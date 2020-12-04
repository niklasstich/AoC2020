package AoC2020

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

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
		if line=="" {
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
		if line=="" {
			sarr = append(sarr, buffer)
			buffer = ""
			continue
		}
		// else just append line to buffer
		if buffer != ""{
			buffer += " "
		}
		buffer += line
	}
	return sarr, nil
}

func Contains(slice []string, s string) bool{
	for _, item := range slice {
		if s == item {
			return true
		}
	}
	return false
}