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
	cpu.Registers.H=0xFF
	cpu.Registers.L=0x00
	cpu.Memory[0xFF00] = 0x86
	cpu.Debug()
	fmt.Printf("%08b\n",cpu.Memory[0xFF00])
	cpu.Prefix[0x06](cpu)
	cpu.Debug()
	fmt.Printf("%08b\n",cpu.Memory[0xFF00])

}
