package calc

import (
    "testing"
    "math"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
    type test struct {
        a    int
        b    int
        want int
    }

    tests := []test{
        {a: 1, b: 2, want: 3},
        {a: 1, b: 0, want: 1},
        {a: -2, b: -3, want: -5},
		{a: -2, b: 3, want: 1},
        {a: math.MaxInt, b: 1, want: math.MinInt},
		{a: math.MinInt, b: -1, want: math.MaxInt},
    }
    for _, tc := range tests {
		got := Sum(tc.a, tc.b)
		assert.Equal(t, tc.want, got, "expected: %v, got: %v", tc.want, got)
    }
}


func TestMultiply (t *testing.T) {
	// arrange
	c := NewCalc()

	// act
	result := c.Multiply(2, 2)

	// assert
	assert.Equal(t, 4, result)
}
