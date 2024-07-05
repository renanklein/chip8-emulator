package chip8

import (
	"math/rand"
)

type Chip8 struct {
	memory      [4096]uint8
	registers   [16]uint8
	I           uint16
	pc          uint16
	gfx         [64 * 32]uint32
	delay_timer uint8
	sound_timer uint8

	draw_screen bool

	// Stack
	stack [16]uint16
	sp    uint8

	// Keypad
	keypad [16]uint8
}

func (chip8 *Chip8) ShouldDraw() bool {
	return chip8.draw_screen
}

func (chip8 *Chip8) SetDraw(draw bool) {
	chip8.draw_screen = draw
}

func (chip8 *Chip8) Initialize(gameData []byte) {
	chip8.pc = 0x200
	chip8.I = 0
	chip8.sp = 0

	chip8.clearMemory()
	chip8.clearRegisters()
	chip8.clearStack()
  chip8.loadFontset()
	chip8.LoadRom(gameData)
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

func (chip8 *Chip8) loadFontset() {
	for i := 0; i < len(fontset); i++ {
		chip8.memory[i] = fontset[i]
	}
}

func (chip8 *Chip8) LoadRom(data []byte) {
	if (4096 - 512) < len(data) {
		panic("Fatal Error ! Cannot load ROM, it's too big")
	}

	for i := 0; i < len(data); i++ {
		chip8.memory[i+512] = data[i]
	}
}

func (chip8 *Chip8) FetchOpcode() uint16 {
	first_part := (uint16(chip8.memory[chip8.pc]) << 8)
	opcode := first_part | uint16(chip8.memory[chip8.pc+1])
	return opcode
}

func (chip8 *Chip8) updateTimers() {
	if chip8.delay_timer > 0 {
		chip8.delay_timer--
	}

	if chip8.sound_timer > 0 {
		chip8.sound_timer--
	}
}


func (chip8 *Chip8) executeOpcodes(x uint16, y uint16, opcode uint16) {
	switch opcode & 0xF000 {

	case 0x0000:
		switch opcode {

		case 0x00E0:
			for i := 0; i < len(chip8.gfx); i++ {
				chip8.gfx[i] = 0x0
			}

			chip8.draw_screen = true
			chip8.pc += 2

		case 0x00EE:
			chip8.sp--
			chip8.pc = chip8.stack[chip8.sp]
		}

	case 0x1000:
		chip8.pc = opcode & 0x0FFF

	case 0x2000:
		chip8.stack[chip8.sp] = chip8.pc
		chip8.sp++
		chip8.pc = opcode & 0x0FFF

	case 0x3000:
		if uint16(chip8.registers[x]) == opcode&0x00FF {
			// jumps instruction
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
	case 0x4000:
		if (uint16(chip8.registers[x])) != opcode&0x00FF {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}

	case 0x5000:

		if (uint16(chip8.registers[x])) == (uint16(chip8.registers[y])) {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
	case 0x6000:
		chip8.registers[x] = uint8(opcode & 0x00FF)
		chip8.pc += 2

	case 0x7000:
		result := uint16(chip8.registers[x]) + opcode&0x00FF

		if result > 255 {
			result -= 256
		}

		chip8.registers[x] += uint8(result)

		chip8.pc += 2

	case 0x8000:
		switch opcode & 0xF {

		case 0x0:
			chip8.registers[x] = chip8.registers[y]

			chip8.pc += 2

		case 0x1:
			chip8.registers[x] |= chip8.registers[y]

			chip8.pc += 2

		case 0x2:
			chip8.registers[x] &= chip8.registers[y]

			chip8.pc += 2

		case 0x3:
			chip8.registers[x] ^= chip8.registers[y]

			chip8.pc += 2

		case 0x4:
			result := chip8.registers[x] + chip8.registers[y]

			if result > 255 {
				chip8.registers[0xf] = 1
			} else {
				chip8.registers[0xf] = 0
			}

			chip8.registers[x] = result & 0xFF

			chip8.pc += 2

		case 0x5:
			if chip8.registers[y] > chip8.registers[x] {
				chip8.registers[0xf] = 0
			} else {
				chip8.registers[0xf] = 1
			}

			result := chip8.registers[x] - chip8.registers[y]

			chip8.registers[x] = result

			chip8.pc += 2

		case 0x6:
			chip8.registers[0xF] = (chip8.registers[x] & 0x1)
			chip8.registers[x] = (chip8.registers[x] >> 1)

			chip8.pc += 2

		case 0x7:
			if chip8.registers[x] > chip8.registers[y] {
				chip8.registers[0xf] = 0
			} else {
				chip8.registers[0xf] = 1
			}

			result := chip8.registers[y] - chip8.registers[x]

			chip8.registers[x] = result

			chip8.pc += 2

		case 0xE:
			chip8.registers[0xF] = chip8.registers[x] >> 7
			chip8.registers[x] = chip8.registers[x] << 1

			chip8.pc += 2
		}

	case 0x9000:
		if chip8.registers[x] != chip8.registers[y] {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}

	case 0xA000:
		chip8.I = (opcode & 0x0FFF)
		chip8.pc += 2
		chip8.updateTimers()

	case 0xB000:
		opcode_addr := (opcode & 0x0FFF)
		chip8.pc = opcode_addr + uint16(chip8.registers[0])

	case 0xC000:
		rand_number := uint8(rand.Intn(256))

		chip8.registers[x] = rand_number & uint8(opcode&0x00FF)
		chip8.pc += 2

	case 0xD000:
		height := (opcode & 0x000F)

		chip8.registers[0xF] = 0

		xPos := chip8.registers[x] % 64
		yPos := chip8.registers[y] % 32

		// Screen line loop
		for line := 0; line < int(height); line++ {
			spriteByte := chip8.memory[chip8.I+uint16(line)]

			// Pixel from each line loop
			for col := 0; col < 8; col++ {
				spritePixel := spriteByte & (0x80 >> col)
				screenPixel := &chip8.gfx[(yPos+uint8(line))*64+(xPos+uint8(col))]
				if spritePixel != 0 {
					if *screenPixel == 0xFFFFFFFF {
						chip8.registers[0xF] = 1
					}

					*screenPixel ^= 0xFFFFFFFF
				}
			}
		}

		chip8.draw_screen = true
		chip8.pc += 2

	case 0xE000:
		switch opcode & 0xFF {
		case 0x9E:
			if IsPressed(int(chip8.registers[x])) {
				chip8.pc += 4
			} else {
				chip8.pc += 2
			}
		case 0xA1:
			if !IsPressed(int(chip8.registers[x])) {
				chip8.pc += 4
			} else {
				chip8.pc += 2
			}

		}
	case 0xF000:
		switch opcode & 0xFF {
		case 0x07:
			chip8.registers[x] = chip8.delay_timer
			chip8.pc += 2

		case 0x0A:
			for i := 0; i < len(GetKeys()); i++ {
				if IsPressed(i) {
					chip8.registers[x] = uint8(i)
				}
			}

			chip8.pc += 2

		case 0x15:
			chip8.delay_timer = chip8.registers[x]
			chip8.pc += 2

		case 0x18:
			chip8.sound_timer = chip8.registers[x]
			chip8.pc += 2

		case 0x1E:
			chip8.I += uint16(chip8.registers[x])
			chip8.pc += 2

		case 0x29:
			digit := chip8.registers[x]

			chip8.I += uint16(digit * 5)
			chip8.pc += 2

		case 0x33:
			number := chip8.registers[x]
			chip8.memory[chip8.I] = number / 100
			chip8.memory[chip8.I+1] = (number % 100) / 10
			chip8.memory[chip8.I+2] = (number % 100) % 10

			chip8.pc += 2

		case 0x55:
			var i uint16
			for i = 0; i <= x; i++ {
				memory_addr := chip8.I + i
				chip8.memory[memory_addr] = chip8.registers[i]
			}

			chip8.pc += 2

		case 0x65:
			var i uint16
			for i = 0; i <= x; i++ {
				chip8.registers[i] = chip8.memory[chip8.I+i]
			}

			chip8.pc += 2

		}

	}
}

func (chip8 *Chip8) EmulationCycle() {
	opcode := chip8.FetchOpcode()

	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

  chip8.executeOpcodes(x, y, opcode)

	chip8.updateTimers()
}
