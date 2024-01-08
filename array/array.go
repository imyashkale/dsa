package array

import "math"

type Array struct {
	elements []int
	len      int
}

// NewArray Create the array
func New(size int) *Array {
	return &Array{
		elements: make([]int, size),
		len:      size,
	}
}

func (a *Array) With(arr []int) *Array {
	a.elements = arr
	return a
}

// LinearSearch
func (a Array) LinearSearch(target int) int {
	for i := 0; i < a.Len(); i++ {
		if target == a.elements[i] {
			return i
		}
	}
	return -1
}

// Len Returns the len of the array
func (a Array) Len() int {
	return a.len
}

// Largest Find the largest Number in array/ Maximum No
func (a Array) Largest() int {
	largest := math.MinInt
	for i := 0; i < a.Len(); i++ {
		if a.elements[i] > largest {
			largest = a.elements[i]
		}
	}
	return largest
}

// Lowest Find the lowest Element in the array
func (a Array) Lowest() int {
	lowest := math.MaxInt
	for i := 0; i < a.Len(); i++ {
		if a.elements[i] < lowest {
			lowest = a.elements[i]
		}
	}
	return lowest
}
