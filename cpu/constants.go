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
	Accumulator
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

	TSX_IMP = 0xBA
	TXS_IMP = 0x9A
	PHA_IMP = 0x48
	PHP_IMP = 0x08
	PLA_IMP = 0x68
	PLP_IMP = 0x28

	AND_IMM = 0x29
	AND_ZER = 0x25
	AND_ZRX = 0x35
	AND_ABS = 0x2D
	AND_ABX = 0x3D
	AND_ABY = 0x39
	AND_IDX = 0x21
	AND_IDY = 0x31

	EOR_IMM = 0x49
	EOR_ZER = 0x45
	EOR_ZRX = 0x55
	EOR_ABS = 0x4D
	EOR_ABX = 0x5D
	EOR_ABY = 0x59
	EOR_IDX = 0x41
	EOR_IDY = 0x51

	ORA_IMM = 0x09
	ORA_ZER = 0x05
	ORA_ZRX = 0x15
	ORA_ABS = 0x0D
	ORA_ABX = 0x1D
	ORA_ABY = 0x19
	ORA_IDX = 0x01
	ORA_IDY = 0x11

	BIT_ZER = 0x24
	BIT_ABS = 0x2C

	ADC_IMM = 0x69
	ADC_ZER = 0x65
	ADC_ZRX = 0x75
	ADC_ABS = 0x6D
	ADC_ABX = 0x7D
	ADC_ABY = 0x79
	ADC_IDX = 0x61
	ADC_IDY = 0x71

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

	ASL_ACC = 0x0A
	ASL_ZER = 0x06
	ASL_ZRX = 0x16
	ASL_ABS = 0x0E
	ASL_ABX = 0x1E

	LSR_ACC = 0x4A
	LSR_ZER = 0x46
	LSR_ZRX = 0x56
	LSR_ABS = 0x4E
	LSR_ABX = 0x5E

	ROL_ACC = 0x2A
	ROL_ZER = 0x26
	ROL_ZRX = 0x36
	ROL_ABS = 0x2E
	ROL_ABX = 0x3E

	ROR_ACC = 0x6A
	ROR_ZER = 0x66
	ROR_ZRX = 0x76
	ROR_ABS = 0x6E
	ROR_ABX = 0x7E

	CLC_IMP = 0x18
	CLD_IMP = 0xD8
	CLI_IMP = 0x58
	CLV_IMP = 0xB8

	SEC_IMP = 0x38
	SED_IMP = 0xF8
	SEI_IMP = 0x78

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

	// Stack
	TSX_IMP: {Code: TSX_IMP, Operation: tsx, ByteSize: 1, Cycles: 2, Mode: Implied},
	TXS_IMP: {Code: TXS_IMP, Operation: txs, ByteSize: 1, Cycles: 2, Mode: Implied},
	PHA_IMP: {Code: PHA_IMP, Operation: pha, ByteSize: 1, Cycles: 3, Mode: Implied},
	PHP_IMP: {Code: PHP_IMP, Operation: php, ByteSize: 1, Cycles: 3, Mode: Implied},
	PLA_IMP: {Code: PLA_IMP, Operation: pla, ByteSize: 1, Cycles: 4, Mode: Implied},
	PLP_IMP: {Code: PLP_IMP, Operation: plp, ByteSize: 1, Cycles: 4, Mode: Implied},

	// Logical
	AND_IMM: {Code: AND_IMM, Operation: and, ByteSize: 2, Cycles: 2, Mode: Immediate},
	AND_ZER: {Code: AND_ZER, Operation: and, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	AND_ZRX: {Code: AND_ZRX, Operation: and, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	AND_ABS: {Code: AND_ABS, Operation: and, ByteSize: 3, Cycles: 4, Mode: Absolute},
	AND_ABX: {Code: AND_ABX, Operation: and, ByteSize: 3, Cycles: 4, Mode: AbsoluteX1},
	AND_ABY: {Code: AND_ABY, Operation: and, ByteSize: 3, Cycles: 4, Mode: AbsoluteY1},
	AND_IDX: {Code: AND_IDX, Operation: and, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	AND_IDY: {Code: AND_IDY, Operation: and, ByteSize: 2, Cycles: 5, Mode: IndirectY1},

	EOR_IMM: {Code: EOR_IMM, Operation: eor, ByteSize: 2, Cycles: 2, Mode: Immediate},
	EOR_ZER: {Code: EOR_ZER, Operation: eor, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	EOR_ZRX: {Code: EOR_ZRX, Operation: eor, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	EOR_ABS: {Code: EOR_ABS, Operation: eor, ByteSize: 3, Cycles: 4, Mode: Absolute},
	EOR_ABX: {Code: EOR_ABX, Operation: eor, ByteSize: 3, Cycles: 4, Mode: AbsoluteX1},
	EOR_ABY: {Code: EOR_ABY, Operation: eor, ByteSize: 3, Cycles: 4, Mode: AbsoluteY1},
	EOR_IDX: {Code: EOR_IDX, Operation: eor, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	EOR_IDY: {Code: EOR_IDY, Operation: eor, ByteSize: 2, Cycles: 5, Mode: IndirectY1},

	ORA_IMM: {Code: ORA_IMM, Operation: aor, ByteSize: 2, Cycles: 2, Mode: Immediate},
	ORA_ZER: {Code: ORA_ZER, Operation: aor, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	ORA_ZRX: {Code: ORA_ZRX, Operation: aor, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	ORA_ABS: {Code: ORA_ABS, Operation: aor, ByteSize: 3, Cycles: 4, Mode: Absolute},
	ORA_ABX: {Code: ORA_ABX, Operation: aor, ByteSize: 3, Cycles: 4, Mode: AbsoluteX1},
	ORA_ABY: {Code: ORA_ABY, Operation: aor, ByteSize: 3, Cycles: 4, Mode: AbsoluteY1},
	ORA_IDX: {Code: ORA_IDX, Operation: aor, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	ORA_IDY: {Code: ORA_IDY, Operation: aor, ByteSize: 2, Cycles: 5, Mode: IndirectY1},

	BIT_ZER: {Code: BIT_ZER, Operation: bit, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	BIT_ABS: {Code: BIT_ABS, Operation: bit, ByteSize: 3, Cycles: 4, Mode: Absolute},

	// Arithmetic
	ADC_IMM: {Code: ADC_IMM, Operation: adc, ByteSize: 2, Cycles: 2, Mode: Immediate},
	ADC_ZER: {Code: ADC_ZER, Operation: adc, ByteSize: 2, Cycles: 3, Mode: ZeroPage},
	ADC_ZRX: {Code: ADC_ZRX, Operation: adc, ByteSize: 2, Cycles: 4, Mode: ZeroPageX},
	ADC_ABS: {Code: ADC_ABS, Operation: adc, ByteSize: 3, Cycles: 4, Mode: Absolute},
	ADC_ABX: {Code: ADC_ABX, Operation: adc, ByteSize: 3, Cycles: 4, Mode: AbsoluteX1},
	ADC_ABY: {Code: ADC_ABY, Operation: adc, ByteSize: 3, Cycles: 4, Mode: AbsoluteY1},
	ADC_IDX: {Code: ADC_IDX, Operation: adc, ByteSize: 2, Cycles: 6, Mode: IndirectX},
	ADC_IDY: {Code: ADC_IDY, Operation: adc, ByteSize: 2, Cycles: 5, Mode: IndirectY1},

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

	// Shifts
	ASL_ACC: {Code: ASL_ACC, Operation: asl, ByteSize: 1, Cycles: 2, Mode: Accumulator},
	ASL_ZER: {Code: ASL_ZER, Operation: asl, ByteSize: 2, Cycles: 5, Mode: ZeroPage},
	ASL_ZRX: {Code: ASL_ZRX, Operation: asl, ByteSize: 2, Cycles: 6, Mode: ZeroPageX},
	ASL_ABS: {Code: ASL_ABS, Operation: asl, ByteSize: 3, Cycles: 6, Mode: Absolute},
	ASL_ABX: {Code: ASL_ABX, Operation: asl, ByteSize: 3, Cycles: 7, Mode: AbsoluteX},

	LSR_ACC: {Code: LSR_ACC, Operation: lsr, ByteSize: 1, Cycles: 2, Mode: Accumulator},
	LSR_ZER: {Code: LSR_ZER, Operation: lsr, ByteSize: 2, Cycles: 5, Mode: ZeroPage},
	LSR_ZRX: {Code: LSR_ZRX, Operation: lsr, ByteSize: 2, Cycles: 6, Mode: ZeroPageX},
	LSR_ABS: {Code: LSR_ABS, Operation: lsr, ByteSize: 3, Cycles: 6, Mode: Absolute},
	LSR_ABX: {Code: LSR_ABX, Operation: lsr, ByteSize: 3, Cycles: 7, Mode: AbsoluteX},

	ROL_ACC: {Code: ROL_ACC, Operation: rol, ByteSize: 1, Cycles: 2, Mode: Accumulator},
	ROL_ZER: {Code: ROL_ZER, Operation: rol, ByteSize: 2, Cycles: 5, Mode: ZeroPage},
	ROL_ZRX: {Code: ROL_ZRX, Operation: rol, ByteSize: 2, Cycles: 6, Mode: ZeroPageX},
	ROL_ABS: {Code: ROL_ABS, Operation: rol, ByteSize: 3, Cycles: 6, Mode: Absolute},
	ROL_ABX: {Code: ROL_ABX, Operation: rol, ByteSize: 3, Cycles: 7, Mode: AbsoluteX},

	ROR_ACC: {Code: ROR_ACC, Operation: ror, ByteSize: 1, Cycles: 2, Mode: Accumulator},
	ROR_ZER: {Code: ROR_ZER, Operation: ror, ByteSize: 2, Cycles: 5, Mode: ZeroPage},
	ROR_ZRX: {Code: ROR_ZRX, Operation: ror, ByteSize: 2, Cycles: 6, Mode: ZeroPageX},
	ROR_ABS: {Code: ROR_ABS, Operation: ror, ByteSize: 3, Cycles: 6, Mode: Absolute},
	ROR_ABX: {Code: ROR_ABX, Operation: ror, ByteSize: 3, Cycles: 7, Mode: AbsoluteX},

	// Status Flag Changes
	CLC_IMP: {Code: CLC_IMP, Operation: clc, ByteSize: 1, Cycles: 2, Mode: Implied},
	CLD_IMP: {Code: CLD_IMP, Operation: cld, ByteSize: 1, Cycles: 2, Mode: Implied},
	CLI_IMP: {Code: CLI_IMP, Operation: cli, ByteSize: 1, Cycles: 2, Mode: Implied},
	CLV_IMP: {Code: CLV_IMP, Operation: clv, ByteSize: 1, Cycles: 2, Mode: Implied},

	SEC_IMP: {Code: SEC_IMP, Operation: sec, ByteSize: 1, Cycles: 2, Mode: Implied},
	SED_IMP: {Code: SED_IMP, Operation: sed, ByteSize: 1, Cycles: 2, Mode: Implied},
	SEI_IMP: {Code: SEI_IMP, Operation: sei, ByteSize: 1, Cycles: 2, Mode: Implied},

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
