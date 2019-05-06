package main

import "github.com/doracpphp/GameBoyEmulator_Golang/gameboy"
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)
import "fmt"
func main() {
	cpu := gameboy.NewGameboy()

	/*for {
		opcode := cpu.Memory[cpu.Registers.PC]
		cpu.Registers.PC+=1
		fmt.Printf("Opcode : %04X\n",opcode)
		cpu.Inst[opcode](cpu)
		cpu.Debug()
	}
	*/
	cpu.Registers.A=0x5
	cpu.Registers.B=0x4
	cpu.Registers.SP=0x04
	cpu.Registers.PC=0x1
	cpu.Registers.F |= 0x10
	cpu.Debug()
	cpu.Inst[0x51](cpu)
	cpu.Debug()

	cart,err := cpu.Memory.Cartridge.NewCartridgeFile("/Users/main/Documents/LT_Program/GameBoyEmu/src/Tetris (World) (Rev A).gb")
	if err != nil{
		fmt.Println("error ",err)
	}
	fmt.Println(cart.GetName())
	fmt.Printf("%08b\n",cpu.Memory.Read(0xFF00))
	pixelgl.Run(run)
}
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Gameboy Emulator",
		Bounds: pixel.R(0, 0, 160, 140),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}