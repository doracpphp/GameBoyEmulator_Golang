package cartridge
type ROM struct{
	rom []uint8
}
func NewROM(data []uint8) CartridgeInterface{
	return &ROM{
		rom: data,
	}
}
func (rom *ROM)Read(addr uint16)uint8{
	return rom.rom[addr]
}
func (rom *ROM)WriteROM(addr uint16, value uint8){

}
func (rom *ROM)WriteRAM(addr uint16, value uint8){
	
}