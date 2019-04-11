package main
import "fmt"
import "doracpphp/instruction"

type Registers struct {
	//Register
	A,B,C,D,E,F,H,L uint8
	//Register pair
	AF,BC,DE,HL uint16

	M, T uint8
	//program counter and stack pointer
	SP,PC uint16	
}
//AF
func (reg *Registers) setAF(value uint16){
	reg.A = uint8((value & 0xFF00) >> 8)
	reg.F = uint8((value & 0x00F0))
	reg.AF = (uint16(reg.A) << 8) | uint16(reg.F)
}
func (reg *Registers)carry()uint8{
	return (reg.F >> 4) & 0x1
}
func (reg *Registers) setCarry(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 4) * -1)) | (v << 4)
}
func (reg *Registers)halfCarry()uint8{
	return (reg.F >> 5) & 0x1
}

func (reg *Registers) setHalfCarry(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 5) * -1)) | (v << 5)
}
func (reg *Registers)subtract()uint8{
	return (reg.F >> 6) & 0x1
}
func (reg *Registers) setSubtract(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 6) * -1)) | (v << 6)
}
func (reg *Registers)zero()uint8{
	return (reg.F >> 7) & 0x1
}
func (reg *Registers) setZero(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 7) * -1)) | (v << 7)
}

type Instruction func(*Emulator)
var instructionSet = map[uint8]Instruction{
	0x00:NOP,

	0x80:ADDr_b,
	0x81:ADDr_c,
	0x82:ADDr_d,
	0x83:ADDr_e,
	0x84:ADDr_h,
	0x85:ADDr_l,
	//0x86
	0x87:ADDr_a,
	0x88:ADCr_b,
	0x89:ADCr_c,
	0x8A:ADCr_d,
	0x8B:ADCr_e,
	0x8C:ADCr_h,
	0x8D:ADCr_l,
	//0x8E
	0x8F:ADCr_a,
}

type Emulator struct{
	registers Registers
	memory [0xFFFF]uint8
}
func NewEmulator() *Emulator{
	emu := new(Emulator)
	return emu
}
func (emu *Emulator)memoryRead(address uint16)uint8 {
	return emu.memory[address]
}
func (emu *Emulator)memoryWrite(address uint16, value uint8){
	if address >= 0x8000 {
		emu.memory[address] = value
	}
}

func (emu *Emulator)Debug(){
	//fmt.Printf("A : %08b  D : %08b\nB : %08b  E : %08b \nC : %08b \n",emu.registers.A,emu.registers.D,emu.registers.B,emu.registers.E,emu.registers.C)
	fmt.Printf("A : %08b  B : %08b  C : %08b\nD : %08b  E : %08b  H : %08b  L : %08b\n",
				emu.registers.A,emu.registers.B,emu.registers.C,
				emu.registers.D,emu.registers.E,emu.registers.H,emu.registers.L)
	fmt.Printf("SP : %04X  PC : %04X  Flag : %08b\n",emu.registers.SP,emu.registers.PC,emu.registers.F)

}


//instruction 0x00
func NOP(emu *Emulator){
	emu.registers.M = 1;
	emu.registers.T = 4;
}

func POPBC(emu *Emulator){
	emu.registers.B = emu.memory[emu.registers.SP]
	emu.registers.C = emu.memory[emu.registers.SP+1]
	emu.registers.SP+=2
	emu.registers.M = 3;
	emu.registers.T = 12;
}

func PUSHBC(emu *Emulator){
	emu.registers.SP--
	emu.memory[emu.registers.SP] = emu.registers.B
	emu.registers.SP--
	emu.memory[emu.registers.SP] = emu.registers.C
	emu.registers.M = 3;
	emu.registers.T = 12;
}

func POPHL(emu *Emulator){
	emu.registers.L = emu.memory[emu.registers.SP]
	emu.registers.H = emu.memory[emu.registers.SP+1]
	emu.registers.SP+=2
	emu.registers.M = 3;
	emu.registers.T = 12;
}

//0x80
func ADDr_b(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.B) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.B
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADDr_c(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.C) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.C
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADDr_d(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.D) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.D
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADDr_e(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.E) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.E
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADDr_h(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.H) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.H
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADDr_l(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.L) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.L
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADDr_a(emu *Emulator){
	emu.registers.F = 0
	if int(emu.registers.A)+int(emu.registers.A) > 255{
		emu.registers.F |= 0x10
	}
	emu.registers.A += emu.registers.A
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_b(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.B) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_c(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.C) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_d(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.D) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_e(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.E) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_h(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.H) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_l(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.L) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
func ADCr_a(emu *Emulator){
	emu.registers.F = 0
	var tmp uint16 = uint16(emu.registers.A) + uint16(emu.registers.A) + (uint16(emu.registers.F & 0x10) >> 4)
	if tmp >255{
		emu.registers.F |= 0x10
	}
	emu.registers.A = uint8(tmp & 0xFF)
	if (emu.registers.A & 255) == 0{
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
//0xB8
func CPr_b(emu *Emulator){
	if int(emu.registers.A) - int(emu.registers.B) < 0{
		emu.registers.F |= 0x10
	}
	tmp := emu.registers.A
	tmp -= emu.registers.B
	emu.registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
//0xB9
func CPr_c(emu *Emulator){
	if int(emu.registers.A) - int(emu.registers.C) < 0{
		emu.registers.F |= 0x10
	}
	tmp := emu.registers.A
	tmp -= emu.registers.C
	emu.registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
//0xBA
func CPr_d(emu *Emulator){
	if int(emu.registers.A) - int(emu.registers.D) < 0{
		emu.registers.F |= 0x10
	}
	tmp := emu.registers.A
	tmp -= emu.registers.D
	emu.registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
//0xBB
func CPr_e(emu *Emulator){
	if int(emu.registers.A) - int(emu.registers.E) < 0{
		emu.registers.F |= 0x10
	}
	tmp := emu.registers.A
	tmp -= emu.registers.E
	emu.registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
//0xBC
func CPr_h(emu *Emulator){
	if int(emu.registers.A) - int(emu.registers.H) < 0{
		emu.registers.F |= 0x10
	}
	tmp := emu.registers.A
	tmp -= emu.registers.H
	emu.registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}
//0xBD
func CPr_l(emu *Emulator){
	if int(emu.registers.A) - int(emu.registers.L) < 0{
		emu.registers.F |= 0x10
	}
	tmp := emu.registers.A
	tmp -= emu.registers.L
	emu.registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.registers.F |= 0x80
	}
	emu.registers.M = 1;
	emu.registers.T = 4;
}



func main(){
	cpu := NewEmulator()
	cpu.registers.setCarry(1)
	cpu.registers.SP=0xFFFE
	cpu.registers.B=0x2
	cpu.registers.A=0x1
	cpu.Debug()
	ADDr_b(cpu);
	cpu.Debug()
	instructionSet[0x00](cpu)

}