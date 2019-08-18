package gameboy
import "github.com/doracpphp/GameBoyEmulator_Golang/cartridge"

type MMU struct{
	Bios [256]uint8
	Cartridge *cartridge.Cartridge
	VRAM [0x4000]uint8
	WRAM [0x9000]uint8
	OAM [0x100] uint8
	HRAM[0x100]uint8
}
func NewMMU() * MMU{
	mmu := new(MMU)
	return mmu
}
func (mmu *MMU)Read(addr uint16)uint8{
	switch {
	case addr < 0x8000:
		return mmu.Cartridge.Read(addr)
	case addr < 0xA000:
		return mmu.VRAM[addr - 0x8000]
	case addr < 0xC000:
		return mmu.Cartridge.Read(addr)
	//メインメモリ
	case addr < 0xD000:
		return mmu.WRAM[addr - 0xC000]
	case addr < 0xE000:
		return mmu.WRAM[addr - 0xD000]
	case addr < 0xFE00:
		return 0xFF
	//スプライト属性テーブル
	case addr < 0xFEA0:
		return mmu.OAM[addr - 0xFE00]
	case addr < 0xFF00:
		return 0xFF
	default:
		return mmu.ReadHRAM(addr)
	}
	return 0x00
}
func (mmu *MMU)ReadHRAM(addr uint16)uint8{
	return mmu.HRAM[addr - 0xFF00]
}
func (mmu *MMU)Write(addr uint16,value uint8){
	switch{
	case addr < 0x8000:
		mmu.Cartridge.WriteROM(addr, value)
	case addr < 0xA000:
		mmu.VRAM[addr - 0x8000] = value
	case addr < 0xC000:
		mmu.Cartridge.WriteROM(addr, value)
	case addr < 0xD000:
		mmu.WRAM[addr - 0xC000] = value
	case addr < 0xE000:
		mmu.WRAM[addr - 0xD000] = value
	case addr < 0xFE00:
	case addr < 0xFEA0:
		mmu.OAM[addr - 0xFE00] = value
	case addr < 0xFF00:
		break
	default:
		mmu.WriteHRAM(addr, value)
	}
}
func (mmu *MMU)WriteHRAM(addr uint16,value uint8){
	switch {
	case addr == 0xFF41:
		mmu.HRAM[0x41] = value
	default:
		mmu.HRAM[addr - 0xFF00] = value
	}
}
