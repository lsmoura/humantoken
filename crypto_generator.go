package humantoken

import (
	"crypto/rand"
	"fmt"
)

type TokenCryptoGen struct{}

func (g TokenCryptoGen) Generate(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	intList := make([]int, size)
	for i, v := range b {
		intList[i] = int(v)
	}

	return GenerateWithIndices(intList)
}

func NewCryptoGenerator() TokenGenerator {
	return TokenCryptoGen{}
}
