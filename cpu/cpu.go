package cpu

type Cpu struct {
	Cycle int
	Registers
}

/*
func (c *Cpu) getNextByte(memory Memory) uint8 {
	opcode := memory.readByte(uint16(c.ProgramCounter))
	c.ProgramCounter++
	return opcode
}
*/

/*
func (c *Cpu) ldaImm(memory Memory) {
	operand := c.getNextByte(memory)
	c.Accumulator = Register8(operand)

	c.updateZeroAndNegativeFlags(c.Accumulator)

}

func (c *Cpu) taxImp() {
	c.XIndex = c.Accumulator

	c.updateZeroAndNegativeFlags(c.XIndex)
}

func (c *Cpu) inxImp() {
	c.XIndex += 1

	c.updateZeroAndNegativeFlags(c.XIndex)
}
*/

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

// TODO: add private func to clear this
func (c *Cpu) getOperandAddress(mem *Memory, mode int) uint16 {
	switch mode {
	case Implied:
		panic("Implied mode not supported")
	case Immediate:
		return uint16(c.ProgramCounter)
	case ZeroPage:
		return uint16(mem.readByte(uint16(c.ProgramCounter)))
	case ZeroPageX:
		operandAddr := mem.readByte(uint16(c.ProgramCounter))
		operandAddr += uint8(c.XIndex)
		return uint16(operandAddr)
	case ZeroPageY:
		operandAddr := mem.readByte(uint16(c.ProgramCounter))
		operandAddr += uint8(c.YIndex)
		return uint16(operandAddr)
	case Absolute:
		return mem.readWord(uint16(c.ProgramCounter))
	case AbsoluteX:
		operandAddr := mem.readWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.XIndex)
		return res
	case AbsoluteX1:
		operandAddr := mem.readWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.XIndex)
		if (res >> 8) != (operandAddr >> 8) {
			c.Cycle++
		}
		return res
	case AbsoluteY:
		operandAddr := mem.readWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.YIndex)
		return res
	case AbsoluteY1:
		operandAddr := mem.readWord(uint16(c.ProgramCounter))
		res := operandAddr + uint16(c.YIndex)
		if (res >> 8) != (operandAddr >> 8) {
			c.Cycle++
		}
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
		if (res >> 8) != (word >> 8) {
			c.Cycle++
		}
		return res
	default:
		panic("Addressing mode not implemented")
	}
}

func (c *Cpu) lda(mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.Accumulator = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func (c *Cpu) ldx(mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.XIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func (c *Cpu) ldy(mem *Memory, mode int) {
	operand := mem.readByte(c.getOperandAddress(mem, mode))
	c.YIndex = Register8(operand)
	c.updateZeroAndNegativeFlags(Register8(operand))
}

func (c *Cpu) sta(mem *Memory, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.writeByte(uint16(addrToFill), uint8(c.Accumulator))
}

func (c *Cpu) stx(mem *Memory, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.writeByte(uint16(addrToFill), uint8(c.XIndex))
}

func (c *Cpu) sty(mem *Memory, mode int) {
	addrToFill := c.getOperandAddress(mem, mode)
	mem.writeByte(uint16(addrToFill), uint8(c.YIndex))
}

func (c *Cpu) Reset(mem *Memory) {
	c.Accumulator = 0
	c.XIndex = 0
	c.YIndex = 0
	c.StackPointer = 0
	c.Status = 0

	c.ProgramCounter = Register16(mem.readWord(0xFFFC))
}

// TODO: maybe add EACH by group string like "STA" "LDA"
func (c *Cpu) interpret(opcode uint8, memory *Memory) bool {
	opc := Opcodes[opcode]

	switch opcode {
	case LDA_IMM, LDA_ZER, LDA_ZRX, LDA_ABS, LDA_ABX, LDA_ABY, LDA_IDX, LDA_IDY:
		c.lda(memory, opc.Mode)
	case LDX_IMM, LDX_ZER, LDX_ZRY, LDX_ABS, LDX_ABY:
		c.ldx(memory, opc.Mode)
	case LDY_IMM, LDY_ZER, LDY_ZRX, LDY_ABS, LDY_ABX:
		c.ldy(memory, opc.Mode)
	case STA_ZER, STA_ZRX, STA_ABS, STA_ABX, STA_ABY, STA_IDX, STA_IDY:
		c.sta(memory, opc.Mode)
	case STX_ZER, STX_ZRY, STX_ABS:
		c.stx(memory, opc.Mode)
	case STY_ZER, STY_ZRX, STY_ABS:
		c.sty(memory, opc.Mode)
	case BRK_IMP:
		c.Cycle += 7
		c.ProgramCounter += Register16(opc.ByteSize - 1)
		return true
	default:
		panic("Unknown opcode!")
	}

	c.Cycle += opc.Cycles
	c.ProgramCounter += Register16(opc.ByteSize - 1)

	return false
}

func (c *Cpu) Run(memory *Memory) {
	for {
		opcode := memory.readByte(uint16(c.ProgramCounter))
		c.ProgramCounter++

		if c.interpret(opcode, memory) {
			break
		}
	}
}
