package main

import "math"

// Find the smallest difference between two numbers in two arrays
// return it as an array of 2 elements
func SmallestDifference(array1, array2 []int) []int {
	numOne := array1[0]
	numTwo := array2[0]
	absDiff := math.Abs(float64(numOne) - float64(numTwo))

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if math.Abs(float64(array1[i])-float64(array2[j])) < absDiff {
				absDiff = math.Abs(float64(array1[i]) - float64(array2[j]))
				numOne = array1[i]
				numTwo = array2[j]
			}
		}
	}
	result := []int{numOne, numTwo}
	return result
}
