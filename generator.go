package humantoken

import (
	"math/rand"
	"time"
)

type TokenGen struct {
	r *rand.Rand
}

func (g TokenGen) Generate(size int) string {
	return Generate(size, g.r)
}

func NewGenerator(r *rand.Rand) TokenGenerator {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return TokenGen{r}
}

func NewGeneratorRandom() TokenGenerator {
	return TokenGen{r: rand.New(rand.NewSource(time.Now().UnixNano()))}
}
