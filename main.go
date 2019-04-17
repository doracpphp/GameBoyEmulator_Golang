package main

import "github.com/doracpphp/GameBoyEmulator_Golang/instruction"
import "fmt"
func main() {
	cpu := instruction.NewEmulator()
	cpu.Debug()
	for {
		opcode := cpu.Memory[cpu.Registers.PC]
		cpu.Registers.PC+=1
		fmt.Printf("Opcode : %04X\n",opcode)
		cpu.Inst[opcode](cpu)
		cpu.Debug()
	}
	
}
