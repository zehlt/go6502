package go6502

/*
	CPU MEM MAP

	0xFFFF |
	0x8000 | PRG ROM map from cartridge

	0x8000 |
	0x6000 | Save RAM map from cartridge

	0x6000 |
	0x4020 | Expansion Rom (Mappers)

	0x4020 |
	0x2000 | IO Registers PPU APU...

	0x2000 |
	0x0000 | CPU RAM
*/

type Mem [0xFFFF]uint8

func (m *Mem) ReadByte(addr uint16) uint8 {
	return m[addr]
}

func (m *Mem) WriteByte(addr uint16, data uint8) {
	m[addr] = data
}

func (m *Mem) ReadWord(addr uint16) uint16 {
	var lo uint16 = uint16(m.ReadByte(addr))
	var hi uint16 = uint16(m.ReadByte(addr + 1))
	var word uint16 = ((hi << 8) | lo)
	return word
}

func (m *Mem) WriteWord(addr uint16, data uint16) {
	hi := uint8((data >> 8))
	lo := uint8(data)
	m.WriteByte(addr, lo)
	m.WriteByte(addr+1, hi)
}

func (m *Mem) WriteBytes(addr uint16, data []uint8) {
	for index, value := range data {
		m.WriteByte(addr+uint16(index), value)
	}
}
