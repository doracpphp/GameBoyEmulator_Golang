package mmu
type MMU struct{
	bios [256]uint8
}
func NewMMU() * MMU{
	mmu := new(MMU)
	return mmu
}
