package humantoken

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		name           string
		generator      TokenGenerator // if nil, NewGeneratorRandom() is used
		size           int
		expectedLength int
	}{
		{
			name:           "zero size",
			size:           0,
			expectedLength: 0,
		},
		{
			name:           "ten size",
			size:           10,
			expectedLength: 10,
		},
		{
			name:           "negative size",
			size:           -10,
			expectedLength: 0,
		},
		{
			name:           "zero seed generator",
			generator:      NewGenerator(rand.New(rand.NewSource(0))),
			size:           10,
			expectedLength: 10,
		},
		{
			name:           "nil generator",
			generator:      NewGenerator(nil),
			size:           12,
			expectedLength: 12,
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			g := test.generator
			if g == nil {
				g = NewGeneratorRandom()
			}
			got := g.Generate(test.size)

			assert.Len(t, got, test.expectedLength)
		})
	}
}

func TestNewGeneratorWithRandSameAsGenerateWithRand(t *testing.T) {
	sources := make([]int64, 10)

	for i := range sources {
		sources[i] = rand.Int63()
	}

	for _, sourceSeed := range sources {
		g := NewGenerator(rand.New(rand.NewSource(0)))
		want := Generate(10, rand.New(rand.NewSource(0)))
		got := g.Generate(10)

		assert.Equalf(t, want, got, "for Generate with %d seed to have same result", sourceSeed)
	}
}
