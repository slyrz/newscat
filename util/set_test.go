package util

import (
	"testing"
)

func TestBitset(t *testing.T) {
	s := NewBitset()

	// This should fill all buckets.
	for i := 1; i <= s.Cap(); i++ {
		if s.Add(uint32(i)); s.Len() != i {
			t.Errorf("unexpected length")
		}
	}

	if s.Len() != s.Cap() {
		t.Errorf("test did not utilize all buckets")
	}
}

func TestBitsetUnion(t *testing.T) {
	u := NewBitset()
	u.Add(2)
	u.Add(3)
	u.Add(5)
	u.Add(8)

	v := NewBitset()
	v.Add(1)
	v.Add(2)
	v.Add(6)
	v.Add(8)

	if u.Len() != 4 || v.Len() != 4 {
		t.Errorf("unexpected length")
	}

	if u.Union(v); u.Len() != 6 {
		t.Errorf("union failed")
	}
}

func TestStringset(t *testing.T) {
	s := NewStringset()
	s.Add("hello")
	s.Add("world")
	if s.Len() != 2 {
		t.Errorf("unexpected length")
	}
}
