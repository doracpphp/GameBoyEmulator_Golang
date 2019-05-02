package cartridge
import "fmt"
type CartridgeType struct {
	ID          byte
	Description string
}
type CartridgeInterface interface{
	Read(addr uint16)uint8
	WriteROM(addr uint16, value uint8)
	WriteRAM(addr uint16, value uint8)
}
type Cartridge struct {
	CartridgeInterface
	Title string
	FileName string
}
func NewCartridge(rom []uint8, filename string)*Cartridge{
	cart := &Cartridge{FileName:filename}
	if rom[0x147] == 0x00{
		fmt.Printf("ROM ONLY\n")
	}
	return cart
}

