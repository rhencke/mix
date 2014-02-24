package mix

import (
	"fmt"
	"io"
)

// The 56 characters of the MIX character set, indexed by its character code.
const CharacterCodes = " ABCDEFGHIΔJKLMNOPQRΣΠSTUVWXYZ0123456789.,()+-*/=$<>@;:'"

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

// The Comparison indicator is one of: LESS, EQUAL, GREATER
type Comparison byte

const (
	LESS Comparison = iota
	EQUAL
	GREATER
)

// The Computer is a MIX computer.
type Computer struct {
	// Accumulator register
	A Word

	// Extension register
	X Word

	// Index registers
	I1, I2, I3, I4, I5, I6 Address

	// Jump register
	J JumpAddress

	// Overflow toggle
	Overflow bool

	// Comparison indicator
	Comparison Comparison

	// Memory cells
	Memory [4000]Word

	// Input/output devices
	Tape        [8]io.ReadWriteSeeker
	Disk        [8]io.ReadWriteSeeker
	CardReader  io.Reader
	CardPunch   io.Writer
	LinePrinter io.Writer
	Typewriter  io.Reader
	PaperTape   io.ReadWriteSeeker

	// The address of the current instruction, or ⊛
	LocationCounter int
}

func (c *Computer) Step() {
	mem := c.Memory[c.LocationCounter]
	switch mem.C() {

	}
}

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

// C is the operation code of this word.
func (w Word) C() Byte {
	return w[5]
}

// F is the modification of the operation code of this word.
func (w Word) F() Byte {
	return w[4]
}

// I is the index register to be used, or 0 if none.
func (w Word) I() Byte {
	return w[3]
}

// AA is the address of this word.
func (w Word) AA() Address {
	return Address{w[0], w[1], w[2]}
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

func (c Comparison) String() string {
	switch c {
	case LESS:
		return "LESS"
	case EQUAL:
		return "EQUAL"
	case GREATER:
		return "GREATER"
	default:
		return "??"
	}
}
