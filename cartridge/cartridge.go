package cartridge
import "fmt"
import "os"
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
func (cart *Cartridge)GetName()string{
	if cart.Title == ""{
		for i := uint16(0x134);i < 0x142;i++{
			ch := cart.Read(i)
			if ch != 0x00{
				cart.Title += string(ch)
			}
		}
	}
	return cart.Title
}
func NewCartridge(rom []uint8, filename string)*Cartridge{
	cart := &Cartridge{FileName:filename}
	if rom[0x147] == 0x00{
		fmt.Printf("ROM ONLY\n")
	}
	cart.CartridgeInterface = NewROM(rom)
	fmt.Println("ROM size ",rom[0x148])
	return cart
}
func (cart *Cartridge)NewCartridgeFile(filename string)(*Cartridge,error){
	rom ,err := loadROMFile(filename)
	if err != nil{
		return nil,err
	}
	return NewCartridge(rom,filename),nil
}
func loadROMFile(filename string)([]uint8,error){
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("err : ",err)
		//return nil,err
	}
	buf := make([]uint8, 0x1)
	var data []uint8
	for{
		n,err := file.Read(buf)
		if n==0{
			break;
		}
		if err != nil{
			fmt.Println("error : ",err)
			return nil,err
		}
		data = append(data,buf[0])
	}
	return data,nil
}

