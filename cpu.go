package go6502

// import "fmt"

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

func zeroPageOperandAddr(pc Register16, mem *Mem, register Register8) uint16 {
	operandAddr := mem.ReadByte(uint16(pc))
	operandAddr += uint8(register)
	return uint16(operandAddr)
}

func absoluteOperandAddr(pc Register16, mem *Mem, register Register8) uint16 {
	operandAddr := mem.ReadWord(uint16(pc))
	addr := operandAddr + uint16(register)
	return addr
}

func incrementWhenPageCrossed(c *Cpu, operandAddr uint16, resAddr uint16) {
	if (resAddr >> 8) != (operandAddr >> 8) {
		c.Cycle++
	}
}

func (c *Cpu) getOperandAddress(mem *Mem, mode int) uint16 {
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
		operandAddr := mem.ReadWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.XIndex)
		incrementWhenPageCrossed(c, operandAddr, res)
		return res
	case AbsoluteY:
		return absoluteOperandAddr(c.ProgramCounter, mem, c.YIndex)
	case AbsoluteY1:
		operandAddr := mem.ReadWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.YIndex)
		incrementWhenPageCrossed(c, operandAddr, res)
		return res
	case Indirect:
		operand := mem.ReadWord(uint16(c.ProgramCounter))
		indirectValue := mem.ReadWord(operand)
		return indirectValue
	case IndirectX:
		operand := mem.ReadByte(uint16(c.ProgramCounter))
		operand += uint8(c.XIndex)
		return mem.ReadWord(uint16(operand))
	case IndirectY:
		operand := mem.ReadByte(uint16(c.ProgramCounter))
		word := mem.ReadWord(uint16(operand))
		res := word + uint16(c.YIndex)
		return res
	case IndirectY1:
		operand := mem.ReadByte(uint16(c.ProgramCounter))
		word := mem.ReadWord(uint16(operand))
		res := word + uint16(c.YIndex)
		incrementWhenPageCrossed(c, word, res)
		return res
	default:
		panic("Addressing mode not implemented")
	}
}

