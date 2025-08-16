package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phab/math"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, math.Add(1, 2))
}

func TestSub(t *testing.T) {
	assert.Equal(t, 1, math.Sub(3, 2))
}
