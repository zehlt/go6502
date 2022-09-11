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

func zeroPageOperandAddr(pc Register16, bus Bus, register Register8) uint16 {
	operandAddr := bus.Read(uint16(pc))
	operandAddr += uint8(register)
	return uint16(operandAddr)
}

func absoluteOperandAddr(pc Register16, bus Bus, register Register8) uint16 {
	operandAddr := bus.ReadWord(uint16(pc))
	addr := operandAddr + uint16(register)
	return addr
}

func incrementWhenPageCrossed(c *Cpu, operandAddr uint16, resAddr uint16) {
	if (resAddr >> 8) != (operandAddr >> 8) {
		c.Cycle++
	}
}

func (c *Cpu) getOperandAddress(bus Bus, mode int) uint16 {
	switch mode {
	case Implied:
		panic("Implied mode not supported")
	case Immediate:
		return uint16(c.ProgramCounter)
	case ZeroPage:
		return zeroPageOperandAddr(c.ProgramCounter, bus, 0)
	case ZeroPageX:
		return zeroPageOperandAddr(c.ProgramCounter, bus, c.XIndex)
	case ZeroPageY:
		return zeroPageOperandAddr(c.ProgramCounter, bus, c.YIndex)
	case Absolute:
		return absoluteOperandAddr(c.ProgramCounter, bus, 0)
	case AbsoluteX:
		return absoluteOperandAddr(c.ProgramCounter, bus, c.XIndex)
	case AbsoluteX1:
		operandAddr := bus.ReadWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.XIndex)
		incrementWhenPageCrossed(c, operandAddr, res)
		return res
	case AbsoluteY:
		return absoluteOperandAddr(c.ProgramCounter, bus, c.YIndex)
	case AbsoluteY1:
		operandAddr := bus.ReadWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.YIndex)
		incrementWhenPageCrossed(c, operandAddr, res)
		return res
	case Indirect:
		operand := bus.ReadWord(uint16(c.ProgramCounter))
		indirectValue := bus.ReadWord(operand)
		return indirectValue
	case IndirectX:
		operand := bus.Read(uint16(c.ProgramCounter))
		operand += uint8(c.XIndex)
		return bus.ReadWord(uint16(operand))
	case IndirectY:
		operand := bus.Read(uint16(c.ProgramCounter))
		word := bus.ReadWord(uint16(operand))
		res := word + uint16(c.YIndex)
		return res
	case IndirectY1:
		operand := bus.Read(uint16(c.ProgramCounter))
		word := bus.ReadWord(uint16(operand))
		res := word + uint16(c.YIndex)
		incrementWhenPageCrossed(c, word, res)
		return res
	default:
		panic("Addressing mode not implemented")
	}
}

