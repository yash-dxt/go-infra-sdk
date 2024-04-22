package helpers_test

import (
	"testing"

	ut "github.com/yash-dxt/go-infra-sdk/helpers"

	"github.com/stretchr/testify/assert"
)

type testObj struct {
	a float64
	b float64
}

func TestArrayMap(t *testing.T) {
	objArr := []testObj{
		{
			a: 21,
			b: 0.5,
		},
		{
			a: 42,
			b: 1,
		},
	}

	mapped := ut.ArrayMap(objArr, func(ms testObj) float64 { return ms.a })

	assert.Equal(t, []float64{21, 42}, mapped, "ArrayMap")
	assert.Equal(t, 63.0, ut.Sum(mapped), "Sum")
}

func TestArrayChunk(t *testing.T) {
	objArr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assert.Equal(
		t,
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8},
		},
		ut.ArrayChunk(objArr, 3),
		"ArrayChunk",
	)

	assert.Equal(
		t,
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		ut.ArrayChunk(objArr, 4),
		"ArrayChunk",
	)

	assert.Equal(
		t,
		[][]int{
			{1, 2, 3, 4, 5, 6, 7, 8},
		},
		ut.ArrayChunk(objArr, 10),
		"ArrayChunk",
	)
}
