package go6502

import (
	"testing"

	"github.com/zehlt/go6502/asrt"
)

func initSampleMemory() Mem {
	mem := Mem{
		0x10, 0x00, 0x80, 0xE0,
	}
	return mem
}

func TestReadSingleByte(t *testing.T) {
	mem := initSampleMemory()

	var addr uint16 = 0x02
	b := mem.ReadByte(addr)

	asrt.Equal(t, mem[addr], b)
}

func TestWriteSingleByte(t *testing.T) {
	mem := initSampleMemory()

	var addr uint16 = 0xCA
	var data uint8 = 0xEF

	mem.WriteByte(addr, data)

	asrt.Equal(t, mem[addr], data)
}

func TestMultipleWriteAndRead(t *testing.T) {
	mem := initSampleMemory()

	var addr uint16 = 0xCA
	var data uint8 = 0xEF

	mem.WriteByte(addr, data)
	mem.WriteByte(addr+2, data+2)
	mem.WriteByte(addr+10, data+10)

	got1 := mem.ReadByte(addr)
	got2 := mem.ReadByte(addr + 2)
	got3 := mem.ReadByte(addr + 10)

	asrt.Equal(t, mem[addr], got1)
	asrt.Equal(t, mem[addr+2], got2)
	asrt.Equal(t, mem[addr+10], got3)
}

// TODO: maybe something special should happen at 0xFFFF overflow ?
func TestReadSingleWord(t *testing.T) {
	mem := initSampleMemory()
	var addr uint16 = 0x01

	var lo uint16 = uint16(mem[addr])
	var hi uint16 = uint16(mem[addr+1])

	var want uint16 = ((hi << 8) | lo)
	got := mem.ReadWord(addr)

	asrt.Equal(t, want, got)
}

func TestWriteSingleWord(t *testing.T) {
	mem := initSampleMemory()
	var addr uint16 = 0xDE01
	var word uint16 = 0x3080

	mem.WriteWord(addr, word)

	asrt.Equal(t, mem[addr], uint8(0x80))
	asrt.Equal(t, mem[addr+1], uint8(0x30))
}

func TestWriteTwoWord(t *testing.T) {
	mem := initSampleMemory()
	var addr uint16 = 0xDE01
	var word1 uint16 = 0x3080
	var word2 uint16 = 0xEADF

	mem.WriteWord(addr, word1)
	mem.WriteWord(addr+1, word2)

	asrt.Equal(t, mem[addr], uint8(0x80))
	asrt.Equal(t, mem[addr+1], uint8(0xDF))
	asrt.Equal(t, mem[addr+2], uint8(0xEA))
}

// TODO: Probably test for overloading ?
func TestWriteBytes(t *testing.T) {
	mem := initSampleMemory()

	var addr uint16 = 0xA3E4

	want := []uint8{
		0x01, 0x58, 0xCE, 0x9D,
	}

	mem.WriteBytes(addr, want)

	for i := 0; i < len(want); i++ {
		asrt.Equal(t, mem[addr+(uint16(i))], want[i])
	}
}
