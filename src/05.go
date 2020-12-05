package AoC2020

func Day5_1() (highestChecksum int) {
	lines, err := ReadAllLines("../inputs/05.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		seat := getSeatNumber(line)
		if seat > highestChecksum {
			highestChecksum = seat
		}
	}
	return
}

func Day5_2() (checksum int) {
	lines, err := ReadAllLines("../inputs/05.txt")
	if len(lines) == 0 {
		panic("empty input")
	}
	if err != nil {
		panic(err)
	}
	min, max := 1024, -1
	for _, line := range lines {
		num := getSeatNumber(line)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		checksum ^= num
	}
	for i := 0; i < min; i++ {
		checksum ^= i
	}
	for i := max + 1; i < 1024; i++ {
		checksum ^= i
	}
	return
}

func getSeatNumber(seatStr string) int {
	var row, column = 0, 0
	for i := 0; i < 7; i++ {
		if seatStr[i] == 'F' {
			row = row<<1 | 0
		} else if seatStr[i] == 'B' {
			row = row<<1 | 1
		} else {
			panic("did not expect symbol in input stream: " + string(seatStr[i]))
		}
	}
	for i := 7; i < 10; i++ {
		if seatStr[i] == 'L' {
			column = column<<1 | 0
		} else if seatStr[i] == 'R' {
			column = column<<1 | 1
		} else {
			panic("did not expect symbol in input stream: " + string(seatStr[i]))
		}
	}
	// sanity check
	return row*8 + column
}
