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
	cpu.Registers.A=0x5
	cpu.Registers.B=0x4
	cpu.Registers.F |= 0x10
	cpu.Debug()
	cpu.Inst[0x97](cpu)
	cpu.Debug()
	fmt.Printf("%08b\n",cpu.Memory[0xFF00])

}
