package cpu

import (
	"testing"

	"github.com/zehlt/go6502/asrt"
)

func TestLdaImmediatePositiveValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x10, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x10))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestLdaImmediateZeroValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestLdaImmediateNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_IMM, 0x80, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x80))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPagePositiveValue(t *testing.T) {
	memory := Memory{
		LDA_ZER, 0xAA, BRK_IMP,
	}
	memory[0xAA] = 0x67

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x67))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ZER, 0xAA, BRK_IMP,
	}
	memory[0xAA] = 0xDF

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xDF))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ZER, 0xAA, BRK_IMP,
	}
	memory[0xAA] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXPositiveValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x20, BRK_IMP,
	}
	memory[0x30] = 0x79

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x79))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x20, BRK_IMP,
	}
	memory[0x30] = 0xEF

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xEF))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x20, BRK_IMP,
	}
	memory[0x30] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaZeroPageXWrappingValue(t *testing.T) {
	memory := Memory{
		LDA_ZRX, 0x80, BRK_IMP,
	}
	memory[0x7F] = 0x22

	cpu := Cpu{}
	cpu.XIndex = 0xFF
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x22))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		LDA_ABS, 0xFE, 0x01, BRK_IMP,
	}
	memory[0x01fe] = 0x33

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x33))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ABS, 0xFE, 0x01, BRK_IMP,
	}
	memory[0x01fe] = 0xAA

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xAA))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ABS, 0xFE, 0x01, BRK_IMP,
	}
	memory[0x01fe] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXPositiveValueCrossedPage(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0xFF, 0x01, BRK_IMP,
	}
	// 0x01ff+0x0010 = 0x020f
	memory[0x020f] = 0x66

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x66))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXPositiveValueNotCrossedPage(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0x00, 0x01, BRK_IMP,
	}
	memory[0x0110] = 0x77

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x77))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0x20, 0x01, BRK_IMP,
	}
	memory[0x0130] = 0x90

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x90))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteXZeroValue(t *testing.T) {
	memory := Memory{
		LDA_ABX, 0x00, 0x01, BRK_IMP,
	}
	memory[0x0110] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteYPositiveValueNotCrossed(t *testing.T) {
	memory := Memory{
		LDA_ABY, 0x10, 0x0A, BRK_IMP,
	}
	memory[0x0A30] = 0x45

	cpu := Cpu{}
	cpu.YIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x45))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaAbsoluteYNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_ABY, 0x10, 0x0A, BRK_IMP,
	}
	memory[0x0A30] = 0xA5

	cpu := Cpu{}
	cpu.YIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xA5))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_ABY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectXPositiveValue(t *testing.T) {
	memory := Memory{
		LDA_IDX, 0x10, BRK_IMP,
	}
	memory[0x0015] = 0x07
	memory[0x0016] = 0x09

	memory[0x0907] = 0x79

	cpu := Cpu{}
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x79))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectXNegativeValue(t *testing.T) {
	memory := Memory{
		LDA_IDX, 0x10, BRK_IMP,
	}
	memory[0x0015] = 0x07
	memory[0x0016] = 0x09

	memory[0x0907] = 0xFF

	cpu := Cpu{}
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xFF))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectXZeroValue(t *testing.T) {
	memory := Memory{
		LDA_IDX, 0x10, BRK_IMP,
	}
	memory[0x0015] = 0x07
	memory[0x0016] = 0x09

	memory[0x0907] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectYPositiveValueNotCrossed(t *testing.T) {
	memory := Memory{
		LDA_IDY, 0x20, BRK_IMP,
	}
	memory[0x0020] = 0x03
	memory[0x0021] = 0x07

	memory[0x0704] = 0x55

	cpu := Cpu{}
	cpu.YIndex = 0x01
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x55))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdaIndirectYPositiveValueCrossed(t *testing.T) {
	memory := Memory{
		LDA_IDY, 0x20, BRK_IMP,
	}
	memory[0x0020] = 0xFF
	memory[0x0021] = 0x07

	memory[0x080f] = 0x66

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x66))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDA_IDY].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxImmediatePositiveValue(t *testing.T) {
	memory := Memory{
		LDX_IMM, 0x50, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x50))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxImmediateNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_IMM, 0x81, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x81))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxImmediateZeroValue(t *testing.T) {
	memory := Memory{
		LDX_IMM, 0x00, BRK_IMP,
	}
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPagePositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0x22

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x22))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPageNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0xAE

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xAE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

