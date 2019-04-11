package instruction

import "github.com/doracpphp/GameBoyEmulator_Golang/register"
import "fmt"

type Instruction func(*Emulator)

var instructionSet = map[uint8]Instruction{
	0x00: NOP,
	//0x40
	0x40: LDrr_bb, 0x41: LDrr_bc, 0x42: LDrr_bd,
	0x43: LDrr_be, 0x44: LDrr_bh, 0x45: LDrr_bl,
	0x46: LDrHL_b, 0x47: LDrr_ba, 0x48: LDrr_cb,
	0x49: LDrr_cc, 0x4A: LDrr_cd, 0x4B: LDrr_ce,
	0x4C: LDrr_ch, 0x4D: LDrr_cl, 0x4E: LDrHL_c,
	0x4F: LDrr_ca,
	//0x50
	0x50: LDrr_db, 0x51: LDrr_dc, 0x52: LDrr_dd,
	0x53: LDrr_de, 0x54: LDrr_dh, 0x55: LDrr_dl,
	0x56: LDrHL_d, 0x57: LDrr_da, 0x58: LDrr_eb,
	0x59: LDrr_ec, 0x5A: LDrr_ed, 0x5B: LDrr_ee,
	0x5C: LDrr_eh, 0x5D: LDrr_el, 0x5E: LDrHL_e,
	0x5F: LDrr_ea,
	//0x80
	0x80: ADDr_b, 0x81: ADDr_c, 0x82: ADDr_d,
	0x83: ADDr_e, 0x84: ADDr_h, 0x85: ADDr_l,
	0x86: ADDr_hl,0x87: ADDr_a, 0x88: ADCr_b,
	0x89: ADCr_c, 0x8A: ADCr_d, 0x8B: ADCr_e, 
	0x8C: ADCr_h, 0x8D: ADCr_l,
	//0x8E
	0x8F: ADCr_a,
	//0xA0
	0xA0: ANDr_b, 0xA1: ANDr_c, 0xA2: ANDr_d,
	0xA3: ANDr_e, 0xA4: ANDr_h, 0xA5: ANDr_l,
	0xA6: ANDHL, 0xA7: ANDr_a, 0xA8: XORr_b,
	0xA9: XORr_c, 0xAA: XORr_d, 0xAB: XORr_e,
	0xAC: XORr_h, 0xAD: XORr_l, 0xAE: XORr_hl,
	0xAF: XORr_a,
	//0xB0
	0xB0: ORr_b, 0xB1: ORr_c, 0xB2: ORr_d,
	0xB3: ORr_e, 0xB4: ORr_h, 0xB5: ORr_l,
	0xB6: ORHL, 0xB7: ORr_a, 0xB8: CPr_b,
	0xB9: CPr_c, 0xBA: CPr_d, 0xBB: CPr_e,
	0xBC: CPr_h, 0xBD: CPr_l, 0xBE: CPHL,
	0xBF: CPr_a,
}

type Emulator struct {
	Registers register.Register
	Memory    [0xFFFF]uint8
	Inst      map[uint8]Instruction
}

func NewEmulator() *Emulator {
	emu := new(Emulator)
	emu.Inst = instructionSet
	bios := [...]uint8{
		0x31, 0xFE, 0xFF, 0xAF, 0x21, 0xFF, 0x9F, 0x32, 0xCB, 0x7C, 0x20, 0xFB, 0x21, 0x26, 0xFF, 0x0E,
	}
	for i, v := range bios {
		emu.Memory[i] = v
	}
	return emu
}
func (emu *Emulator) MemoryRead(address uint16) uint8 {
	return emu.Memory[address]
}
func (emu *Emulator) memoryWrite(address uint16, value uint8) {
	if address >= 0x8000 {
		emu.Memory[address] = value
	}
}

func (emu *Emulator) Debug() {
	//fmt.Printf("A : %08b  D : %08b\nB : %08b  E : %08b \nC : %08b \n",emu.Registers.A,emu.Registers.D,emu.Registers.B,emu.Registers.E,emu.Registers.C)
	fmt.Printf("------------------------------------\nA : %08b  B : %08b  C : %08b\nD : %08b  E : %08b  H : %08b  L : %08b\n",
		emu.Registers.A, emu.Registers.B, emu.Registers.C,
		emu.Registers.D, emu.Registers.E, emu.Registers.H, emu.Registers.L)
	fmt.Printf("SP : %04X  PC : %04X  Flag : %08b\n------------------------------------\n", emu.Registers.SP, emu.Registers.PC, emu.Registers.F)

}

