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
