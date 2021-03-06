package main

import "fmt"

type hashMap map[string]int

type MinMaxStack struct {
	minMaxVals []hashMap
	stack      []int
}

func main() {
	x := NewMinMaxStack()
	x.Push(5)
	fmt.Println(x.Peek())
	fmt.Println(x.GetMin())
	fmt.Println(x.GetMax())
	x.Push(7)
	fmt.Println(x.Peek())
	fmt.Println(x.GetMin())
	fmt.Println(x.GetMax())
	x.Push(2)
	fmt.Println(x.Peek())
	fmt.Println(x.GetMin())
	fmt.Println(x.GetMax())
	x.Pop()
	fmt.Println(x.Peek())
	fmt.Println(x.GetMin())
	fmt.Println(x.GetMax())
	x.Pop()
	fmt.Println(x.Peek())
	fmt.Println(x.GetMin())
	fmt.Println(x.GetMax())
}

func NewMinMaxStack() *MinMaxStack {
	var minMaxStack *MinMaxStack
	minMaxStack = new(MinMaxStack)
	minMaxStack.minMaxVals = []hashMap{
		map[string]int{
			// is the smallest possible integer in int32
			"max": -2147483648,
			"min": 2147483648,
		},
	}

	return minMaxStack
}

func (stack *MinMaxStack) Peek() int {
	if len(stack.stack) == 0 {
		return -1
	}
	lastIndex := len(stack.stack) - 1
	return stack.stack[lastIndex]
}

func (stack *MinMaxStack) Pop() int {
	if len(stack.stack) == 0 {
		return -1
	}
	lastIndex := len(stack.stack) - 1
	lastValue := stack.stack[lastIndex]

	// delete items from the stack
	stack.stack = stack.stack[:len(stack.stack)-1]
	stack.minMaxVals = stack.minMaxVals[:len(stack.minMaxVals)-1]

	return lastValue
}

func (stack *MinMaxStack) Push(number int) {
	stack.stack = append(stack.stack, number)
	currMinValue := stack.minMaxVals[len(stack.minMaxVals)-1]["min"]
	currMaxValue := stack.minMaxVals[len(stack.minMaxVals)-1]["max"]
	m := map[string]int{
		"min": currMinValue,
		"max": currMaxValue,
	}
	stack.minMaxVals = append(stack.minMaxVals, m)
	if currMinValue > number {
		stack.minMaxVals[len(stack.minMaxVals)-1]["min"] = number
	}
	if currMaxValue < number {
		stack.minMaxVals[len(stack.minMaxVals)-1]["max"] = number
	}
}

func (stack *MinMaxStack) GetMin() int {
	return stack.minMaxVals[len(stack.minMaxVals)-1]["min"]
}

func (stack *MinMaxStack) GetMax() int {
	return stack.minMaxVals[len(stack.minMaxVals)-1]["max"]
}
