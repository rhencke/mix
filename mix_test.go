package mix

import (
	"testing"
)

func TestByteString(t *testing.T) {
	var tests = []struct {
		value    Byte
		expected string
	}{
		{0, "00"},
		{99, "99"},
		{100, "??"},
	}

	for i, v := range tests {
		if v.value.String() != v.expected {
			t.Errorf("test #%v, expected %v, got %v", i, v.expected, v.value.String())
		}
	}
}
