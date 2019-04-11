package main

import "github.com/doracpphp/GameBoyEmulator_Golang/instruction"


func main(){
	cpu := instruction.NewEmulator()
	cpu.Registers.SetCarry(1)
	cpu.Registers.SP=0xFFFE
	cpu.Registers.B=0x2
	cpu.Registers.A=0x1
	cpu.Debug()
	instruction.ADDr_b(cpu);
	cpu.Debug()
	cpu.Inst[0x00](cpu)

}