package array_test

import (
	"testing"

	"github.com/imyashkale/dsa/array"
	"github.com/magiconair/properties/assert"
)

func TestArray(t *testing.T) {
	a := array.NewArray()
	assert.Equal(t, a.Len, 0)
}
