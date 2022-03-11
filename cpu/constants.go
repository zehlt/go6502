package cpu

const (
	Implied = iota
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Absolute
	AbsoluteX
	AbsoluteX1
	AbsoluteY
	AbsoluteY1
	IndirectX
	IndirectY
	IndirectY1
)

const (
	LDA_IMM = 0xA9
	LDA_ZER = 0xA5
	LDA_ZRX = 0xB5
	LDA_ABS = 0xAD
	LDA_ABX = 0xBD
	LDA_ABY = 0xB9
	LDA_IDX = 0xA1
	LDA_IDY = 0xB1

	LDX_IMM = 0xA2
	LDX_ZER = 0xA6
	LDX_ZRY = 0xB6
	LDX_ABS = 0xAE
	LDX_ABY = 0xBE

	LDY_IMM = 0xA0
	LDY_ZER = 0xA4
	LDY_ZRX = 0xB4
	LDY_ABS = 0xAC
	LDY_ABX = 0xBC

	STA_ZER = 0x85
	STA_ZRX = 0x95
	STA_ABS = 0x8D
	STA_ABX = 0x9D
	STA_ABY = 0x99
	STA_IDX = 0x81
	STA_IDY = 0x91

	STX_ZER = 0x86
	STX_ZRY = 0x96
	STX_ABS = 0x8E

	STY_ZER = 0x84
	STY_ZRX = 0x94
	STY_ABS = 0x8C

	BRK_IMP = 0x00
	TAX_IMP = 0xAA
	INX_IMP = 0xE8
)

type Opcode struct {
	Code     uint8
	ByteSize int
	Cycles   int
	Mode     int
}

var Opcodes = map[uint8]Opcode{
	LDA_IMM: {Code: LDA_IMM, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDA_ZER: {Code: LDA_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDA_ZRX: {Code: LDA_ZRX, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	LDA_ABS: {Code: LDA_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDA_ABX: {Code: LDA_ABX, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteX1},
	LDA_ABY: {Code: LDA_ABY, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteY1},
	LDA_IDX: {Code: LDA_IDX, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	LDA_IDY: {Code: LDA_IDY, ByteSize: 2, Cycles: 5 /*+1 crossed*/, Mode: IndirectY1},

	LDX_IMM: {Code: LDX_IMM, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDX_ZER: {Code: LDX_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDX_ZRY: {Code: LDX_ZRY, ByteSize: 2, Cycles: 4, Mode: ZeroPageY},
	LDX_ABS: {Code: LDX_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDX_ABY: {Code: LDX_ABY, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteY1},

	LDY_IMM: {Code: LDY_IMM, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDY_ZER: {Code: LDY_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDY_ZRX: {Code: LDY_ZRX, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	LDY_ABS: {Code: LDY_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDY_ABX: {Code: LDY_ABX, ByteSize: 3, Cycles: 4 /*+1 crossed*/, Mode: AbsoluteX1},

	STA_ZER: {Code: STA_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	STA_ZRX: {Code: STA_ZRX, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	STA_ABS: {Code: STA_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},
	STA_ABX: {Code: STA_ABX, ByteSize: 3, Cycles: 5, Mode: AbsoluteX},
	STA_ABY: {Code: STA_ABY, ByteSize: 3, Cycles: 5, Mode: AbsoluteY},
	STA_IDX: {Code: STA_IDX, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	STA_IDY: {Code: STA_IDY, ByteSize: 2, Cycles: 6, Mode: IndirectY},

	STX_ZER: {Code: STX_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	STX_ZRY: {Code: STX_ZRY, ByteSize: 2, Cycles: 4, Mode: ZeroPageY},
	STX_ABS: {Code: STX_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},

	STY_ZER: {Code: STY_ZER, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	STY_ZRX: {Code: STY_ZRX, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	STY_ABS: {Code: STY_ABS, ByteSize: 3, Cycles: 4, Mode: Absolute},

	BRK_IMP: {Code: BRK_IMP, ByteSize: 1, Cycles: 7, Mode: Implied},
}

const (
	Carry     = 0b0000_0001
	Zero      = 0b0000_0010
	Interrupt = 0b0000_0100
	Decimal   = 0b0000_1000
	Break     = 0b0001_0000
	Verflow   = 0b0100_0000
	Negative  = 0b1000_0000
)
