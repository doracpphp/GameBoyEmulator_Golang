package gameboy

type Instruction func(*Emulator)
type Emulator struct {
	Registers Register
	Inst      map[uint8]Instruction
	Prefix    map[uint8]Instruction
	Halt,Ime      uint8
	ClockSpeed uint64
	Memory	MMU
}
func NewGameboy() *Emulator {
	emu := new(Emulator)
	emu.Inst = instructionSet
	emu.Prefix = prefixset
	emu.ClockSpeed = 4194304
	bios := [...]uint8{
		0x31, 0xFE, 0xFF, 0xAF, 0x21, 0xFF, 0x9F, 0x32, 0xCB, 0x7C, 0x20, 0xFB, 0x21, 0x26, 0xFF, 0x0E,
	}
	for i, v := range bios {
		emu.Memory.Bios[i]=v
	}
	return emu
}