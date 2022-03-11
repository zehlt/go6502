package cpu

type Nes struct {
	cpu Cpu
	mem Memory
}

func (n *Nes) reset() {
	prgRom := []uint8{
		LDA_IMM, 0x41, BRK_IMP, TAX_IMP, BRK_IMP,
	}

	const prgStart = 0x8000
	n.mem.writeBytes(prgStart, prgRom)
	n.mem.writeWord(0xFFFC, prgStart)

	n.cpu.Reset(&n.mem)
}

func (n *Nes) Start() {
	n.reset()

	n.cpu.Run(&n.mem)
}
