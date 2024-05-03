package models

type Chip8 struct {
	opcode      uint16
	memory      [4096]byte
	registers   [16]byte
	I           uint8
	pc          uint16
	gfx         [64 * 32]byte
	delay_timer byte
	sound_timer byte

	// Stack
	stack [16]uint8
	sp    uint8

	// Keypad
	keypad [16]byte
}

func (chip8 *Chip8) Initialize() {
	chip8.pc = 0x200
	chip8.opcode = 0
	chip8.I = 0
	chip8.sp = 0

}

func (chip8 *Chip8) clearMemory() {
	for i := 0; i < 4096; i++ {
		chip8.memory[i] = 0
	}
}

func (chip8 *Chip8) clearStack() {
	for i := 0; i < 16; i++ {
		chip8.stack[i] = 0
	}
}

func (chip8 *Chip8) clearRegisters() {
	for i := 0; i < 16; i++ {
		chip8.registers[i] = 0
	}
}
