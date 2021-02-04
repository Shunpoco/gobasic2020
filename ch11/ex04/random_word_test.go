package word

import (
	"math/rand"
	"testing"
	"time"
)

var symbols = []rune{0x20, 0x2c, 0x2e} // space, comma, period

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n+1)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	// Insert a symbol in the slice at a given index.
	iS := rng.Intn(len(symbols)) // An index in symbols
	iRs := rng.Intn(n + 1)       // An index in runes
	copy(runes[iRs+1:], runes[iRs:len(runes)-1])
	runes[iRs] = symbols[iS]
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false, %d", p, i)
		}
	}
}
