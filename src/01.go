package AoC2020

type tuple struct {
	a int
	b int
}

func Day1_1() int{
	nums, err := ReadFileAsInts("../inputs/01.txt")
	if err != nil {
		panic(err)
	}
	//create buffered channel and write every number into it once
	work := make(chan int, len(nums))
	m := make(map[int]bool)
	for _, n := range nums {
		work <- n
		m[n] = true
	}
	close(work)
	//according to the problem its safe to assume only one pair of numbers make up the number 2020
	res := make(chan int, 1)
	//to parallelize, run as many goroutines as we have logical processors
	//for i := 0; i<runtime.NumCPU(); i++ {
		go func() {
			for w := range work {
				if m[2020-w] {
					res <- w
				}
			}
		}()
	//}
	correct := <- res
	return correct * (2020-correct)
}

func Day1_2() int{
	nums, err := ReadFileAsInts("../inputs/01.txt")
	if err != nil {
		panic(err)
	}
	//create tuples
	//work := make(chan tuple, ((len(nums)-1)*len(nums))/2)
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}
	for a := len(nums)-1; a>=0; a-- {
		for b := 0; b < a; b++ {
			if m[2020-nums[a]-nums[b]] {
				return (2020-nums[a]-nums[b]) * nums[a] * nums[b]
			}
		}
	}
	return -1
}