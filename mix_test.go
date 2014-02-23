package mix

import (
	"testing"
)

func TestByteString(t *testing.T) {
	var tests = []struct {
		value    Byte
		expected string
	}{
		{-1, "??"},
		{0, "00"},
		{99, "99"},
		{100, "??"},
	}

	for i, v := range tests {
		actual := v.value.String()
		if actual != v.expected {
			t.Errorf("test #%v, expected %v, got %v", i, v.expected, actual)
		}
	}
}

func TestNilByteString(t *testing.T) {
	actual := (*Byte)(nil).String()
	if actual != "??" {
		t.Error("expected ??, got", actual)
	}
}
