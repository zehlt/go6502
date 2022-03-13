package cpu

type Cpu struct {
	Cycle int
	Registers
}

func (c *Cpu) updateZeroAndNegativeFlags(value Register8) {
	if value == 0 {
		c.Status.Add(Zero)
	} else {
		c.Status.Remove(Zero)
	}

	if value.IsNegative() {
		c.Status.Add(Negative)
	} else {
		c.Status.Remove(Negative)
	}
}

func zeroPageOperandAddr(pc Register16, mem *Memory, register Register8) uint16 {
	operandAddr := mem.readByte(uint16(pc))
	operandAddr += uint8(register)
	return uint16(operandAddr)
}

func absoluteOperandAddr(pc Register16, mem *Memory, register Register8) uint16 {
	operandAddr := mem.readWord(uint16(pc))
	addr := operandAddr + uint16(register)
	return addr
}

func incrementWhenPageCrossed(c *Cpu, operandAddr uint16, resAddr uint16) {
	if (resAddr >> 8) != (operandAddr >> 8) {
		c.Cycle++
	}
}

func (c *Cpu) getOperandAddress(mem *Memory, mode int) uint16 {
	switch mode {
	case Implied:
		panic("Implied mode not supported")
	case Immediate:
		return uint16(c.ProgramCounter)
	case ZeroPage:
		return zeroPageOperandAddr(c.ProgramCounter, mem, 0)
	case ZeroPageX:
		return zeroPageOperandAddr(c.ProgramCounter, mem, c.XIndex)
	case ZeroPageY:
		return zeroPageOperandAddr(c.ProgramCounter, mem, c.YIndex)
	case Absolute:
		return absoluteOperandAddr(c.ProgramCounter, mem, 0)
	case AbsoluteX:
		return absoluteOperandAddr(c.ProgramCounter, mem, c.XIndex)
	case AbsoluteX1:
		operandAddr := mem.readWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.XIndex)
		incrementWhenPageCrossed(c, operandAddr, res)
		return res
	case AbsoluteY:
		return absoluteOperandAddr(c.ProgramCounter, mem, c.YIndex)
	case AbsoluteY1:
		operandAddr := mem.readWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.YIndex)
		incrementWhenPageCrossed(c, operandAddr, res)
		return res
	case IndirectX:
		operand := mem.readByte(uint16(c.ProgramCounter))
		operand += uint8(c.XIndex)
		return mem.readWord(uint16(operand))
	case IndirectY:
		operand := mem.readByte(uint16(c.ProgramCounter))
		word := mem.readWord(uint16(operand))
		res := word + uint16(c.YIndex)
		return res
	case IndirectY1:
		operand := mem.readByte(uint16(c.ProgramCounter))
		word := mem.readWord(uint16(operand))
		res := word + uint16(c.YIndex)
		incrementWhenPageCrossed(c, word, res)
		return res
	default:
		panic("Addressing mode not implemented")
	}
}