func lda(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.Accumulator = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func ldx(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.XIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func ldy(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.YIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func sta(c *Cpu, mem *Mem, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.WriteByte(uint16(addrToFill), uint8(c.Accumulator))
}

func stx(c *Cpu, mem *Mem, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.WriteByte(uint16(addrToFill), uint8(c.XIndex))
}

// TODO: impl this func
// should call the interrupt
func brk(c *Cpu, mem *Mem, mode int) {
}

func nop(c *Cpu, mem *Mem, mode int) {
	// purposely does nothing
}

// TODO: add some tests maybe ?
func rti(c *Cpu, mem *Mem, mode int) {
	statusFlags := popStack(c, mem)
	pc := popStack(c, mem)

	c.ProgramCounter = Register16(pc)
	c.Status = Register8(statusFlags)
}

func sty(c *Cpu, mem *Mem, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.WriteByte(uint16(addrToFill), uint8(c.YIndex))
}

func tax(c *Cpu, mem *Mem, mode int) {
	c.XIndex = c.Accumulator
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func tay(c *Cpu, mem *Mem, mode int) {
	c.YIndex = c.Accumulator
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func txa(c *Cpu, mem *Mem, mode int) {
	c.Accumulator = c.XIndex
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func tya(c *Cpu, mem *Mem, mode int) {
	c.Accumulator = c.YIndex
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: Maybe add tests for other mods
func and(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.Accumulator &= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func eor(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.Accumulator ^= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func aor(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.Accumulator |= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: may need some verif
func bit(c *Cpu, mem *Mem, mode int) {
	operand := mem.ReadByte(c.getOperandAddress(mem, mode))

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

func tsx(c *Cpu, mem *Mem, mode int) {
	c.XIndex = c.StackPointer
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func txs(c *Cpu, mem *Mem, mode int) {
	c.StackPointer = c.XIndex
}

func pushStack(c *Cpu, mem *Mem, value uint8) {
	mem.WriteByte(0x0100+uint16(c.StackPointer), value)
	c.StackPointer--
}

func popStack(c *Cpu, mem *Mem) uint8 {
	c.StackPointer++
	val := mem.ReadByte(0x0100 + uint16(c.StackPointer))

	return val
}

// TODO: may wrap the stack pointer
func pha(c *Cpu, mem *Mem, mode int) {
	pushStack(c, mem, uint8(c.Accumulator))
}

func php(c *Cpu, mem *Mem, mode int) {
	flags := c.Status
	flags.Add(Break)
	flags.Add(Break2)

	pushStack(c, mem, uint8(c.Status))
}

// TODO: may wrap the stack pointer also
// Should learn about the B flag
func pla(c *Cpu, mem *Mem, mode int) {
	val := popStack(c, mem)
	c.Accumulator = Register8(val)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: should go deeper for tests
func plp(c *Cpu, mem *Mem, mode int) {
	val := popStack(c, mem)
	c.Status = Register8(val)

	c.Status.Remove(Break)
	c.Status.Add(Break2)
}

func inc(c *Cpu, mem *Mem, mode int) {
	addr := c.getOperandAddress(mem, mode)
	b := mem.ReadByte(addr)
	mem.WriteByte(addr, b+1)
	c.updateZeroAndNegativeFlags(Register8(b + 1))
}

func inx(c *Cpu, mem *Mem, mode int) {
	c.XIndex++
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func iny(c *Cpu, mem *Mem, mode int) {
	c.YIndex++
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func dec(c *Cpu, mem *Mem, mode int) {
	addr := c.getOperandAddress(mem, mode)
	b := mem.ReadByte(addr)
	mem.WriteByte(addr, b-1)
	c.updateZeroAndNegativeFlags(Register8(b - 1))
}

func dex(c *Cpu, mem *Mem, mode int) {
	c.XIndex--
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func dey(c *Cpu, mem *Mem, mode int) {
	c.YIndex--
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func rol(c *Cpu, mem *Mem, mode int) {
	if mode == Accumulator {
		isOldBit7Set := c.Accumulator&0b1000_0000 != 0
		c.Accumulator <<= 1
		c.Accumulator.Set(0, c.Status.Has(Carry))
		c.Status.Set(0, isOldBit7Set)
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(mem, mode)
		value := mem.ReadByte(addr)

		isOldBit7Set := value&0b1000_0000 != 0
		value <<= 1
		if c.Status.Has(Carry) {
			value |= 0b0000_0001
		} else {
			value &^= 0b0000_0001
		}
		c.Status.Set(0, isOldBit7Set)

		mem.WriteByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

// TODO: Maybe adding more tests
func ror(c *Cpu, mem *Mem, mode int) {
	if mode == Accumulator {
		isOldBit0Set := c.Accumulator&0b0000_0001 != 0
		c.Accumulator >>= 1
		c.Accumulator.Set(7, c.Status.Has(Carry))
		c.Status.Set(0, isOldBit0Set)
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(mem, mode)
		value := mem.ReadByte(addr)

		isOldBit0Set := value&0b0000_0001 != 0
		value >>= 1

		if c.Status.Has(Carry) {
			value |= 0b1000_0000
		} else {
			value &^= 0b1000_0000
		}
		c.Status.Set(0, isOldBit0Set)

		mem.WriteByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func lsr(c *Cpu, mem *Mem, mode int) {
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
		value := mem.ReadByte(addr)

		if value&0b0000_0001 == 0 {
			c.Status.Remove(Carry)
		} else {
			c.Status.Add(Carry)
		}

		value >>= 1

		mem.WriteByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func asl(c *Cpu, mem *Mem, mode int) {
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
		value := mem.ReadByte(addr)

		if value&Negative == 0 {
			c.Status.Remove(Carry)
		} else {
			c.Status.Add(Carry)
		}

		value <<= 1

		mem.WriteByte(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func (c *Cpu) addToRegisterA(value uint8) {
	var sum uint16 = uint16(c.Accumulator) + uint16(value)
	if c.Status.Has(Carry) {
		sum += 1
	}

	isGreaterThan8bit := sum > 0xff

	if isGreaterThan8bit {
		c.Status.Add(Carry)
	} else {
		c.Status.Remove(Carry)
	}

	var result uint8 = uint8(sum)
	if (value^result)&(result^uint8(c.Accumulator))&0x80 != 0 {
		c.Status.Add(Verflow)
	} else {
		c.Status.Remove(Verflow)
	}
	c.Accumulator = Register8(result)
}

// TODO: adding decilam addition later
func adc(c *Cpu, mem *Mem, mode int) {
	addr := c.getOperandAddress(mem, mode)
	data := mem.ReadByte(addr)
	c.addToRegisterA(data)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: add tests
func sbc(c *Cpu, mem *Mem, mode int) {
	addr := c.getOperandAddress(mem, mode)
	data := mem.ReadByte(addr)
	data = -data
	c.addToRegisterA(data)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: add tests
func (c *Cpu) compare(mem *Mem, compareWith uint8) {

	if compareWith <= uint8(c.Accumulator) {
		c.Status.Add(Carry)
	} else {
		c.Status.Remove(Carry)
	}

}

func cmp(c *Cpu, mem *Mem, mode int) {
	data := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.compare(mem, data)
	c.updateZeroAndNegativeFlags(c.Accumulator - Register8(data))
}

func cpx(c *Cpu, mem *Mem, mode int) {
	data := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.compare(mem, data)
	c.updateZeroAndNegativeFlags(c.XIndex - Register8(data))
}

func cpy(c *Cpu, mem *Mem, mode int) {
	data := mem.ReadByte(c.getOperandAddress(mem, mode))
	c.compare(mem, data)
	c.updateZeroAndNegativeFlags(c.YIndex - Register8(data))
}

func jsr(c *Cpu, mem *Mem, mode int) {
	operand := c.getOperandAddress(mem, mode)
	pushStack(c, mem, uint8((c.ProgramCounter+2)>>8))
	pushStack(c, mem, uint8((c.ProgramCounter + 2)))
	c.ProgramCounter = Register16(operand)
}

// TODO: check if the order of operation is right
func rts(c *Cpu, mem *Mem, mode int) {
	var lo uint16 = uint16(popStack(c, mem))
	var hi uint16 = uint16(popStack(c, mem))

	// fmt.Printf("before rts: %04x\n", c.ProgramCounter)
	c.ProgramCounter = Register16(((hi << 8) | lo))
	// fmt.Printf("after  rts: %04x\n", c.ProgramCounter)
}

func branch(c *Cpu, mem *Mem, condition bool) {
	if condition {
		var jump int8 = int8(mem.ReadByte(uint16(c.ProgramCounter)))
		c.ProgramCounter += Register16(jump)
	}
}

// TODO: may check if branching works correctly to the TODO
func bcc(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, !c.Status.Has(Carry))
}

func bcs(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, c.Status.Has(Carry))
}

func beq(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, c.Status.Has(Zero))
}

func bmi(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, c.Status.Has(Negative))
}

func bne(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, !c.Status.Has(Zero))
}

func bpl(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, !c.Status.Has(Negative))
}

func bvc(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, !c.Status.Has(Verflow))
}

func bvs(c *Cpu, mem *Mem, mode int) {
	branch(c, mem, c.Status.Has(Verflow))
}

// TODO: check if the increment doesn't happen in the same instruction cycle
func jmp(c *Cpu, mem *Mem, mode int) {
	operand := c.getOperandAddress(mem, mode)
	c.ProgramCounter = Register16(operand)
}

func clc(c *Cpu, mem *Mem, mode int) {
	c.Status.Remove(Carry)
}

func cld(c *Cpu, mem *Mem, mode int) {
	c.Status.Remove(Decimal)
}

func cli(c *Cpu, mem *Mem, mode int) {
	c.Status.Remove(Interrupt)
}

func clv(c *Cpu, mem *Mem, mode int) {
	c.Status.Remove(Verflow)
}

func sec(c *Cpu, mem *Mem, mode int) {
	c.Status.Add(Carry)
}

func sed(c *Cpu, mem *Mem, mode int) {
	c.Status.Add(Decimal)
}

func sei(c *Cpu, mem *Mem, mode int) {
	c.Status.Add(Interrupt)
}

func (c *Cpu) interpret(opcode uint8, mem *Mem) {
	opc := Opcodes[opcode]

	opc.Operation(c, mem, opc.Mode)

	c.Cycle += opc.Cycles
	c.ProgramCounter += Register16(opc.ByteSize - 1)
}

func (c *Cpu) Reset(mem *Mem) {
	c.Accumulator = 0
	c.XIndex = 0
	c.YIndex = 0
	c.StackPointer = 0xFF
	c.Status = 0

	c.ProgramCounter = Register16(mem.ReadWord(0xFFFC))
}

func (c *Cpu) Step(mem *Mem) bool {
	opcode := mem.ReadByte(uint16(c.ProgramCounter))
	// fmt.Printf("[%04x]: %02x, (SP: %02x, A: %02x)\n", c.ProgramCounter, opcode, c.StackPointer, c.Accumulator)

	c.ProgramCounter++

	c.interpret(opcode, mem)

	return opcode == BRK_IMP
}

func (c *Cpu) Run(mem *Mem) {
	for {
		shouldExit := c.Step(mem)

		if shouldExit {
			break
		}
	}
}
