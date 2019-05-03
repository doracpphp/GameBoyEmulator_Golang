package main

import "github.com/doracpphp/GameBoyEmulator_Golang/instruction"
import "fmt"
func main() {
	cpu := instruction.NewEmulator()

	/*for {
		opcode := cpu.Memory[cpu.Registers.PC]
		cpu.Registers.PC+=1
		fmt.Printf("Opcode : %04X\n",opcode)
		cpu.Inst[opcode](cpu)
		cpu.Debug()
	}
	*/
	/*cpu.Registers.A=0x5
	cpu.Registers.B=0x4
	cpu.Registers.SP=0x04
	cpu.Registers.PC=0x1
	cpu.Registers.F |= 0x10
	cpu.Debug()
	cpu.Inst[0xC3](cpu)
	cpu.Debug()
	*/
	cart,err := cpu.MMU.Cartridge.NewCartridgeFile("test.gb")
	if err != nil{
		fmt.Println("error ",err)
	}
	fmt.Println(cart.GetName())
	fmt.Printf("%08b\n",cpu.Memory[0xFF00])
}
