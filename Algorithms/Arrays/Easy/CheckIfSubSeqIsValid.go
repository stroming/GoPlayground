package main

// Input an array of integers and a sequence you would like to check if it's contained in the array
// Returns bool

func IsValidSubsequence(array []int, sequence []int) bool {
	j := 0
	for i := 0; i < len(array); i++ {
		if sequence[j] == array[i] {
			if j == len(sequence)-1 {
				return true
			}
			j++
		}
	}
	return false
}
