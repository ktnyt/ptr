package ptr

import "testing"

func TestNew(t *testing.T) {
	in := 42
	out := New(in)
	if in != *out {
		t.Errorf("*New(%v) = %v, want %v", in, *out, in)
	}
}
