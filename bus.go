package go6502

type Bus interface {
	Write(addr uint16, data uint8)
	Read(addr uint16) uint8
	WriteWord(addr uint16, data uint16)
	ReadWord(addr uint16) uint16
}

type BusEx struct {
	m *Mem
}

func (b BusEx) Read(addr uint16) uint8 {
	return b.m.Read(addr)
}

func (b BusEx) Write(addr uint16, data uint8) {
	b.m.Write(addr, data)
}

func (b BusEx) ReadWord(addr uint16) uint16 {
	return b.m.ReadWord(addr)
}

func (b BusEx) WriteWord(addr uint16, data uint16) {
	b.m.WriteWord(addr, data)
}

func (b BusEx) WriteBytes(addr uint16, data []uint8) {
	b.m.WriteBytes(addr, data)
}
