package main

import "github.com/doracpphp/GameBoyEmulator_Golang/instruction"
import "fmt"
func main() {
	cpu := instruction.NewEmulator()
	cpu.Registers.SetCarry(1)
	cpu.Registers.SP = 0xFFFE
	cpu.Registers.B = 0xFF
	cpu.Registers.A = 0x00
	cpu.Registers.L = 0x00
	cpu.Registers.H = 0x80
	cpu.Debug()
	cpu.Inst[0x70](cpu)
	cpu.Debug()
	fmt.Println(cpu.Memory[0x8000])
}
