package main

// Given an array of integers return if the given array is monotinic
// Return Bool
func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	for x := 0; x < len(array)-1; {
		switch {
		case array[x] < array[x+1]:
			for i := x; i < len(array)-1; i++ {
				if array[i] > array[i+1] {
					return false
				}
			}
			return true
		case array[x] > array[x+1]:
			for i := x; i < len(array)-1; i++ {
				if array[i] < array[i+1] {
					return false
				}
			}
			return true
		default:
			x++
		}
	}
	return true
}
