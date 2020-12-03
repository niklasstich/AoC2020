package AoC2020



func Day3_1() (treeCount int) {
	field, err := ReadAllLines("../inputs/03.txt")
	if err != nil {
		panic(err)
	}
	return countTreesInPath(field, 3, 1)
}

func Day3_2() (treeCount int) {
	field, err := ReadAllLines("../inputs/03.txt")
	if err != nil {
		panic(err)
	}
	treeCount = 1
	treeCount*=countTreesInPath(field,1,1)
	treeCount*=countTreesInPath(field,3,1)
	treeCount*=countTreesInPath(field,5,1)
	treeCount*=countTreesInPath(field,7,1)
	treeCount*=countTreesInPath(field,1,2)
	return
}

func countTreesInPath(field []string, horiz, vert int) (treeCount int) {
	rowLen := len(field[0])
	horizPos := horiz
	for vertPos := vert; vertPos < len(field); vertPos+=vert {
		if field[vertPos][horizPos] == '#' {
			treeCount++
		}
		horizPos += horiz
		horizPos %= rowLen
	}
	return
}