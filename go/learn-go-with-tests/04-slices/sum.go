package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sum []int

	for _, numslice := range numbers {
		sum = append(sum, Sum(numslice))
	}

	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sum []int

	for _, numslice := range numbers {
		if len(numslice) > 0 {
			sum = append(sum, Sum(numslice[1:]))
		} else {
			sum = append(sum, 0)
		}
	}

	return sum
}
