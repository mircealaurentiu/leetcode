package source

func SingleNumber(nums []int) int {

	freqNums := make(map[int]int)

	for _, num := range nums {
		freqNums[num]++
	}

	for num, count := range freqNums {
		//fmt.Println(num, count)
		if count == 1 {
			return num
		}
	}

	return -1
}
