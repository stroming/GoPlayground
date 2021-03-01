package main

// Given a square matrix of type int return the matrix's values in spiral order
// Return an array of the matrix's values in spiral order
func SpiralTraverse(sl [][]int) []int {
	result := []int{}
	eRow := len(sl) - 1
	eCol := len(sl[0]) - 1
	for sCol := 0; sCol <= eCol; sCol++ {
		result = append(result, sl[0][sCol])
	}
	// Delete the empty row
	sl = sl[1:]
	for sRow := 0; sRow < eRow; sRow++ {
		result = append(result, sl[sRow][eCol])
		sl[sRow] = sl[sRow][:eCol]
	}
	eRow = len(sl) - 1
	if len(sl) == 0 {
		return result
	}
	for eCol = len(sl[0]) - 1; eCol >= 0; eCol-- {
		result = append(result, sl[eRow][eCol])
	}
	// Delete the last empty row
	sl = sl[:eRow]
	for eRow = len(sl) - 1; eRow >= 0 && len(sl[eRow]) > 0; eRow-- {
		result = append(result, sl[eRow][0])
		sl[eRow] = sl[eRow][1:]
	}
	// Checks if the Rows of are empty
	for i := len(sl) - 1; i >= 0; i-- {
		if len(sl[i]) == 0 {
			sl = sl[:i]
		}
	}
	if len(sl) != 0 {
		result = append(result, SpiralTraverse(sl)...)
	}
	return result
}

// Example
// matrix = [][]int {
// 	{1,2,3,4},
// 	{12,13,14,5},
// 	{11,16,15,6},
// 	{10,9,8,7},
// }

// returns array{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
