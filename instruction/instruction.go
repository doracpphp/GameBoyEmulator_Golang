package instruction
import "github.com/doracpphp/GameBoyEmulator_Golang/register"
import "fmt"
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
	Registers register.Register
	Memory [0xFFFF]uint8
	Inst map[uint8]Instruction
}
func NewEmulator() *Emulator{
	emu := new(Emulator)
	emu.Inst = instructionSet
	return emu
}
func (emu *Emulator)memoryRead(address uint16)uint8 {
	return emu.Memory[address]
}
func (emu *Emulator)memoryWrite(address uint16, value uint8){
	if address >= 0x8000 {
		emu.Memory[address] = value
	}
}

func (emu *Emulator)Debug(){
	//fmt.Printf("A : %08b  D : %08b\nB : %08b  E : %08b \nC : %08b \n",emu.Registers.A,emu.Registers.D,emu.Registers.B,emu.Registers.E,emu.Registers.C)
	fmt.Printf("A : %08b  B : %08b  C : %08b\nD : %08b  E : %08b  H : %08b  L : %08b\n",
				emu.Registers.A,emu.Registers.B,emu.Registers.C,
				emu.Registers.D,emu.Registers.E,emu.Registers.H,emu.Registers.L)
	fmt.Printf("SP : %04X  PC : %04X  Flag : %08b\n",emu.Registers.SP,emu.Registers.PC,emu.Registers.F)

}


//instruction 0x00
func NOP(emu *Emulator){
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}

func POPBC(emu *Emulator){
	emu.Registers.B = emu.Memory[emu.Registers.SP]
	emu.Registers.C = emu.Memory[emu.Registers.SP+1]
	emu.Registers.SP+=2
	emu.Registers.M = 3;
	emu.Registers.T = 12;
}

func PUSHBC(emu *Emulator){
	emu.Registers.SP--
	emu.Memory[emu.Registers.SP] = emu.Registers.B
	emu.Registers.SP--
	emu.Memory[emu.Registers.SP] = emu.Registers.C
	emu.Registers.M = 3;
	emu.Registers.T = 12;
}

func POPHL(emu *Emulator){
	emu.Registers.L = emu.Memory[emu.Registers.SP]
	emu.Registers.H = emu.Memory[emu.Registers.SP+1]
	emu.Registers.SP+=2
	emu.Registers.M = 3;
	emu.Registers.T = 12;
}

//0x80
func ADDr_b(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.B) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.B
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADDr_c(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.C) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.C
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADDr_d(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.D) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.D
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADDr_e(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.E) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.E
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADDr_h(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.H) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.H
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADDr_l(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.L) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.L
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADDr_a(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.A) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.A
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_b(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.B) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_c(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.C) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_d(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.D) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_e(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.E) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_h(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.H) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_l(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.L) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
func ADCr_a(emu *Emulator){
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.A) + (uint16(emu.Registers.F & 0x10) >> 4)
	if tmp >255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
//0xB8
func CPr_b(emu *Emulator){
	if int(emu.Registers.A) - int(emu.Registers.B) < 0{
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.B
	emu.Registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
//0xB9
func CPr_c(emu *Emulator){
	if int(emu.Registers.A) - int(emu.Registers.C) < 0{
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.C
	emu.Registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
//0xBA
func CPr_d(emu *Emulator){
	if int(emu.Registers.A) - int(emu.Registers.D) < 0{
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.D
	emu.Registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
//0xBB
func CPr_e(emu *Emulator){
	if int(emu.Registers.A) - int(emu.Registers.E) < 0{
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.E
	emu.Registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
//0xBC
func CPr_h(emu *Emulator){
	if int(emu.Registers.A) - int(emu.Registers.H) < 0{
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.H
	emu.Registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
//0xBD
func CPr_l(emu *Emulator){
	if int(emu.Registers.A) - int(emu.Registers.L) < 0{
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.L
	emu.Registers.F |= 0x40
	if tmp & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1;
	emu.Registers.T = 4;
}
