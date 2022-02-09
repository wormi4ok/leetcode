package leetcode

import (
	"math"
	"math/bits"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test991_BrokenCalculator(t *testing.T) {
	assert.Equal(t, 1, brokenCalc(1, 2))
	assert.Equal(t, 2, brokenCalc(2, 3))
	assert.Equal(t, 2, brokenCalc(5, 8))
	assert.Equal(t, 3, brokenCalc(3, 10))
	assert.Equal(t, 7, brokenCalc(10, 3))
	assert.Equal(t, 39, brokenCalc(1, 1000000000))
}

// Following solution from https://ansonvandoren.com/posts/broken-calculator/
func brokenCalc(startValue int, target int) int {
	if target <= startValue {
		return startValue - target
	}

	n := bits.Len64(uint64((target - 1) / startValue))
	bucket := int(math.Pow(2, float64(n)))

	i := startValue*bucket - target

	binary := strconv.FormatInt(int64(i), 2)
	if len(binary) < n {
		binary = strings.Repeat("0", n-len(binary)) + binary
	}
	lsbOnes := strings.Count(binary[len(binary)-n:], "1")
	msbVal := i >> n

	return n + lsbOnes + msbVal
}