//instruction 0x00
func NOP(emu *Emulator) {
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_bb(emu *Emulator) {
	emu.Registers.B = emu.Registers.B
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_bc(emu *Emulator) {
	emu.Registers.B = emu.Registers.C
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_bd(emu *Emulator) {
	emu.Registers.B = emu.Registers.D
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_be(emu *Emulator) {
	emu.Registers.B = emu.Registers.E
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_bh(emu *Emulator) {
	emu.Registers.B = emu.Registers.H
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_bl(emu *Emulator) {
	emu.Registers.B = emu.Registers.L
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrHL_b(emu *Emulator) {
	emu.Registers.B = emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func LDrr_ba(emu *Emulator) {
	emu.Registers.B = emu.Registers.A
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_cb(emu *Emulator) {
	emu.Registers.C = emu.Registers.B
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_cc(emu *Emulator) {
	emu.Registers.C = emu.Registers.C
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_cd(emu *Emulator) {
	emu.Registers.C = emu.Registers.D
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_ce(emu *Emulator) {
	emu.Registers.C = emu.Registers.E
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_ch(emu *Emulator) {
	emu.Registers.C = emu.Registers.H
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_cl(emu *Emulator) {
	emu.Registers.C = emu.Registers.L
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrHL_c(emu *Emulator) {
	emu.Registers.C = emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func LDrr_ca(emu *Emulator) {
	emu.Registers.C = emu.Registers.A
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_db(emu *Emulator) {
	emu.Registers.D = emu.Registers.B
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_dc(emu *Emulator) {
	emu.Registers.D = emu.Registers.C
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_dd(emu *Emulator) {
	emu.Registers.D = emu.Registers.D
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_de(emu *Emulator) {
	emu.Registers.D = emu.Registers.E
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_dh(emu *Emulator) {
	emu.Registers.D = emu.Registers.H
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_dl(emu *Emulator) {
	emu.Registers.D = emu.Registers.L
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrHL_d(emu *Emulator) {
	emu.Registers.D = emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func LDrr_da(emu *Emulator) {
	emu.Registers.D = emu.Registers.A
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_eb(emu *Emulator) {
	emu.Registers.E = emu.Registers.B
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_ec(emu *Emulator) {
	emu.Registers.E = emu.Registers.C
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_ed(emu *Emulator) {
	emu.Registers.E = emu.Registers.D
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_ee(emu *Emulator) {
	emu.Registers.E = emu.Registers.E
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_eh(emu *Emulator) {
	emu.Registers.E = emu.Registers.H
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrr_el(emu *Emulator){
	emu.Registers.E = emu.Registers.L
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func LDrHL_e(emu *Emulator) {
	emu.Registers.E = emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func LDrr_ea(emu *Emulator){
	emu.Registers.E = emu.Registers.A
	emu.Registers.M = 1
	emu.Registers.T = 4
}


func POPBC(emu *Emulator) {
	emu.Registers.B = emu.Memory[emu.Registers.SP]
	emu.Registers.C = emu.Memory[emu.Registers.SP+1]
	emu.Registers.SP += 2
	emu.Registers.M = 3
	emu.Registers.T = 12
}

func PUSHBC(emu *Emulator) {
	emu.Registers.SP--
	emu.Memory[emu.Registers.SP] = emu.Registers.B
	emu.Registers.SP--
	emu.Memory[emu.Registers.SP] = emu.Registers.C
	emu.Registers.M = 3
	emu.Registers.T = 12
}

func POPHL(emu *Emulator) {
	emu.Registers.L = emu.Memory[emu.Registers.SP]
	emu.Registers.H = emu.Memory[emu.Registers.SP+1]
	emu.Registers.SP += 2
	emu.Registers.M = 3
	emu.Registers.T = 12
}

//0x80
func ADDr_b(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.B) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.B
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADDr_c(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.C) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.C
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADDr_d(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.D) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.D
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADDr_e(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.E) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.E
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADDr_h(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.H) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.H
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADDr_l(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.L) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.L
	if emu.Registers.A & 255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADDr_hl(emu *Emulator){
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))) > 255{
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	if emu.Registers.A & 255 == 0{
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func ADDr_a(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)+int(emu.Registers.A) > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A += emu.Registers.A
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_b(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.B) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_c(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.C) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_d(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.D) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_e(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.E) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_h(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.H) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_l(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.L) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ADCr_a(emu *Emulator) {
	emu.Registers.F = 0
	var tmp uint16 = uint16(emu.Registers.A) + uint16(emu.Registers.A) + (uint16(emu.Registers.F&0x10) >> 4)
	if tmp > 255 {
		emu.Registers.F |= 0x10
	}
	emu.Registers.A = uint8(tmp & 0xFF)
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xA0
func ANDr_b(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.B
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ANDr_c(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.C
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ANDr_d(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.D
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ANDr_e(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.E
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ANDr_h(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.H
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ANDr_l(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.L
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ANDHL(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func ANDr_a(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.F |= 0x20
	emu.Registers.A &= emu.Registers.A
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_b(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.B
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_c(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.C
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_d(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.D
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_e(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.E
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_h(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.H
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_l(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.L
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func XORr_hl(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func XORr_a(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A ^= emu.Registers.A
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORr_b(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.B
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORr_c(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.C
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORr_d(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.D
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORr_e(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.E
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORr_h(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.H
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORr_l(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.L
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func ORHL(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func ORr_a(emu *Emulator) {
	emu.Registers.F = 0
	emu.Registers.A |= emu.Registers.A
	if (emu.Registers.A & 255) == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xB8
func CPr_b(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.B) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.B
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xB9
func CPr_c(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.C) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.C
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xBA
func CPr_d(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.D) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.D
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xBB
func CPr_e(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.E) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.E
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xBC
func CPr_h(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.H) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.H
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}

//0xBD
func CPr_l(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.L) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.L
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
func CPHL(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)- int(emu.MemoryRead(uint16(emu.Registers.H)<<8|uint16(emu.Registers.L))) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.MemoryRead(uint16(emu.Registers.H)<<8 | uint16(emu.Registers.L))
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 8
}
func CPr_a(emu *Emulator) {
	emu.Registers.F = 0
	if int(emu.Registers.A)-int(emu.Registers.A) < 0 {
		emu.Registers.F |= 0x10
	}
	tmp := emu.Registers.A
	tmp -= emu.Registers.A
	emu.Registers.F |= 0x40
	if tmp&255 == 0 {
		emu.Registers.F |= 0x80
	}
	emu.Registers.M = 1
	emu.Registers.T = 4
}
