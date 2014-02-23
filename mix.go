package mix

import (
	"fmt"
)

// A Byte represents one of 100 distinct values, on this particular MIX computer.
type Byte int8

// The Sign of a field is either positive or negative.
type Sign int8

// A Word is [Sign, Byte, Byte, Byte, Byte, Byte]
type Word [6]Byte

// An Address is [Sign, Byte, Byte]
type Address [3]Byte

// A JumpAddress is [Byte, Byte]
type JumpAddress [2]Byte

func (b Byte) String() string {
	if b > 99 || b < 0 {
		return "??"
	}
	return fmt.Sprintf("%02d", byte(b))
}

func (s Sign) String() string {
	if s < 0 {
		return "-"
	}
	return "+"
}

func (w Word) String() string {
	return fmt.Sprint(Sign(w[0]), w[1], w[2], w[3], w[4], w[5])
}

func (a Address) String() string {
	return fmt.Sprint(Sign(a[0]), a[1], a[2])
}

func (j JumpAddress) String() string {
	return fmt.Sprint(j[0], j[1])
}
