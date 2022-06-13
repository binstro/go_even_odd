package goevenodd

import "strings"

func CheckEvenOdd(nums ...int) string {
	var result = make([]string, len(nums))

	for index, num := range nums {
		if num%2 == 0 {
			result[index] = "genap"
		} else {
			result[index] = "ganjil"
		}
	}

	return strings.Join(result, ", ")
}
