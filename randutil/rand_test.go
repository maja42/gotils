package randutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustRandString(t *testing.T) {
	testCount := 1000
	numeric := []rune("0123456789")

	for i := 0; i < testCount; i++ {
		str := MustRandString(200, numeric)
		illegal := strings.Trim(str, string(numeric))
		assert.Empty(t, illegal)
	}
}

func TestMustRandString_noDuplicates(t *testing.T) {
	testCount := 1000

	numeric := []rune("0123456789")
	found := make(map[string]struct{}, testCount)
	for i := 0; i < testCount; i++ {
		str := MustRandString(200, numeric)
		found[str] = struct{}{}
	}
	if len(found) != testCount {
		assert.Fail(t, "random generator produced duplicate results")
	}
}

func TestMustRandBytes_noDuplicates(t *testing.T) {
	testCount := 1000

	found := make(map[[10]byte]struct{}, testCount)
	for i := 0; i < testCount; i++ {
		var arr [10]byte
		copy(arr[:], MustRandBytes(20))
		found[arr] = struct{}{}
	}
	if len(found) != testCount {
		assert.Fail(t, "random generator produced duplicate results")
	}
}
