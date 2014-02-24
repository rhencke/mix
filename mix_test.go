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
		13: {LESS, "LESS"},
		14: {EQUAL, "EQUAL"},
		15: {GREATER, "GREATER"},
		16: {Comparison(255), "??"},
	}

	for i, v := range tests {
		actual := v.value.String()
		if actual != v.expected {
			t.Errorf("test #%v, expected %v, got %v", i, v.expected, actual)
		}
	}
}

func TestCFIAA(t *testing.T) {
	w := Word{-1, 2, 3, 4, 5, 6}

	var expectedA = Address{-1, 2, 3}
	if a := w.AA(); a != expectedA {
		t.Errorf("expected address - 02 03, got %v", a)
	}

	if i := w.I(); i != 4 {
		t.Errorf("expected index register 04, got %v", i)
	}
	if f := w.F(); f != 5 {
		t.Errorf("expected field value 05, got %v", f)
	}

	if c := w.C(); c != 6 {
		t.Errorf("expected operation code 06, got %v", c)
	}
}
