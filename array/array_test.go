package array_test

import (
	"testing"

	"github.com/imyashkale/dsa/array"
	"github.com/magiconair/properties/assert"
)

func TestArray(t *testing.T) {
	a := array.New(0)
	assert.Equal(t, a.Len(), 0)
}

func TestLinearSearchFound(t *testing.T) {
	a := array.New(10).With([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	targetIndex := a.LinearSearch(2)
	assert.Equal(t, targetIndex, 1)
	assert.Equal(t, a.Len(), 10)
}

func TestLinearSearchNotFound(t *testing.T) {
	a := array.New(4).With([]int{1, 2, 3, 4})

	targetIndex := a.LinearSearch(6)
	assert.Equal(t, targetIndex, -1)
	assert.Equal(t, a.Len(), 4)
}

func TestLargest(t *testing.T) {

	a := array.New(6).With([]int{1, 2, 3, 4, 5, 6})

	largest := a.Largest()
	assert.Equal(t, largest, 6)
	assert.Equal(t, a.Len(), 6)
}

func TestLowest(t *testing.T) {

	a := array.New(4).With([]int{3, 4, 1, 2})

	lowest := a.Lowest()
	assert.Equal(t, lowest, 1)
	assert.Equal(t, a.Len(), 4)
}
