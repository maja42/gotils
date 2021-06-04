package randutil

import (
	"crypto/rand"
	"math/big"
)

var reader = rand.Reader

// MustRandBytes returns securely generated random bytes.
// Panics if the system's secure random number generator fails.
func MustRandBytes(n int) []byte {
	r, err := RandBytes(n)
	if err != nil {
		panic(err)
	}
	return r
}

// RandBytes returns securely generated random bytes.
// Returns an error if the system's secure random number generator fails.
func RandBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := reader.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// MustRandString returns a securely generated random string.
// Panics if the system's secure random number generator fails.
func MustRandString(n int, runes []rune) string {
	s, err := RandString(n, runes)
	if err != nil {
		panic(err)
	}
	return s
}

// RandString returns a securely generated random string.
// Returns an error if the system's secure random number generator fails.
func RandString(n int, runes []rune) (string, error) {
	runeCount := big.NewInt(int64(len(runes)))

	ret := make([]rune, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(reader, runeCount)
		if err != nil {
			return "", err
		}
		ret[i] = runes[num.Int64()]
	}
	return string(ret), nil
}
