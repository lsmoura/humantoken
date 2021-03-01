package humantoken

import (
	"math/rand"
	"testing"
)

func Test_ZeroSizeTokenWithRandomGenerator(t *testing.T) {
	g := NewGeneratorRandom()
	got := g.Generate(0)
	want := ""
	if got != want {
		t.Errorf("for Generate(0) want empty string, but got '%s'", got)
		t.Fail()
	}
}

func Test_SpecificSizeTokenWithRandomGenerator(t *testing.T) {
	g := NewGeneratorRandom()
	got := g.Generate(10)
	if len(got) != 10 {
		t.Errorf("for Generate(10) want string with size 10, but got '%s'", got)
		t.Fail()
	}
}

func Test_NewGeneratorWithRand(t *testing.T) {
	g := NewGenerator(rand.New(rand.NewSource(0)))
	got := g.Generate(10)
	if len(got) != 10 {
		t.Errorf("for Generate(10) want string with size 10, but got '%s'", got)
		t.Fail()
	}
}

func Test_NewGeneratorWithoutRand(t *testing.T) {
	g := NewGenerator(nil)
	got := g.Generate(12)
	if len(got) != 12 {
		t.Errorf("for Generate(12) want string with size 12, but got '%s'", got)
		t.Fail()
	}
}

func Test_NewGeneratorWithRandSameAsGenerateWithRand(t *testing.T) {
	g := NewGenerator(rand.New(rand.NewSource(0)))
	want := Generate(10, rand.New(rand.NewSource(0)))
	got := g.Generate(10)
	if got != want {
		t.Errorf("for Generate with same rand to have same result, but got '%s' and '%s'", want, got)
		t.Fail()
	}
}
