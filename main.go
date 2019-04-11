package main

import "github.com/doracpphp/GameBoyEmulator_Golang/instruction"

func main() {
	cpu := instruction.NewEmulator()
	cpu.Registers.SetCarry(1)
	cpu.Registers.SP = 0xFFFE
	cpu.Registers.B = 0xFF
	cpu.Registers.A = 0x00
	cpu.Registers.L = 0x01
	cpu.Registers.H = 0x00
	cpu.Debug()
	cpu.Inst[0xA8](cpu)
	cpu.Debug()
}
