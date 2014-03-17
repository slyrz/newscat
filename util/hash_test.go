package util

import (
	"testing"
)

func TestHash(t *testing.T) {
	if Hash("abc") != Hash("abc") {
		t.Errorf("Hash(x) != Hash(x)")
	}

	if Hash("abc") == Hash("xyz") {
		t.Errorf("Hash(x) == Hash(y): what are the odds?")
	}
}