func lda(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.Accumulator = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func ldx(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.XIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func ldy(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.YIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func sta(c *Cpu, mem *Memory, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.writeByte(uint16(addrToFill), uint8(c.Accumulator))
}

func stx(c *Cpu, mem *Memory, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.writeByte(uint16(addrToFill), uint8(c.XIndex))
}

// TODO: impl this func
func brk(c *Cpu, mem *Memory, mode int) {

}

func sty(c *Cpu, mem *Memory, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.writeByte(uint16(addrToFill), uint8(c.YIndex))
}

func tax(c *Cpu, mem *Memory, mode int) {
	c.XIndex = c.Accumulator
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func tay(c *Cpu, mem *Memory, mode int) {
	c.YIndex = c.Accumulator
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func txa(c *Cpu, mem *Memory, mode int) {
	c.Accumulator = c.XIndex
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func tya(c *Cpu, mem *Memory, mode int) {
	c.Accumulator = c.YIndex
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: Maybe add tests for other mods
func and(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.Accumulator &= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func eor(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.Accumulator ^= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func aor(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.Accumulator |= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: may need some verif
func bit(c *Cpu, mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))

	aCopy := c.Accumulator
	aCopy &= Register8(operand)

	if aCopy == 0 {
		c.Status.Add(Zero)
	} else {
		c.Status.Remove(Zero)
	}

	if operand&Negative == 0 {
		c.Status.Remove(Negative)
	} else {
		c.Status.Add(Negative)
	}

	if operand&Verflow == 0 {
		c.Status.Remove(Verflow)
	} else {
		c.Status.Add(Verflow)
	}
}

func tsx(c *Cpu, mem *Memory, mode int) {
	c.XIndex = c.StackPointer
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func txs(c *Cpu, mem *Memory, mode int) {
	c.StackPointer = c.XIndex
}

// TODO: may wrap the stack pointer
func pha(c *Cpu, mem *Memory, mode int) {
	mem.writeByte(uint16(c.StackPointer), uint8(c.Accumulator))

	c.StackPointer--
}

func php(c *Cpu, mem *Memory, mode int) {
	mem.writeByte(uint16(c.StackPointer), uint8(c.Status))

	c.StackPointer--
}

// TODO: may wrap the stack pointer also
// Should learn about the B flag
func pla(c *Cpu, mem *Memory, mode int) {
	c.StackPointer++
	val := mem.readByte(uint16(c.StackPointer))
	c.Accumulator = Register8(val)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: should go deeper for tests
func plp(c *Cpu, mem *Memory, mode int) {
	c.StackPointer++
	val := mem.readByte(uint16(c.StackPointer))
	c.Status = Register8(val)
}

func inc(c *Cpu, mem *Memory, mode int) {
	addr := c.getOperandAddress(mem, mode)
	b := mem.readByte(addr)
	mem.writeByte(addr, b+1)
	c.updateZeroAndNegativeFlags(Register8(b + 1))
}

func inx(c *Cpu, mem *Memory, mode int) {
	c.XIndex++
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func iny(c *Cpu, mem *Memory, mode int) {
	c.YIndex++
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func dec(c *Cpu, mem *Memory, mode int) {
	addr := c.getOperandAddress(mem, mode)
	b := mem.readByte(addr)
	mem.writeByte(addr, b-1)
	c.updateZeroAndNegativeFlags(Register8(b - 1))
}

func dex(c *Cpu, mem *Memory, mode int) {
	c.XIndex--
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func dey(c *Cpu, mem *Memory, mode int) {
	c.YIndex--
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func rol(c *Cpu, mem *Memory, mode int) {
	if mode == Accumulator {
		isOldBit7Set := c.Accumulator&0b1000_0000 != 0
		c.Accumulator <<= 1
		c.Accumulator.Set(0, c.Status.Has(Carry))
		c.Status.Set(0, isOldBit7Set)
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(mem, mode)
		value := mem.readByte(addr)

		isOldBit7Set := value&0b1000_0000 != 0
		value <<= 1
		if c.Status.Has(Carry) {
			value |= 0b0000_0001
		} else {
			value &^= 0b0000_0001
		}
		c.Status.Set(0, isOldBit7Set)

		mem.writeByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

// TODO: Maybe adding more tests
func ror(c *Cpu, mem *Memory, mode int) {
	if mode == Accumulator {
		isOldBit0Set := c.Accumulator&0b0000_0001 != 0
		c.Accumulator >>= 1
		c.Accumulator.Set(7, c.Status.Has(Carry))
		c.Status.Set(0, isOldBit0Set)
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(mem, mode)
		value := mem.readByte(addr)

		isOldBit0Set := value&0b0000_0001 != 0
		value >>= 1

		if c.Status.Has(Carry) {
			value |= 0b1000_0000
		} else {
			value &^= 0b1000_0000
		}
		c.Status.Set(0, isOldBit0Set)

		mem.writeByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func lsr(c *Cpu, mem *Memory, mode int) {
	if mode == Accumulator {
		if c.Accumulator.Has(0b0000_0001) {
			c.Status.Add(Carry)
		} else {
			c.Status.Remove(Carry)
		}

		c.Accumulator >>= 1
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(mem, mode)
		value := mem.readByte(addr)

		if value&0b0000_0001 == 0 {
			c.Status.Remove(Carry)
		} else {
			c.Status.Add(Carry)
		}

		value >>= 1

		mem.writeByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func asl(c *Cpu, mem *Memory, mode int) {
	if mode == Accumulator {
		if c.Accumulator.Has(0b1000_0000) {
			c.Status.Add(Carry)
		} else {
			c.Status.Remove(Carry)
		}

		c.Accumulator <<= 1
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(mem, mode)
		value := mem.readByte(addr)

		if value&Negative == 0 {
			c.Status.Remove(Carry)
		} else {
			c.Status.Add(Carry)
		}

		value <<= 1

		mem.writeByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

// TODO: need to implement it
// TODO: check if the carry should be add before
// TODO: adding decilam addition later
func adc(c *Cpu, mem *Memory, mode int) {
	//c.Status.Remove(Carry)
}

func clc(c *Cpu, mem *Memory, mode int) {
	c.Status.Remove(Carry)
}

func cld(c *Cpu, mem *Memory, mode int) {
	c.Status.Remove(Decimal)
}

func cli(c *Cpu, mem *Memory, mode int) {
	c.Status.Remove(Interrupt)
}

func clv(c *Cpu, mem *Memory, mode int) {
	c.Status.Remove(Verflow)
}

func sec(c *Cpu, mem *Memory, mode int) {
	c.Status.Add(Carry)
}

func sed(c *Cpu, mem *Memory, mode int) {
	c.Status.Add(Decimal)
}

func sei(c *Cpu, mem *Memory, mode int) {
	c.Status.Add(Interrupt)
}

func (c *Cpu) interpret(opcode uint8, memory *Memory) {
	opc := Opcodes[opcode]

	opc.Operation(c, memory, opc.Mode)

	c.Cycle += opc.Cycles
	c.ProgramCounter += Register16(opc.ByteSize - 1)
}

func (c *Cpu) Reset(mem *Memory) {
	c.Accumulator = 0
	c.XIndex = 0
	c.YIndex = 0
	c.StackPointer = 0
	c.Status = 0

	c.ProgramCounter = Register16(mem.readWord(0xFFFC))
}

func (c *Cpu) Run(memory *Memory) {
	for {
		opcode := memory.readByte(uint16(c.ProgramCounter))
		c.ProgramCounter++

		c.interpret(opcode, memory)

		if opcode == BRK_IMP {
			break
		}
	}
}
