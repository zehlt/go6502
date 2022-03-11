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

	TAX_IMP = 0xAA
	TAY_IMP = 0xA8
	TXA_IMP = 0x8A
	TYA_IMP = 0x98

	INC_ZER = 0xE6
	INC_ZRX = 0xF6
	INC_ABS = 0xEE
	INC_ABX = 0xFE
	INX_IMP = 0xE8
	INY_IMP = 0xC8

	DEC_ZER = 0xC6
	DEC_ZRX = 0xD6
	DEC_ABS = 0xCE
	DEC_ABX = 0xDE
	DEX_IMP = 0xCA
	DEY_IMP = 0x88

	BRK_IMP = 0x00

	OTHER = 0xFF
)

type Opcode struct {
	Code      uint8
	ByteSize  int
	Cycles    int
	Mode      int
	Operation func(cpu *Cpu, mem *Memory, mode int)
}

var Opcodes = map[uint8]Opcode{

	// Load Operations
	LDA_IMM: {Code: LDA_IMM, Operation: lda, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDA_ZER: {Code: LDA_ZER, Operation: lda, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDA_ZRX: {Code: LDA_ZRX, Operation: lda, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	LDA_ABS: {Code: LDA_ABS, Operation: lda, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDA_ABX: {Code: LDA_ABX, Operation: lda, ByteSize: 3, Cycles: 4, Mode: AbsoluteX1},
	LDA_ABY: {Code: LDA_ABY, Operation: lda, ByteSize: 3, Cycles: 4, Mode: AbsoluteY1},
	LDA_IDX: {Code: LDA_IDX, Operation: lda, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	LDA_IDY: {Code: LDA_IDY, Operation: lda, ByteSize: 2, Cycles: 5, Mode: IndirectY1},

	LDX_IMM: {Code: LDX_IMM, Operation: ldx, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDX_ZER: {Code: LDX_ZER, Operation: ldx, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDX_ZRY: {Code: LDX_ZRY, Operation: ldx, ByteSize: 2, Cycles: 4, Mode: ZeroPageY},
	LDX_ABS: {Code: LDX_ABS, Operation: ldx, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDX_ABY: {Code: LDX_ABY, Operation: ldx, ByteSize: 3, Cycles: 4, Mode: AbsoluteY1},

	LDY_IMM: {Code: LDY_IMM, Operation: ldy, ByteSize: 2, Cycles: 2, Mode: Immediate},
	LDY_ZER: {Code: LDY_ZER, Operation: ldy, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	LDY_ZRX: {Code: LDY_ZRX, Operation: ldy, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	LDY_ABS: {Code: LDY_ABS, Operation: ldy, ByteSize: 3, Cycles: 4, Mode: Absolute},
	LDY_ABX: {Code: LDY_ABX, Operation: ldy, ByteSize: 3, Cycles: 4, Mode: AbsoluteX1},

	// Store Operations
	STA_ZER: {Code: STA_ZER, Operation: sta, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	STA_ZRX: {Code: STA_ZRX, Operation: sta, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	STA_ABS: {Code: STA_ABS, Operation: sta, ByteSize: 3, Cycles: 4, Mode: Absolute},
	STA_ABX: {Code: STA_ABX, Operation: sta, ByteSize: 3, Cycles: 5, Mode: AbsoluteX},
	STA_ABY: {Code: STA_ABY, Operation: sta, ByteSize: 3, Cycles: 5, Mode: AbsoluteY},
	STA_IDX: {Code: STA_IDX, Operation: sta, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	STA_IDY: {Code: STA_IDY, Operation: sta, ByteSize: 2, Cycles: 6, Mode: IndirectY},

	STX_ZER: {Code: STX_ZER, Operation: stx, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	STX_ZRY: {Code: STX_ZRY, Operation: stx, ByteSize: 2, Cycles: 4, Mode: ZeroPageY},
	STX_ABS: {Code: STX_ABS, Operation: stx, ByteSize: 3, Cycles: 4, Mode: Absolute},

	STY_ZER: {Code: STY_ZER, Operation: sty, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	STY_ZRX: {Code: STY_ZRX, Operation: sty, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	STY_ABS: {Code: STY_ABS, Operation: sty, ByteSize: 3, Cycles: 4, Mode: Absolute},

	// Register Transfers
	TAX_IMP: {Code: TAX_IMP, Operation: tax, ByteSize: 1, Cycles: 2, Mode: Implied},
	TAY_IMP: {Code: TAY_IMP, Operation: tay, ByteSize: 1, Cycles: 2, Mode: Implied},
	TXA_IMP: {Code: TXA_IMP, Operation: txa, ByteSize: 1, Cycles: 2, Mode: Implied},
	TYA_IMP: {Code: TYA_IMP, Operation: tya, ByteSize: 1, Cycles: 2, Mode: Implied},

	// Increments
	INC_ZER: {Code: INC_ZER, Operation: inc, ByteSize: 2, Cycles: 5, Mode: ZeroPage},
	INC_ZRX: {Code: INC_ZRX, Operation: inc, ByteSize: 2, Cycles: 6, Mode: ZeroPageX},
	INC_ABS: {Code: INC_ABS, Operation: inc, ByteSize: 3, Cycles: 6, Mode: Absolute},
	INC_ABX: {Code: INC_ABX, Operation: inc, ByteSize: 3, Cycles: 7, Mode: AbsoluteX},
	INX_IMP: {Code: INX_IMP, Operation: inx, ByteSize: 1, Cycles: 2, Mode: Implied},
	INY_IMP: {Code: INY_IMP, Operation: iny, ByteSize: 1, Cycles: 2, Mode: Implied},

	// Decrements
	DEC_ZER: {Code: DEC_ZER, Operation: dec, ByteSize: 2, Cycles: 5, Mode: ZeroPage},
	DEC_ZRX: {Code: DEC_ZRX, Operation: dec, ByteSize: 2, Cycles: 6, Mode: ZeroPageX},
	DEC_ABS: {Code: DEC_ABS, Operation: dec, ByteSize: 3, Cycles: 6, Mode: Absolute},
	DEC_ABX: {Code: DEC_ABX, Operation: dec, ByteSize: 3, Cycles: 7, Mode: AbsoluteX},
	DEX_IMP: {Code: DEX_IMP, Operation: dex, ByteSize: 1, Cycles: 2, Mode: Implied},
	DEY_IMP: {Code: DEY_IMP, Operation: dey, ByteSize: 1, Cycles: 2, Mode: Implied},

	BRK_IMP: {Code: BRK_IMP, Operation: brk, ByteSize: 1, Cycles: 7, Mode: Implied},
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
