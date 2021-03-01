package main

func LongestPeak(array []int) int {
	result := 0
	counter := 0
	// result < len(array) - i makes sure we don't
	// iterate needlessly if what is left of the array is smaller
	// than the longest peak we've found
	for i := 1; i < len(array)-1 && result < len(array)-i; i++ {
		// Looks for a peak
		if array[i-1] < array[i] && array[i] > array[i+1] {
			// Counts how long is the peak
			// on the left side of the peak
			for k := i; k > 0; k-- {
				if array[k] > array[k-1] {
					counter++
				} else {
					break
				}
			}
			// Counts how long is the peak
			// on the right side of the peak
			for j := i; j < len(array)-1; j++ {
				if array[j] > array[j+1] {
					counter++
				} else {
					// This will make sure we don't
					// iterate needlessly through
					// the right side of the current peak
					i = j
					break
				}
			}
			// Switch keeps track of the longest peak
			switch {
			case result < counter:
				result = counter
				counter = 0
			default:
				counter = 0
			}
		}
	}
	// If there is a peak
	// add 1 because the previous logic
	// does not count the peak itself
	if result > 0 {
		result++
	}
	return result
}
