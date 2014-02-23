package mix

import (
	"fmt"
)

// A Byte represents one of 100 distinct values, on this particular MIX computer.
type Byte int8

// A Word is composed of 6 fields, addressed from 0 to 5:
//   [Sign, Byte, Byte, Byte, Byte, Byte]

func (b *Byte) String() string {
	if b == nil || *b > 99 || *b < 0 {
		return "??"
	}
	return fmt.Sprintf("%02d", byte(*b))
}
