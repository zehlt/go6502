package cpu

type Register16 uint16
type Register8 uint8
type RegisterBitset8 uint8

type Registers struct {
	ProgramCounter Register16
	StackPointer   Register8
	Accumulator    Register8
	XIndex         Register8
	YIndex         Register8
	Status         RegisterBitset8
}

func (r *Register8) IsNegative() bool {
	return *r >= 0b1000_0000
}

func (r *RegisterBitset8) Add(bits uint8) {
	*r |= RegisterBitset8(bits)
}

func (r *RegisterBitset8) Remove(bits uint8) {
	*r &^= RegisterBitset8(bits)
}

func (r *RegisterBitset8) Has(bits uint8) bool {
	toTest := RegisterBitset8(bits)
	toTest &= *r
	return toTest == RegisterBitset8(bits)
}

func (r *RegisterBitset8) Clear() {
	*r = 0
}
