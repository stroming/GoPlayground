package main

// Moves all instances of an inputed int to the back of the given array
// Returns the array
func MoveElementToEnd(array []int, toMove int) []int {
	for j, i := 0, 0; i < len(array); i++ {
		if array[i] != toMove {
			array[i], array[j] = array[j], array[i]
			j++
		}
	}
	return array
}
