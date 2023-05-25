package humantoken

import (
	"math/rand"
	"time"
)

type TokenGenerator interface {
	Generate(size int) string
}

func GenerateWithIndices(indices []int) string {
	const validChars = "23456789abcdefghjkmnpqrstvwxyz"
	s := make([]byte, len(indices))

	for i, v := range indices {
		s[i] = validChars[v%len(validChars)]
	}

	return string(s)
}

func Generate(size int, r *rand.Rand) string {
	if size <= 0 {
		return ""
	}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	intList := make([]int, size)
	for i := range intList {
		intList[i] = r.Intn(256)
	}

	return GenerateWithIndices(intList)
}
