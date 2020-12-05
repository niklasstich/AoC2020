package AoC2020

func Day5_1() (highestChecksum int){
	lines, err := ReadAllLines("../inputs/05.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		var upperRow, lowerRow, leftColumn, rightColumn int = 127, 0, 0, 7
		for i := 0;i < 7; i++ {
			step := (upperRow - lowerRow + 1)/2
			if line[i] == 'F' {
				upperRow -= step
			} else if line[i] == 'B' {
				lowerRow += step
			} else {
				panic("did not expect symbol in input stream: " + string(line[i]))
			}
		}
		for i := 7; i < 9; i++ {
			step := (rightColumn - leftColumn + 1)/2
			if line[i] == 'L' {
				leftColumn += step
			} else if line[i] == 'R' {
				rightColumn -= step
			} else {
				panic("did not expect symbol in input stream: " + string(line[i]))
			}
		}
		// sanity check
		if upperRow != lowerRow || leftColumn != rightColumn {
			panic("sanity check failed")
		}
		checkSum := lowerRow * 8 + leftColumn
		if checkSum > highestChecksum {
			highestChecksum = checkSum
		}
	}
	return
}
