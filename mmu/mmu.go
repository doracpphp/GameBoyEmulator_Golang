package mmu
import "github.com/doracpphp/GameBoyEmulator_Golang/cartridge"

type MMU struct{
	bios [256]uint8
	cartridge *cartridge.Cartridge
	VRAM [0x4000]uint8
	WRAM [0x1000]uint8
	
}
func NewMMU() * MMU{
	mmu := new(MMU)
	return mmu
}
func (mmu *MMU)Read(addr uint16)uint8{
	switch {
	case addr < 0x8000:
		return mmu.cartridge.Read(addr)
	}
	return 0x00
}