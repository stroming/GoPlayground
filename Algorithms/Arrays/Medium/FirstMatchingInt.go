package main

// Given an array of ints find the first one that is duplicated
// Return -1 if no such int exist in the given array
func FirstDuplicateValue(array []int) int {
	m := make(map[int]int)

	for i := 0; i < len(array); i++ {
		m[array[i]]++

		if m[array[i]] == 2 {
			return array[i]
		}
	}
	return -1
}
