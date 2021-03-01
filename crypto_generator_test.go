package humantoken

import "testing"

func Test_ZeroSizeTokenWithCryptoGenerator(t *testing.T) {
	g := NewCryptoGenerator()
	got := g.Generate(0)
	want := ""
	if got != want {
		t.Errorf("for Generate(0) want empty string, but got '%s'", got)
		t.Fail()
	}
}

func Test_FixedSizeTokenWithCryptoGenerator(t *testing.T) {
	g := NewCryptoGenerator()
	got := g.Generate(10)
	if len(got) != 10 {
		t.Errorf("for Generate(10) want string with 10 characters, but got '%s'", got)
		t.Fail()
	}
}
