package humantoken

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestZeroSizeToken(t *testing.T) {
	got := Generate(0, nil)

	assert.Len(t, got, 0)
}

func TestSpecificSizeToken(t *testing.T) {
	got := Generate(10, nil)

	assert.Len(t, got, 10)
}

func TestRespectUseOfRand(t *testing.T) {
	want := Generate(50, rand.New(rand.NewSource(0)))
	got := Generate(50, rand.New(rand.NewSource(0)))

	assert.Equal(t, got, want)
}

func TestReusingRandDiffResults(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	want := Generate(50, r)
	got := Generate(50, r)

	assert.NotEqual(t, got, want)
}

func TestGenerateWithSameIndices(t *testing.T) {
	idx0 := []int{1, 1, 1}
	idx1 := []int{1, 1, 1}

	want := GenerateWithIndices(idx0)
	got := GenerateWithIndices(idx1)

	assert.Equal(t, got, want)
}

func TestGenerateWithDifferentIndices(t *testing.T) {
	idx0 := []int{1, 1, 1}
	idx1 := []int{2, 2, 2}

	want := GenerateWithIndices(idx0)
	got := GenerateWithIndices(idx1)

	assert.NotEqual(t, got, want)
}
