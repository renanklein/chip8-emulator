package models

import (
	"fmt"
	"math/rand"
)

var chip8_fontset = [80]uint16{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

type Chip8 struct {
	memory      [4096]uint8
	registers   [16]uint8
	I           uint16
	pc          uint16
	gfx         [64 * 32]uint8
	delay_timer uint8
	sound_timer uint8

	// Stack
	stack [16]uint16
	sp    uint8

	// Keypad
	keypad [16]uint8
}

func (chip8 *Chip8) Initialize() {
	chip8.pc = 0x200
	chip8.I = 0
	chip8.sp = 0

	chip8.clearMemory()
	chip8.clearRegisters()
	chip8.clearStack()
	chip8.LoadRom()
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

func (chip8 *Chip8) LoadRom() {
	//TODO
}

func (chip8 *Chip8) FetchOpcode() uint16 {
	first_part := chip8.memory[chip8.pc] << 8
	opcode := first_part | chip8.memory[chip8.pc+1]
	return uint16(opcode)
}

func (chip8 *Chip8) updateTimers() {
	if chip8.delay_timer > 0 {
		chip8.delay_timer--
	}

	if chip8.sound_timer > 0 && chip8.sound_timer == 1 {
		chip8.sound_timer--
	}
}

func (chip8 *Chip8) EmulationCycle() {
	opcode := chip8.FetchOpcode()

	switch opcode & 0xF000 {

	case 0x1000:
		chip8.pc = opcode & 0x0FFF

	case 0x2000:
		chip8.stack[chip8.sp] = chip8.pc
		chip8.sp++
		chip8.pc = opcode & 0x0FFF

	case 0x3000:
		register_number := opcode & 0x0F00
		if uint16(chip8.registers[register_number]) == opcode&0x00FF {
			// jumps instruction
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
	case 0x4000:
		register_number := opcode & 0x0F00
		if uint16(chip8.registers[register_number]>>8) != opcode&0x00FF {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}

	case 0x5000:
		first_register_index := opcode & 0x0F00
		second_register_index := opcode & 0x00F0

		if (chip8.registers[first_register_index] >> 8) == (chip8.registers[second_register_index] >> 4) {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
	case 0x6000:
		register_index := (opcode & 0x0F00) >> 8
		chip8.registers[register_index] = (opcode & 0x00FF)

	case 0x7000:
		register_index := (opcode & 0x0F00) >> 8
		result := chip8.registers[register_index] + (opcode & 0x00FF)

		if result > 256 {
			result = result - 256
		}

		chip8.registers[register_index] = result

		chip8.pc += 2

	case 0xA000:
		chip8.I = (opcode & 0x0FFF)
		chip8.pc += 2
		chip8.updateTimers()

	case 0xB000:
		opcode_addr := (opcode & 0x0FFF)
		chip8.pc = opcode_addr + uint16(chip8.registers[0])

	case 0xC000:
		register_index := (opcode & 0x0F00) >> 8
		rand_number := rand.Intn(256)

		chip8.registers[register_index] = rand_number & (opcode & 0x00FF)
		chip8.pc += 2

	case 0xD000:
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		height := (opcode & 0x000F)

		chip8.registers[0xF] = 0

		// Screen line loop
		for yLine := 0; yLine < int(height); yLine++ {
			line := chip8.memory[chip8.I+uint16(yLine)]

			//Pixel from each line loop
			for xLine := 0; xLine < 8; xLine++ {

			}
		}

	default:
		fmt.Printf("Unknown opcode: %d", opcode)

	}
}
