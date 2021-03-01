package humantoken

import (
	"fmt"
	"math/rand"
	"time"
)

type TokenGenerator interface {
	Generate(size int) string
}

func Generate(size int, r *rand.Rand) string {
	const validChars = "23456789abcdefghjkmnpqrstvwxyz"

	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	s := ""

	for i := 0; i < size; i++ {
		idx := r.Intn(len(validChars))
		s = fmt.Sprintf("%s%c", s, validChars[idx])
	}

	return s
}