func lda(c *Cpu, bus Bus, mode int) {
	operand := bus.Read(c.getOperandAddress(bus, mode))
	c.Accumulator = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func ldx(c *Cpu, bus Bus, mode int) {
	operand := bus.Read(c.getOperandAddress(bus, mode))
	c.XIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func ldy(c *Cpu, bus Bus, mode int) {
	operand := bus.Read(c.getOperandAddress(bus, mode))
	c.YIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func sta(c *Cpu, bus Bus, mode int) {
	addrToFill := c.getOperandAddress(bus, mode)
	bus.Write(uint16(addrToFill), uint8(c.Accumulator))
}

func stx(c *Cpu, bus Bus, mode int) {
	addrToFill := c.getOperandAddress(bus, mode)
	bus.Write(uint16(addrToFill), uint8(c.XIndex))
}

// TODO: impl this func
// should call the interrupt
func brk(c *Cpu, bus Bus, mode int) {
}

func nop(c *Cpu, bus Bus, mode int) {
	// purposely does nothing
}

// TODO: add some tests maybe ?
func rti(c *Cpu, bus Bus, mode int) {
	statusFlags := popStack(c, bus)
	pc := popStack(c, bus)

	c.ProgramCounter = Register16(pc)
	c.Status = Register8(statusFlags)
}

func sty(c *Cpu, bus Bus, mode int) {
	addrToFill := c.getOperandAddress(bus, mode)
	bus.Write(uint16(addrToFill), uint8(c.YIndex))
}

func tax(c *Cpu, bus Bus, mode int) {
	c.XIndex = c.Accumulator
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func tay(c *Cpu, bus Bus, mode int) {
	c.YIndex = c.Accumulator
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func txa(c *Cpu, bus Bus, mode int) {
	c.Accumulator = c.XIndex
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func tya(c *Cpu, bus Bus, mode int) {
	c.Accumulator = c.YIndex
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: Maybe add tests for other mods
func and(c *Cpu, bus Bus, mode int) {
	operand := bus.Read(c.getOperandAddress(bus, mode))
	c.Accumulator &= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func eor(c *Cpu, bus Bus, mode int) {
	operand := bus.Read(c.getOperandAddress(bus, mode))
	c.Accumulator ^= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func aor(c *Cpu, bus Bus, mode int) {
	operand := bus.Read(c.getOperandAddress(bus, mode))
	c.Accumulator |= Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: may need some verif
func bit(c *Cpu, bus Bus, mode int) {
	data := bus.Read(c.getOperandAddress(bus, mode))

	and := uint8(c.Accumulator) & data

	if and == 0 {
		c.Status.Add(Zero)
	} else {
		c.Status.Remove(Zero)
	}

	if data&0b1000_0000 > 0 {
		c.Status.Add(Negative)
	} else {
		c.Status.Remove(Negative)
	}

	if data&0b0100_0000 > 0 {
		c.Status.Add(Verflow)
	} else {
		c.Status.Remove(Verflow)
	}

	// aCopy := c.Accumulator
	// aCopy &= Register8(operand)

	// if aCopy == 0 {
	// 	c.Status.Add(Zero)
	// } else {
	// 	c.Status.Remove(Zero)
	// }

	// if operand&Negative == 0 {
	// 	c.Status.Remove(Negative)
	// } else {
	// 	c.Status.Add(Negative)
	// }

	// if operand&Verflow == 0 {
	// 	c.Status.Remove(Verflow)
	// } else {
	// 	c.Status.Add(Verflow)
	// }
}

func tsx(c *Cpu, bus Bus, mode int) {
	c.XIndex = c.StackPointer
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func txs(c *Cpu, bus Bus, mode int) {
	c.StackPointer = c.XIndex
}

func pushStack(c *Cpu, bus Bus, value uint8) {
	bus.Write(0x0100+uint16(c.StackPointer), value)
	c.StackPointer--
}

func popStack(c *Cpu, bus Bus) uint8 {
	c.StackPointer++
	val := bus.Read(0x0100 + uint16(c.StackPointer))

	return val
}

// TODO: may wrap the stack pointer
func pha(c *Cpu, bus Bus, mode int) {
	pushStack(c, bus, uint8(c.Accumulator))
}

func php(c *Cpu, bus Bus, mode int) {
	flags := c.Status
	flags.Add(Break)
	flags.Add(Break2)

	pushStack(c, bus, uint8(c.Status))
}

// TODO: may wrap the stack pointer also
// Should learn about the B flag
func pla(c *Cpu, bus Bus, mode int) {
	val := popStack(c, bus)
	c.Accumulator = Register8(val)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: should go deeper for tests
func plp(c *Cpu, bus Bus, mode int) {
	val := popStack(c, bus)
	c.Status = Register8(val)

	c.Status.Remove(Break)
	c.Status.Add(Break2)
}

func inc(c *Cpu, bus Bus, mode int) {
	addr := c.getOperandAddress(bus, mode)
	b := bus.Read(addr)
	bus.Write(addr, b+1)
	c.updateZeroAndNegativeFlags(Register8(b + 1))
}

func inx(c *Cpu, bus Bus, mode int) {
	c.XIndex++
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func iny(c *Cpu, bus Bus, mode int) {
	c.YIndex++
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func dec(c *Cpu, bus Bus, mode int) {
	addr := c.getOperandAddress(bus, mode)
	b := bus.Read(addr)
	bus.Write(addr, b-1)
	c.updateZeroAndNegativeFlags(Register8(b - 1))
}

func dex(c *Cpu, bus Bus, mode int) {
	c.XIndex--
	c.updateZeroAndNegativeFlags(c.XIndex)
}

func dey(c *Cpu, bus Bus, mode int) {
	c.YIndex--
	c.updateZeroAndNegativeFlags(c.YIndex)
}

func rol(c *Cpu, bus Bus, mode int) {
	if mode == Accumulator {
		isOldBit7Set := c.Accumulator&0b1000_0000 != 0
		c.Accumulator <<= 1
		c.Accumulator.Set(0, c.Status.Has(Carry))
		c.Status.Set(0, isOldBit7Set)
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(bus, mode)
		value := bus.Read(addr)

		isOldBit7Set := value&0b1000_0000 != 0
		value <<= 1
		if c.Status.Has(Carry) {
			value |= 0b0000_0001
		} else {
			value &^= 0b0000_0001
		}
		c.Status.Set(0, isOldBit7Set)

		bus.Write(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

// TODO: Maybe adding more tests
func ror(c *Cpu, bus Bus, mode int) {
	if mode == Accumulator {
		isOldBit0Set := c.Accumulator&0b0000_0001 != 0
		c.Accumulator >>= 1
		c.Accumulator.Set(7, c.Status.Has(Carry))
		c.Status.Set(0, isOldBit0Set)
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(bus, mode)
		value := bus.Read(addr)

		isOldBit0Set := value&0b0000_0001 != 0
		value >>= 1

		if c.Status.Has(Carry) {
			value |= 0b1000_0000
		} else {
			value &^= 0b1000_0000
		}
		c.Status.Set(0, isOldBit0Set)

		bus.Write(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func lsr(c *Cpu, bus Bus, mode int) {
	if mode == Accumulator {
		if c.Accumulator.Has(0b0000_0001) {
			c.Status.Add(Carry)
		} else {
			c.Status.Remove(Carry)
		}

		c.Accumulator >>= 1
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(bus, mode)
		value := bus.Read(addr)

		if value&0b0000_0001 == 0 {
			c.Status.Remove(Carry)
		} else {
			c.Status.Add(Carry)
		}

		value >>= 1

		bus.Write(addr, value)
		c.updateZeroAndNegativeFlags(Register8(value))
	}
}

func asl(c *Cpu, bus Bus, mode int) {
	if mode == Accumulator {
		if c.Accumulator.Has(0b1000_0000) {
			c.Status.Add(Carry)
		} else {
			c.Status.Remove(Carry)
		}

		c.Accumulator <<= 1
		c.updateZeroAndNegativeFlags(c.Accumulator)
	} else {
		addr := c.getOperandAddress(bus, mode)
		value := bus.Read(addr)

		if value&Negative == 0 {
			c.Status.Remove(Carry)
		} else {
			c.Status.Add(Carry)
		}

		value <<= 1

		bus.Write(addr, value)
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
func adc(c *Cpu, bus Bus, mode int) {
	addr := c.getOperandAddress(bus, mode)
	data := bus.Read(addr)
	c.addToRegisterA(data)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: add tests
func sbc(c *Cpu, bus Bus, mode int) {
	addr := c.getOperandAddress(bus, mode)
	data := bus.Read(addr)
	data = -data
	c.addToRegisterA(data)
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

// TODO: add tests
func (c *Cpu) compare(bus Bus, compareWith uint8) {

	if compareWith <= uint8(c.Accumulator) {
		c.Status.Add(Carry)
	} else {
		c.Status.Remove(Carry)
	}

}

func cmp(c *Cpu, bus Bus, mode int) {
	data := bus.Read(c.getOperandAddress(bus, mode))
	c.compare(bus, data)
	c.updateZeroAndNegativeFlags(c.Accumulator - Register8(data))
}

func cpx(c *Cpu, bus Bus, mode int) {
	data := bus.Read(c.getOperandAddress(bus, mode))
	c.compare(bus, data)
	c.updateZeroAndNegativeFlags(c.XIndex - Register8(data))
}

func cpy(c *Cpu, bus Bus, mode int) {
	data := bus.Read(c.getOperandAddress(bus, mode))
	c.compare(bus, data)
	c.updateZeroAndNegativeFlags(c.YIndex - Register8(data))
}

func jsr(c *Cpu, bus Bus, mode int) {
	operand := c.getOperandAddress(bus, mode)
	pushStack(c, bus, uint8((c.ProgramCounter+2)>>8))
	pushStack(c, bus, uint8((c.ProgramCounter + 2)))
	c.ProgramCounter = Register16(operand)
}

// TODO: check if the order of operation is right
func rts(c *Cpu, bus Bus, mode int) {
	var lo uint16 = uint16(popStack(c, bus))
	var hi uint16 = uint16(popStack(c, bus))

	// fmt.Printf("before rts: %04x\n", c.ProgramCounter)
	c.ProgramCounter = Register16(((hi << 8) | lo))
	// fmt.Printf("after  rts: %04x\n", c.ProgramCounter)
}

func branch(c *Cpu, bus Bus, condition bool) {
	if condition {
		var jump int8 = int8(bus.Read(uint16(c.ProgramCounter)))
		c.ProgramCounter += Register16(jump)
	}
}

// TODO: may check if branching works correctly to the TODO
func bcc(c *Cpu, bus Bus, mode int) {
	branch(c, bus, !c.Status.Has(Carry))
}

func bcs(c *Cpu, bus Bus, mode int) {
	branch(c, bus, c.Status.Has(Carry))
}

func beq(c *Cpu, bus Bus, mode int) {
	branch(c, bus, c.Status.Has(Zero))
}

func bmi(c *Cpu, bus Bus, mode int) {
	branch(c, bus, c.Status.Has(Negative))
}

func bne(c *Cpu, bus Bus, mode int) {
	branch(c, bus, !c.Status.Has(Zero))
}

func bpl(c *Cpu, bus Bus, mode int) {
	branch(c, bus, !c.Status.Has(Negative))
}

func bvc(c *Cpu, bus Bus, mode int) {
	branch(c, bus, !c.Status.Has(Verflow))
}

func bvs(c *Cpu, bus Bus, mode int) {
	branch(c, bus, c.Status.Has(Verflow))
}

// TODO: check if the increment doesn't happen in the same instruction cycle
func jmp(c *Cpu, bus Bus, mode int) {
	operand := c.getOperandAddress(bus, mode)
	c.ProgramCounter = Register16(operand)
}

func clc(c *Cpu, bus Bus, mode int) {
	c.Status.Remove(Carry)
}

func cld(c *Cpu, bus Bus, mode int) {
	c.Status.Remove(Decimal)
}

func cli(c *Cpu, bus Bus, mode int) {
	c.Status.Remove(Interrupt)
}

func clv(c *Cpu, bus Bus, mode int) {
	c.Status.Remove(Verflow)
}

func sec(c *Cpu, bus Bus, mode int) {
	c.Status.Add(Carry)
}

func sed(c *Cpu, bus Bus, mode int) {
	c.Status.Add(Decimal)
}

func sei(c *Cpu, bus Bus, mode int) {
	c.Status.Add(Interrupt)
}

func (c *Cpu) interpret(opcode uint8, bus Bus) {
	opc, ok := Opcodes[opcode]
	if !ok {
		panic("UNKOWN OPCODE")
	}

	opc.Operation(c, bus, opc.Mode)

	c.Cycle += opc.Cycles
	c.ProgramCounter += Register16(opc.ByteSize - 1)
}

func (c *Cpu) Reset(bus Bus) {
	c.Accumulator = 0
	c.XIndex = 0
	c.YIndex = 0
	c.StackPointer = 0xFF
	c.Status = 0

	c.ProgramCounter = Register16(bus.ReadWord(0xFFFC))
}

func (c *Cpu) Step(bus Bus) bool {
	opcode := bus.Read(uint16(c.ProgramCounter))
	c.ProgramCounter++

	c.interpret(opcode, bus)
	return opcode == BRK_IMP
}

func (c *Cpu) Run(bus Bus) {
	for {
		shouldExit := c.Step(bus)

		if shouldExit {
			break
		}
	}
}
