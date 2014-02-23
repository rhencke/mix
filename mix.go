package mix

import (
	"fmt"
)

// A Byte represents one of 100 distinct values, on this particular MIX computer.
type Byte byte

func (b *Byte) String() string {
	if *b > 99 {
		return "??"
	}
	return fmt.Sprintf("%02d", byte(*b))
}

type Computer struct {
}
