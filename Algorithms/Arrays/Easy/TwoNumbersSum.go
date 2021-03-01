package main

// Input an array of integers and target sum
// if any two numbers summed up make the targeted int
// return it otherwise return an empty array

func TwoNumberSum(array []int, target int) []int {

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if (array[i] + array[j]) == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}
