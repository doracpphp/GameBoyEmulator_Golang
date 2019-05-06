package gameboy

import "fmt"

func (emu *Emulator) MemoryRead(address uint16) uint8 {
	return emu.Memory.Read(address)
}
func (emu *Emulator) MemoryWrite(address uint16, value uint8) {
	if address >= 0x8000 {
		emu.Memory.Write(address, value)
	}
}

func (emu *Emulator) Debug() {
	//fmt.Printf("A : %08b  D : %08b\nB : %08b  E : %08b \nC : %08b \n",emu.Registers.A,emu.Registers.D,emu.Registers.B,emu.Registers.E,emu.Registers.C)
	fmt.Printf("------------------------------------\nA : %08b  B : %08b  C : %08b\nD : %08b  E : %08b  H : %08b  L : %08b\n",
		emu.Registers.A, emu.Registers.B, emu.Registers.C,
		emu.Registers.D, emu.Registers.E, emu.Registers.H, emu.Registers.L)
	fmt.Printf("SP : %04X  PC : %04X  Flag : %08b\n------------------------------------\n", emu.Registers.SP, emu.Registers.PC, emu.Registers.F)

}

var instructionSet = map[uint8]Instruction{
	//0x00
	0x00: NOP, 0x01: LDBCnn, 0x02: LDBCr_a,
	0x03: INCBC, 0x04: INCB, 0x05: DECB,
	0x06: LDr_bn, 0x07: RLCA, 0x08: LDmmSP,
	0x09: ADDHLBC, 0x0A: LDAmBC, 0x0B: DECBC,
	0x0C: INCC, 0x0D: DECC, 0x0E: LDr_bn,
	0x0F: RRCA,
	//0x10
	0x10: STOP, 0x11: LDDEnn, 0x12: LDDEr_a,
	0x13: INCDE, 0x14: INCD, 0x15: DECD,
	0x16: LDr_dn, 0x17: RLA, 0x18:JRn,
	0x19: ADDHLDE, 0x1A: LDAmDE,0x1B: DECDE,
	0x1C: INCE, 0x1D: DECE,
	//0x20
	0x20: JRNZn, 0x21: LDHLnn, 0x22: LDIHLA,
	0x23: INCHL, 0x24: INCH, 0x25: DECH,
	0x26: LDr_hn,0x27: DAA, 0x28:JRZn,
	0x29:ADDHLHL, 0x2A:LDAHLI, 0x2B: DECHL,
	0x2C: INCL, 0x2D: DECL, 0x2E: LDr_ln,
	0x2F: CPL,
	//0x30
	0x30: JRNCn, 0x31: LDSPnn, 0x32: LDHLDA,
	0x33: INCSP, 0x34: INCHLm, 0x35: DECHLm,
	0x36: LDHLmn, 0x37: SCF, 0x38: JRCn,
	0x39: ADDHLSP, 0x3A: LDAHLD, 0x3B: DECSP,
	0x3C: INCA, 0x3D: DECA, 0x3E: LDr_an,
	0x3F: CCF,
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
	//0x60
	0x60: LDrr_hb, 0x61: LDrr_hc, 0x62: LDrr_hd,
	0x63: LDrr_he, 0x64: LDrr_hh, 0x65: LDrr_hl,
	0x66: LDrHL_h, 0x67: LDrr_ha, 0x68: LDrr_lb,
	0x69: LDrr_lc, 0x6A: LDrr_ld, 0x6B: LDrr_le,
	0x6C: LDrr_lh, 0x6D: LDrr_ll, 0x6E: LDrHL_l,
	0x6F: LDrr_la,
	//0x70
	0x70: LDHLmr_b, 0x71: LDHLmr_c, 0x72: LDHLmr_d,
	0x73: LDHLmr_e, 0x74: LDHLmr_h, 0x75: LDHLmr_l,
	0x76: HALT, 0x77: LDHLmr_a, 0x78: LDrr_ab,
	0x79: LDrr_ac, 0x7A: LDrr_ad, 0x7B: LDrr_ae,
	0x7C: LDrr_ah, 0x7D: LDrr_al, 0x7E: LDrHL_a,
	0x7F: LDrr_aa,
	//0x80
	0x80: ADDr_b, 0x81: ADDr_c, 0x82: ADDr_d,
	0x83: ADDr_e, 0x84: ADDr_h, 0x85: ADDr_l,
	0x86: ADDr_hl, 0x87: ADDr_a, 0x88: ADCr_b,
	0x89: ADCr_c, 0x8A: ADCr_d, 0x8B: ADCr_e,
	0x8C: ADCr_h, 0x8D: ADCr_l, 0x8E: ADCHL,
	0x8F: ADCr_a,
	//0x90
	0x90: SUBr_b, 0x91: SUBr_c, 0x92: SUBr_d,
	0x93: SUBr_e, 0x94: SUBr_h, 0x95: SUBr_l,
	0x96: SUBr_hl,0x97: SUBr_a, 0x98: SBCr_b,
	0x99: SBCr_c, 0x9A: SBCr_d, 0x9B: SUBr_e,
	0x9C: SBCr_h, 0x9D: SBCr_l, 0x9E: SUBr_hl,
	0x9F: SBCr_a,
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
	//0xC0
	0xC0: RETNZ, 0xC1: POPBC, 0xC2: JPNZnn,
	0xC3: JPnn, 0xC4:CALLNZnn,0xC5: PUSHBC,
	0xC6: ADDAn, 0xC7: RST00, 0xC8: RETZ,
	0xC9: RET, 0xCA: JPZnn, 0xCB: PREFIXCB,
	0xCC: CALLZnn, 0xCD: CALL, 0xCE: ADCAn,
	0xCF: RST08,
	//0xD0
	0xD0: RETNC, 0xD1: POPDE, 0xD2: JPNCnn,
	0xD3: void, 0xD4: CALLNCnn, 0xD5: PUSHDE,
	0xD6: SUBn, 0xD7: RST10, 0xD8: RETC,
	0xD9: RETI, 0xDA: JPCnn, 0xDB: void,
	0xDC: CALLCnn, 0xDD: void, 0xDE:SBCn,
	0xDF: RST18,
}
var prefixset = map[uint8]Instruction{
	0x00: RLCB, 0x01: RLCC, 0x02: RLCD,
	0x03: RLCE, 0x04: RLCH, 0x05: RLCL,
	0x06: RLCHLm, 0x07: RLCr_A,
}