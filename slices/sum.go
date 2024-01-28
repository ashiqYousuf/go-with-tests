package slices

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(sliceOfIntSlices ...[]int) []int {
	var sums []int

	for _, numbers := range sliceOfIntSlices {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(sliceOfIntSlices ...[]int) []int {
	var sums []int

	for _, numbers := range sliceOfIntSlices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
