package register
type Register struct {
	//Register
	A,B,C,D,E,F,H,L uint8
	//Register pair
	AF,BC,DE,HL uint16

	M, T uint8
	//program counter and stack pointer
	SP,PC uint16	
}
//AF
func (reg *Register) setAF(value uint16){
	reg.A = uint8((value & 0xFF00) >> 8)
	reg.F = uint8((value & 0x00F0))
	reg.AF = (uint16(reg.A) << 8) | uint16(reg.F)
}
func (reg *Register)carry()uint8{
	return (reg.F >> 4) & 0x1
}
func (reg *Register) SetCarry(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 4) * -1)) | (v << 4)
}
func (reg *Register)halfCarry()uint8{
	return (reg.F >> 5) & 0x1
}

func (reg *Register) setHalfCarry(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 5) * -1)) | (v << 5)
}
func (reg *Register)subtract()uint8{
	return (reg.F >> 6) & 0x1
}
func (reg *Register) setSubtract(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 6) * -1)) | (v << 6)
}
func (reg *Register)zero()uint8{
	return (reg.F >> 7) & 0x1
}
func (reg *Register) setZero(value uint8){
	v := uint8(0)
	if value != 0{
		v=1;
	}
	reg.F = (reg.F & (^(1 << 7) * -1)) | (v << 7)
}