package mix

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	var tests = []struct {
		value    fmt.Stringer
		expected string
	}{
		0:  {Byte(-1), "??"},
		1:  {Byte(0), "00"},
		2:  {Byte(99), "99"},
		3:  {Byte(100), "??"},
		4:  {Sign(-50), "-"},
		5:  {Sign(0), "+"},
		6:  {Sign(50), "+"},
		7:  {Word{0, 0, 0, 0, 0, 0}, "+ 00 00 00 00 00"},
		8:  {Word{-1, 23, 45, 67, 89, 0}, "- 23 45 67 89 00"},
		9:  {Address{0, 0, 0}, "+ 00 00"},
		10: {Address{-1, 23, 45}, "- 23 45"},
		11: {JumpAddress{0, 0}, "00 00"},
		12: {JumpAddress{23, 45}, "23 45"},
	}

	for i, v := range tests {
		actual := v.value.String()
		if actual != v.expected {
			t.Errorf("test #%v, expected %v, got %v", i, v.expected, actual)
		}
	}
}
