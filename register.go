package go6502

type Register16 uint16
type Register8 uint8

type Registers struct {
	ProgramCounter Register16
	StackPointer   Register8
	Accumulator    Register8
	XIndex         Register8
	YIndex         Register8
	Status         Register8
}

func (r *Register8) IsNegative() bool {
	return *r >= 0b1000_0000
}

func (r *Register8) Add(bits uint8) {
	*r |= Register8(bits)
}

func (r *Register8) Remove(bits uint8) {
	*r &^= Register8(bits)
}

func (r *Register8) Set(pos uint8, boolean bool) {
	place := (1 << pos)
	if boolean {
		r.Add(uint8(place))
	} else {
		r.Remove(uint8(place))
	}
}

func (r *Register8) Has(bits uint8) bool {
	toTest := Register8(bits)
	toTest &= *r
	return toTest == Register8(bits)
}

func (r *Register8) Clear() {
	*r = 0
}
