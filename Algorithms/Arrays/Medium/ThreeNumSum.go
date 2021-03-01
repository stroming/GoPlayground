package main

// I first sort them using QuickSort method
func quickSort(arr []int, left int, right int) {
	if left < right {
		pi := partition(arr, left, right)
		quickSort(arr, left, pi-1)
		quickSort(arr, pi+1, right)
	}
}

func partition(arr []int, low int, high int) int {
	// pivot (Element to be placed at right position)
	pivot := arr[high]

	i := (low - 1) // Index of smaller element

	for j := low; j <= high-1; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++ // increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}

// Inputs an array of ints and a target sum
// Returns all triplets that could make that sum in a 2D array
func ThreeNumberSum(s []int, sum int) [][]int {
	quickSort(s, 0, len(s)-1)
	multiSl := [][]int{}
	left := 0
	right := len(s) - 1
	pivot := 1
	for left < pivot {
		if pivot >= right {
			break
		}
		if (s[left] + s[pivot] + s[right]) > sum {
			right--
		}
		if (s[left] + s[pivot] + s[right]) < sum {
			pivot++
		}
		if (s[left]+s[pivot]+s[right]) == sum && left != pivot && pivot != right {
			multiSl = append(multiSl, []int{s[left], s[pivot], s[right]})
			pivot++
		}
		if pivot == right {
			left++
			pivot = left + 1
			right = len(s) - 1
		}
	}
	return multiSl
}
