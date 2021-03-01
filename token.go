package humantoken

import (
	"fmt"
	"math/rand"
	"time"
)

type TokenGenerator interface {
	Generate(size int) string
}

func GenerateWithIndices(indices []int) string {
	const validChars = "23456789abcdefghjkmnpqrstvwxyz"
	s := ""

	for i := 0; i < len(indices); i++ {
		idx := indices[i] % len(validChars)
		s = fmt.Sprintf("%s%c", s, validChars[idx])
	}

	return s
}

func Generate(size int, r *rand.Rand) string {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	intList := make([]int, size)
	for i := range intList {
		intList[i] = r.Intn(256)
	}

	return GenerateWithIndices(intList)
}
