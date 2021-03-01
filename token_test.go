package humantoken

import (
	"math/rand"
	"testing"
)

func Test_ZeroSizeToken(t *testing.T) {
	got := Generate(0, nil)
	want := ""
	if got != want {
		t.Errorf("for Generate(0) want empty string, but got '%s'", got)
		t.Fail()
	}
}

func Test_SpecificSizeToken(t *testing.T) {
	got := Generate(10, nil)
	if len(got) != 10 {
		t.Errorf("for Generate(10) want string with size 10, but got '%s'", got)
		t.Fail()
	}
}

func Test_RespectUseOfRand(t *testing.T) {
	want := Generate(50, rand.New(rand.NewSource(0)))
	got := Generate(50, rand.New(rand.NewSource(0)))
	if got != want {
		t.Errorf("for Generate with same rand to have same result, but got '%s' and '%s'", want, got)
		t.Fail()
	}
}

func Test_ReusingRandDiffResults(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	want := Generate(50, r)
	got := Generate(50, r)
	if got == want {
		t.Errorf("for Generate with sequantial rand to have different results, but got the same '%s'", want)
		t.Fail()
	}
}