// TODO: assure that wrap up doesn't cause any pb
func TestLdxZeroPageZeroValue(t *testing.T) {
	memory := Memory{
		LDX_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPageYPositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ZRY, 0x50, BRK_IMP,
	}

	memory[0x0060] = 0x25

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x25))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZRY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxZeroPageYNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_ZRY, 0x50, BRK_IMP,
	}

	memory[0x0060] = 0xDE

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xDE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ZRY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ABS, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2050] = 0x10

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x10))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		LDX_ABS, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2050] = 0xCE

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xCE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		LDX_ABS, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2050] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteYPositiveValue(t *testing.T) {
	memory := Memory{
		LDX_ABY, 0x50, 0x20, BRK_IMP,
	}

	memory[0x2061] = 0x35

	cpu := Cpu{}
	cpu.YIndex = 0x11
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x35))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABY].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteYPositiveValueCrossedPage(t *testing.T) {
	memory := Memory{
		LDX_ABY, 0xFF, 0x20, BRK_IMP,
	}

	memory[0x210F] = 0x35

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x35))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABY].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdxAbsoluteYPositiveValueCrossedPageAndNegative(t *testing.T) {
	memory := Memory{
		LDX_ABY, 0xFF, 0x20, BRK_IMP,
	}

	memory[0x210F] = 0xCC

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xCC))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDX_ABY].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyImmediatePositiveValue(t *testing.T) {
	memory := Memory{
		LDY_IMM, 0x25, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x25))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyImmediateNegativeValue(t *testing.T) {
	memory := Memory{
		LDY_IMM, 0xBE, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0xBE))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyImmediateZeroValue(t *testing.T) {
	memory := Memory{
		LDY_IMM, 0x00, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdyZeroPagePositiveValue(t *testing.T) {
	memory := Memory{
		LDY_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0x63

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x63))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyZeroPageZeroValue(t *testing.T) {
	memory := Memory{
		LDY_ZER, 0x45, BRK_IMP,
	}

	memory[0x0045] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdyZeroPageXPositiveValue(t *testing.T) {
	memory := Memory{
		LDY_ZRX, 0x45, BRK_IMP,
	}

	memory[0x0055] = 0x22

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x22))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyZeroPageXNegativeValue(t *testing.T) {
	memory := Memory{
		LDY_ZRX, 0x45, BRK_IMP,
	}

	memory[0x0055] = 0xF3

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0xF3))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		LDY_ABS, 0xa5, 0x2f, BRK_IMP,
	}

	memory[0x2fa5] = 0x78

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x78))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		LDY_ABS, 0xa5, 0x2f, BRK_IMP,
	}

	memory[0x2fa5] = 0xaf

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0xaf))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		LDY_ABS, 0xa5, 0x2f, BRK_IMP,
	}

	memory[0x2fa5] = 0x00

	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestLdyAbsoluteXPositiveValueNotCrossed(t *testing.T) {
	memory := Memory{
		LDY_ABX, 0x50, 0x2f, BRK_IMP,
	}

	memory[0x2f70] = 0x78

	cpu := Cpu{}
	cpu.XIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x78))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyAbsoluteXPositiveValueCrossed(t *testing.T) {
	memory := Memory{
		LDY_ABX, 0xff, 0x2f, BRK_IMP,
	}

	memory[0x300f] = 0x78

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x78))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ABX].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
}

