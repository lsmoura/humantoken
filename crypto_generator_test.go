package humantoken

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroSizeTokenWithCryptoGenerator(t *testing.T) {
	g := NewCryptoGenerator()
	got := g.Generate(0)

	assert.Equal(t, "", got)
}

func TestFixedSizeTokenWithCryptoGenerator(t *testing.T) {
	g := NewCryptoGenerator()
	got := g.Generate(10)

	assert.Len(t, got, 10)
}
