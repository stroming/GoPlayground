package main

// Given an array of ints where every element is equal to the
// multiplication of every other element in the array
// Returns an array of the same size
func ArrayOfProducts(array []int) []int {
	result := 1
	resultArr := []int{}
	zeroCounter := 0
	// For multiplies every number by the next number
	// and keeps in a variable
	for i := 0; i < len(array); i++ {
		switch {
		case zeroCounter > 1:
			// If we find more than 1 zero we stop iterating
			// because everything will be a 0
			break
		case array[i] == 0:
			// keeps track of how many zeros we have
			zeroCounter++
		default:
			result = result * array[i]
		}
	}
	// For creates our resultArray and fills it up
	for i := 0; i < len(array); i++ {
		switch {
		case zeroCounter > 1:
			// If zeros > 1 the whole array is filled with 0s
			resultArr = append(resultArr, 0)
		case array[i] == 0:
			// on the spot where there's a zero in our
			// input array we put the multiplication of all the numbers
			// except the zero
			resultArr = append(resultArr, result)
		case zeroCounter == 1:
			// If there is a zero in our input array
			// make every spot in our output array = 0
			resultArr = append(resultArr, 0)
		default:
			resultArr = append(resultArr, result/array[i])
		}
	}
	return resultArr
}