func TestLdyAbsoluteXPositiveValueCrossedAndZero(t *testing.T) {
	memory := Memory{
		LDY_ABX, 0xff, 0x2f, BRK_IMP,
	}

	memory[0x300f] = 0x00

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[LDY_ABX].Cycles+Opcodes[BRK_IMP].Cycles+1)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestStaZeroPage(t *testing.T) {
	memory := Memory{
		STA_ZER, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xAA
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0x0050]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStaZeroPageX(t *testing.T) {
	memory := Memory{
		STA_ZRX, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x20
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0x0070]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStaAbsolute(t *testing.T) {
	memory := Memory{
		STA_ABS, 0x50, 0xFA, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xBC
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0xFA50]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStaAbsoluteX(t *testing.T) {
	memory := Memory{
		STA_ABX, 0x50, 0xFA, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xBC
	cpu.XIndex = 0x30
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0xFA80]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStaAbsoluteY(t *testing.T) {
	memory := Memory{
		STA_ABY, 0x50, 0xFA, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xBC
	cpu.YIndex = 0x30
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0xFA80]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_ABY].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStaIndirectX(t *testing.T) {
	memory := Memory{
		STA_IDX, 0x20, BRK_IMP,
	}

	memory[0x0025] = 0x33
	memory[0x0026] = 0xA7

	cpu := Cpu{}
	cpu.Accumulator = 0xBC
	cpu.XIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0xA733]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_IDX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStaIndirectY(t *testing.T) {
	memory := Memory{
		STA_IDY, 0x20, BRK_IMP,
	}

	memory[0x0020] = 0x33
	memory[0x0021] = 0xA7

	cpu := Cpu{}
	cpu.Accumulator = 0xBC
	cpu.YIndex = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(memory[0xA738]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_IDY].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStxZeroPage(t *testing.T) {
	memory := Memory{
		STX_ZER, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0xAA
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(memory[0x50]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STX_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStxZeroPageY(t *testing.T) {
	memory := Memory{
		STX_ZRY, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x20
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(memory[0x060]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STX_ZRY].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStxAbsolute(t *testing.T) {
	memory := Memory{
		STX_ABS, 0x50, 0xFA, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0xBC
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(memory[0xFA50]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STA_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStyZeroPage(t *testing.T) {
	memory := Memory{
		STY_ZER, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0xAA
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(memory[0x50]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STY_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStyZeroPageX(t *testing.T) {
	memory := Memory{
		STY_ZRX, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x20
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(memory[0x060]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STY_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestStyAbsolute(t *testing.T) {
	memory := Memory{
		STY_ABS, 0x50, 0xFA, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0xBC
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(memory[0xFA50]))
	asrt.Equal(t, cpu.Cycle, Opcodes[STY_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTaxPositiveValue(t *testing.T) {
	memory := Memory{
		TAX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x10
	cpu.XIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.XIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TAX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTaxNegativeValue(t *testing.T) {
	memory := Memory{
		TAX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xCE
	cpu.XIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.XIndex)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TAX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTaxZeroValue(t *testing.T) {
	memory := Memory{
		TAX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x00
	cpu.XIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.XIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestTayPositiveValue(t *testing.T) {
	memory := Memory{
		TAY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x10
	cpu.YIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TAY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTayNegativeValue(t *testing.T) {
	memory := Memory{
		TAY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xCE
	cpu.YIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TAY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTayZeroValue(t *testing.T) {
	memory := Memory{
		TAY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x00
	cpu.YIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestTxaPositiveValue(t *testing.T) {
	memory := Memory{
		TXA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.XIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TXA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTxaNegativeValue(t *testing.T) {
	memory := Memory{
		TXA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0xCE
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.XIndex)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TXA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTxaZeroValue(t *testing.T) {
	memory := Memory{
		TXA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x00
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.XIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestTyaPositiveValue(t *testing.T) {
	memory := Memory{
		TYA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TYA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTyaNegativeValue(t *testing.T) {
	memory := Memory{
		TYA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0xCE
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TYA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTyaZeroValue(t *testing.T) {
	memory := Memory{
		TYA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x00
	cpu.Accumulator = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestIncPositiveValue(t *testing.T) {
	memory := Memory{
		INC_ZER, 0x10, BRK_IMP,
	}

	memory[0x10] = 0x20
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x10], uint8(0x21))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncNegativeValue(t *testing.T) {
	memory := Memory{
		INC_ZER, 0x10, BRK_IMP,
	}

	memory[0x10] = 0x7F
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x10], uint8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncZeroValue(t *testing.T) {
	memory := Memory{
		INC_ZER, 0x10, BRK_IMP,
	}

	memory[0x10] = 0xFF
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x10], uint8(0x00))
	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestIncZeroPageXPositiveValue(t *testing.T) {
	memory := Memory{
		INC_ZRX, 0x10, BRK_IMP,
	}

	memory[0x30] = 0x20
	cpu := Cpu{}
	cpu.XIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x30], uint8(0x21))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncZeroPageXNegativeValue(t *testing.T) {
	memory := Memory{
		INC_ZRX, 0x10, BRK_IMP,
	}

	memory[0x30] = 0x7F
	cpu := Cpu{}
	cpu.XIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x30], uint8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncZeroPageXZeroValue(t *testing.T) {
	memory := Memory{
		INC_ZRX, 0x10, BRK_IMP,
	}

	memory[0x30] = 0xFF
	cpu := Cpu{}
	cpu.XIndex = 0x20
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x30], uint8(0x00))
	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestIncAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		INC_ABS, 0x10, 0xAE, BRK_IMP,
	}

	memory[0xAE10] = 0x20
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE10], uint8(0x21))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		INC_ABS, 0x10, 0xAE, BRK_IMP,
	}

	memory[0xAE10] = 0x7F
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE10], uint8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		INC_ABS, 0x10, 0xAE, BRK_IMP,
	}

	memory[0xAE10] = 0xFF
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE10], uint8(0x00))
	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestIncAbsoluteXPositiveValue(t *testing.T) {
	memory := Memory{
		INC_ABX, 0x10, 0xAE, BRK_IMP,
	}

	memory[0xAE20] = 0x03
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE20], uint8(0x04))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncAbsoluteXNegativeValue(t *testing.T) {
	memory := Memory{
		INC_ABX, 0x10, 0xAE, BRK_IMP,
	}

	memory[0xAE20] = 0x7F
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE20], uint8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INC_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestIncAbsoluteXZeroValue(t *testing.T) {
	memory := Memory{
		INC_ABX, 0x10, 0xAE, BRK_IMP,
	}

	memory[0xAE20] = 0xFF
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE20], uint8(0x00))
	asrt.Equal(t, cpu.Accumulator, cpu.YIndex)
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestInxImpliedPositiveValue(t *testing.T) {
	memory := Memory{
		INX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x12
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x13))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestInxImpliedNegativeValue(t *testing.T) {
	memory := Memory{
		INX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x7F
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestInxImpliedZeroValue(t *testing.T) {
	memory := Memory{
		INX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0xFF
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestInyImpliedPositiveValue(t *testing.T) {
	memory := Memory{
		INY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x12
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x13))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestInyImpliedNegativeValue(t *testing.T) {
	memory := Memory{
		INY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x7F
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestInyImpliedZeroValue(t *testing.T) {
	memory := Memory{
		INY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0xFF
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[INY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecZeroPagePositiveValue(t *testing.T) {
	memory := Memory{
		DEC_ZER, 0xAE, BRK_IMP,
	}

	memory[0xAE] = 0x15
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE], uint8(0x14))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecZeroPageNegativeValue(t *testing.T) {
	memory := Memory{
		DEC_ZER, 0xAE, BRK_IMP,
	}

	memory[0xAE] = 0x81
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE], uint8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecZeroPageZeroValue(t *testing.T) {
	memory := Memory{
		DEC_ZER, 0xAE, BRK_IMP,
	}

	memory[0xAE] = 0x01
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xAE], uint8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecZeroPageXPositiveValue(t *testing.T) {
	memory := Memory{
		DEC_ZRX, 0x5A, BRK_IMP,
	}

	memory[0x6A] = 0x15
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x6A], uint8(0x14))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecZeroPageXNagativeValue(t *testing.T) {
	memory := Memory{
		DEC_ZRX, 0x5A, BRK_IMP,
	}

	memory[0x6A] = 0x81
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x6A], uint8(0x80))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecZeroPageXZeroValue(t *testing.T) {
	memory := Memory{
		DEC_ZRX, 0x5A, BRK_IMP,
	}

	memory[0x6A] = 0x01
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0x6A], uint8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ZRX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecAbsolutePositiveValue(t *testing.T) {
	memory := Memory{
		DEC_ABS, 0x5A, 0xEA, BRK_IMP,
	}

	memory[0xEA5A] = 0x15
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xEA5A], uint8(0x14))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecAbsoluteNegativeValue(t *testing.T) {
	memory := Memory{
		DEC_ABS, 0x5A, 0xEA, BRK_IMP,
	}

	memory[0xEA5A] = 0x90
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xEA5A], uint8(0x8F))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecAbsoluteZeroValue(t *testing.T) {
	memory := Memory{
		DEC_ABS, 0x5A, 0xEA, BRK_IMP,
	}

	memory[0xEA5A] = 0x01
	cpu := Cpu{}
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xEA5A], uint8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ABS].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecAbsoluteXPositiveValue(t *testing.T) {
	memory := Memory{
		DEC_ABX, 0x74, 0xEA, BRK_IMP,
	}

	memory[0xEA84] = 0x15
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xEA84], uint8(0x14))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecAbsoluteXNegativeValue(t *testing.T) {
	memory := Memory{
		DEC_ABX, 0x74, 0xEA, BRK_IMP,
	}

	memory[0xEA84] = 0xDE
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xEA84], uint8(0xDD))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDecAbsoluteXZeroValue(t *testing.T) {
	memory := Memory{
		DEC_ABX, 0x74, 0xEA, BRK_IMP,
	}

	memory[0xEA84] = 0x01
	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, memory[0xEA84], uint8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEC_ABX].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDexImpliedPositiveValue(t *testing.T) {
	memory := Memory{
		DEX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x0F))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDexImpliedNegativeValue(t *testing.T) {
	memory := Memory{
		DEX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0xFD))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDexImpliedZeroValue(t *testing.T) {
	memory := Memory{
		DEX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x01
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDeyImpliedPositiveValue(t *testing.T) {
	memory := Memory{
		DEY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x0F))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDeyImpliedNegativeValue(t *testing.T) {
	memory := Memory{
		DEY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0xFE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0xFD))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestDeyImpliedZeroValue(t *testing.T) {
	memory := Memory{
		DEY_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.YIndex = 0x01
	cpu.Run(&memory)

	asrt.Equal(t, cpu.YIndex, Register8(0x00))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[DEY_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestClcImplied(t *testing.T) {
	memory := Memory{
		CLC_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Status.Add(Carry)
	asrt.True(t, cpu.Status.Has(Carry))
	cpu.Run(&memory)
	asrt.False(t, cpu.Status.Has(Carry))

	asrt.Equal(t, cpu.Cycle, Opcodes[CLC_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestClcImpliedWithOtherFlagSet(t *testing.T) {
	memory := Memory{
		CLC_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Status.Add(Carry)
	cpu.Status.Add(Negative)
	cpu.Status.Add(Zero)
	asrt.True(t, cpu.Status.Has(Carry))
	cpu.Run(&memory)
	asrt.False(t, cpu.Status.Has(Carry))

	asrt.Equal(t, cpu.Cycle, Opcodes[CLC_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestCldImplied(t *testing.T) {
	memory := Memory{
		CLD_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Status.Add(Decimal)
	asrt.True(t, cpu.Status.Has(Decimal))
	cpu.Run(&memory)
	asrt.False(t, cpu.Status.Has(Decimal))

	asrt.Equal(t, cpu.Cycle, Opcodes[CLD_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestCldImpliedWithOtherFlagSet(t *testing.T) {
	memory := Memory{
		CLD_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Status.Add(Decimal)
	cpu.Status.Add(Negative)
	cpu.Status.Add(Zero)
	asrt.True(t, cpu.Status.Has(Decimal))
	cpu.Run(&memory)
	asrt.False(t, cpu.Status.Has(Decimal))

	asrt.Equal(t, cpu.Cycle, Opcodes[CLD_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestCliImplied(t *testing.T) {
	memory := Memory{
		CLI_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Status.Add(Interrupt)
	asrt.True(t, cpu.Status.Has(Interrupt))
	cpu.Run(&memory)
	asrt.False(t, cpu.Status.Has(Interrupt))

	asrt.Equal(t, cpu.Cycle, Opcodes[CLI_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestClvImplied(t *testing.T) {
	memory := Memory{
		CLV_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Status.Add(Verflow)
	asrt.True(t, cpu.Status.Has(Verflow))
	cpu.Run(&memory)
	asrt.False(t, cpu.Status.Has(Verflow))

	asrt.Equal(t, cpu.Cycle, Opcodes[CLV_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestSecImplied(t *testing.T) {
	memory := Memory{
		SEC_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Run(&memory)
	asrt.True(t, cpu.Status.Has(Carry))

	asrt.Equal(t, cpu.Cycle, Opcodes[SEC_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestSedImplied(t *testing.T) {
	memory := Memory{
		SED_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Run(&memory)
	asrt.True(t, cpu.Status.Has(Decimal))

	asrt.Equal(t, cpu.Cycle, Opcodes[SED_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestSeiImplied(t *testing.T) {
	memory := Memory{
		SEI_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Run(&memory)
	asrt.True(t, cpu.Status.Has(Interrupt))

	asrt.Equal(t, cpu.Cycle, Opcodes[SEI_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestAndImmediate(t *testing.T) {
	memory := Memory{
		AND_IMM, 0xF0, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xF5
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xF0))
	asrt.Equal(t, cpu.Cycle, Opcodes[AND_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestAndImmediateZero(t *testing.T) {
	memory := Memory{
		AND_IMM, 0x08, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xD1
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[AND_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.False(t, cpu.Status.Has(Negative))
}

func TestAndImmediateNegative(t *testing.T) {
	memory := Memory{
		AND_IMM, 0xD9, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xD1
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xD1))
	asrt.Equal(t, cpu.Cycle, Opcodes[AND_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.True(t, cpu.Status.Has(Negative))
}

func TestEorImmediate(t *testing.T) {
	memory := Memory{
		EOR_IMM, 0xF0, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0xF5
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x05))
	asrt.Equal(t, cpu.Cycle, Opcodes[EOR_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestEorImmediateZero(t *testing.T) {
	memory := Memory{
		EOR_IMM, 0x45, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x45
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x00))
	asrt.Equal(t, cpu.Cycle, Opcodes[EOR_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Zero))
}

func TestAorImmediate(t *testing.T) {
	memory := Memory{
		ORA_IMM, 0x10, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x15))
	asrt.False(t, cpu.Status.Has(Negative))
	asrt.Equal(t, cpu.Cycle, Opcodes[ORA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestAorImmediateNegative(t *testing.T) {
	memory := Memory{
		ORA_IMM, 0xF0, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x05
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0xF5))
	asrt.Equal(t, cpu.Cycle, Opcodes[ORA_IMM].Cycles+Opcodes[BRK_IMP].Cycles)
	asrt.True(t, cpu.Status.Has(Negative))
}

func TestBitZeroPageTwoLastBitSet(t *testing.T) {
	memory := Memory{
		BIT_ZER, 0x30, BRK_IMP,
	}

	cpu := Cpu{}
	memory[0x30] = 0xF0
	cpu.Accumulator = 0x06
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x06))
	asrt.Equal(t, memory[0x30], uint8(0xF0))

	asrt.True(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.True(t, cpu.Status.Has(Verflow))
	asrt.Equal(t, cpu.Cycle, Opcodes[BIT_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestBitZeroPage(t *testing.T) {
	memory := Memory{
		BIT_ZER, 0x50, BRK_IMP,
	}

	cpu := Cpu{}
	memory[0x50] = 0x07
	cpu.Accumulator = 0x06
	cpu.Run(&memory)

	asrt.Equal(t, cpu.Accumulator, Register8(0x06))
	asrt.Equal(t, memory[0x50], uint8(0x07))

	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.False(t, cpu.Status.Has(Verflow))
	asrt.Equal(t, cpu.Cycle, Opcodes[BIT_ZER].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTsxImplied(t *testing.T) {
	memory := Memory{
		TSX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.StackPointer = 0x45
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x45))

	asrt.False(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TSX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTsxImpliedZeroValue(t *testing.T) {
	memory := Memory{
		TSX_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x10
	cpu.StackPointer = 0x00
	cpu.Run(&memory)

	asrt.Equal(t, cpu.XIndex, Register8(0x00))

	asrt.False(t, cpu.Status.Has(Negative))
	asrt.True(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[TSX_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestTxsImplied(t *testing.T) {
	memory := Memory{
		TXS_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.XIndex = 0x45
	cpu.StackPointer = 0x10
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0x45))
	asrt.Equal(t, cpu.Cycle, Opcodes[TXS_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestPhaImplied(t *testing.T) {
	memory := Memory{
		PHA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x33
	cpu.StackPointer = 0xFF
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFE))
	asrt.Equal(t, memory[0xFF], uint8(0x33))
	asrt.Equal(t, cpu.Cycle, Opcodes[PHA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestPhaImpliedMultiplePush(t *testing.T) {
	memory := Memory{
		PHA_IMP, PHA_IMP, PHA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.Accumulator = 0x33
	cpu.StackPointer = 0xFF
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFC))
	asrt.Equal(t, memory[0xFF], uint8(0x33))
	asrt.Equal(t, memory[0xFE], uint8(0x33))
	asrt.Equal(t, memory[0xFD], uint8(0x33))
	asrt.Equal(t, memory[0xFC], uint8(0x00))
}

func TestPhpImplied(t *testing.T) {
	memory := Memory{
		PHP_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.StackPointer = 0xFF
	cpu.Status.Add(Decimal)
	cpu.Status.Add(Carry)
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFE))
	asrt.Equal(t, memory[0xFF], uint8(0x09))
	asrt.Equal(t, cpu.Cycle, Opcodes[PHP_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestPhpImpliedMultiplePush(t *testing.T) {
	memory := Memory{
		PHP_IMP, PHP_IMP, PHP_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.StackPointer = 0xFF
	cpu.Status.Add(Decimal)
	cpu.Status.Add(Carry)
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFC))
	asrt.Equal(t, memory[0xFF], uint8(0x09))
	asrt.Equal(t, memory[0xFE], uint8(0x09))
	asrt.Equal(t, memory[0xFD], uint8(0x09))
	asrt.Equal(t, memory[0xFC], uint8(0x00))
}

func TestPlaImplied(t *testing.T) {
	memory := Memory{
		PLA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.StackPointer = 0xFE
	memory[0xFF] = 0x15
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFF))
	asrt.Equal(t, memory[0xFE], uint8(0x00))
	asrt.Equal(t, memory[0xFF], uint8(0x15))
	asrt.Equal(t, cpu.Accumulator, Register8(0x15))
	asrt.Equal(t, cpu.Cycle, Opcodes[PLA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestPlaImpliedNegative(t *testing.T) {
	memory := Memory{
		PLA_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.StackPointer = 0xFE
	memory[0xFF] = 0xAE
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFF))
	asrt.Equal(t, memory[0xFE], uint8(0x00))
	asrt.Equal(t, memory[0xFF], uint8(0xAE))
	asrt.Equal(t, cpu.Accumulator, Register8(0xAE))
	asrt.True(t, cpu.Status.Has(Negative))
	asrt.False(t, cpu.Status.Has(Zero))
	asrt.Equal(t, cpu.Cycle, Opcodes[PLA_IMP].Cycles+Opcodes[BRK_IMP].Cycles)
}

func TestPlpImplied(t *testing.T) {
	memory := Memory{
		PLP_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.StackPointer = 0xFE
	memory[0xFF] = 0b0101_0101
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFF))
	asrt.Equal(t, cpu.Status, Register8(0b0101_0101))
}

func TestPlpImpliedMultipleTime(t *testing.T) {
	memory := Memory{
		PLP_IMP, PLP_IMP, BRK_IMP,
	}

	cpu := Cpu{}
	cpu.StackPointer = 0xFD
	memory[0xFE] = 0b0101_0101
	memory[0xFF] = 0b1111_0000
	cpu.Run(&memory)

	asrt.Equal(t, cpu.StackPointer, Register8(0xFF))
	asrt.Equal(t, cpu.Status, Register8(0b1111_0000))
}
