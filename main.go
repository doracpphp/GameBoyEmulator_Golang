package main

import "github.com/doracpphp/GameBoyEmulator_Golang/instruction"
import "fmt"
func main() {
	cpu := instruction.NewEmulator()
	cpu.Registers.SetCarry(1)
	cpu.Registers.SP = 0xFFFE
	cpu.Registers.A = 0xFF
	cpu.Registers.B = 0xFF
	cpu.Registers.F = 0xFF
	cpu.Registers.L = 0xFF
	cpu.Registers.H = 0xFF
	cpu.Memory[0x00] = 0x01
	cpu.Debug()
	cpu.Inst[0x18](cpu)
	cpu.Debug()
	fmt.Println(cpu.Memory[0x8000])
}
