package cartridge


type MBC1 struct{
	rom []uint8
	romBank uint32
	ram []uint8

}
func NewMBC1(data []uint8) CartridgeInterface{
	return &MBC1{
		rom: data,
		romBank: 1,
	}
}
func (mbc *MBC1)Read(addr uint16)uint8{
	return mbc.rom[addr]
}
func (mbc *MBC1)WriteROM(addr uint16, value uint8){
	mbc.rom[addr]=value
}
func (mbc *MBC1)WriteRAM(addr uint16, value uint8){
	mbc.rom[addr]=value
}