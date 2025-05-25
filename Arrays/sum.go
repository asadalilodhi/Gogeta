package main

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(arrays_to_sum ...[]int) []int {
	var result []int

	for _, array := range arrays_to_sum {
		result = append(result, sum(array))
	}

	return result
}

func sumAlltails(arrays_to_sum ...[]int) []int {
	var result []int

	for _, array := range arrays_to_sum {
		if len(array) == 0 {
			result = append(result, 0)
		} else {
			tail := array[1:]
			result = append(result, sum(tail))
		}
	}

	return result
}