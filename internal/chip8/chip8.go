package chip8

import (
	"fmt"
	"math/rand"
)

const VIDEO_WIDTH = 64
const VIDEO_HEIGHT = 32

type Chip8 struct {
	memory      [4096]uint8
	registers   [16]uint8
	I           uint16
	pc          uint16
	gfx         [32][64]byte
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

func (chip8 *Chip8) GetDisplayVector() [32][64]byte {
	return chip8.gfx
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
				for j := 0; j < len(chip8.gfx[i]); j++ {
					chip8.gfx[i][j] = 0x0
				}
			}
			chip8.draw_screen = true

		case 0x00EE:
			chip8.sp--
			chip8.pc = chip8.stack[chip8.sp]
		}

	case 0x1000:
		addr := opcode & 0x0FFF
		chip8.pc = addr

	case 0x2000:
		addr := opcode & 0x0FFF
		chip8.stack[chip8.sp] = chip8.pc
		chip8.sp++
		chip8.pc = addr

	case 0x3000:
		b := (opcode & 0x00FF)
		if uint16(chip8.registers[x]) == b {
			// jumps instruction
			chip8.pc += 2
		}
	case 0x4000:
		b := (opcode & 0x00FF)

		if (uint16(chip8.registers[x])) != b {
			chip8.pc += 2
		}

	case 0x5000:

		if (uint16(chip8.registers[x])) == (uint16(chip8.registers[y])) {
			chip8.pc += 2
		}
	case 0x6000:
		b := opcode & 0x00FF
		chip8.registers[x] = uint8(b)
		chip8.pc += 2

	case 0x7000:
		result := uint16(chip8.registers[x]) + opcode&0x00FF

		if result > 255 {
			result -= 256
		}

		chip8.registers[x] = uint8(result)

	case 0x8000:
		switch opcode & 0xF {

		case 0x0:
			chip8.registers[x] = chip8.registers[y]

		case 0x1:
			chip8.registers[x] |= chip8.registers[y]

		case 0x2:
			chip8.registers[x] &= chip8.registers[y]

		case 0x3:
			chip8.registers[x] ^= chip8.registers[y]

		case 0x4:
			result := chip8.registers[x] + chip8.registers[y]

			if result > 255 {
				chip8.registers[0xf] = 1
			} else {
				chip8.registers[0xf] = 0
			}

			chip8.registers[x] = result & 0xFF

		case 0x5:
			if chip8.registers[y] > chip8.registers[x] {
				chip8.registers[0xf] = 0
			} else {
				chip8.registers[0xf] = 1
			}

			result := chip8.registers[x] - chip8.registers[y]

			chip8.registers[x] = result

		case 0x6:
			chip8.registers[0xF] = (chip8.registers[x] & 0x1)
			chip8.registers[x] = (chip8.registers[x] >> 1)

		case 0x7:
			if chip8.registers[x] > chip8.registers[y] {
				chip8.registers[0xf] = 0
			} else {
				chip8.registers[0xf] = 1
			}

			result := chip8.registers[y] - chip8.registers[x]

			chip8.registers[x] = result

		case 0xE:
			chip8.registers[0xF] = (chip8.registers[x] & 0x80) >> 7
			chip8.registers[x] = chip8.registers[x] << 1

		}

	case 0x9000:
		if chip8.registers[x] != chip8.registers[y] {
			chip8.pc += 2
		}

	case 0xA000:
		chip8.I = (opcode & 0x0FFF)

	case 0xB000:
		opcode_addr := (opcode & 0x0FFF)
		chip8.pc = opcode_addr + uint16(chip8.registers[0])

	case 0xC000:
		rand_number := uint8(rand.Intn(256))

		chip8.registers[x] = rand_number & uint8(opcode&0x00FF)

	case 0xD000:
		height := (opcode & 0x000F)

		chip8.registers[0xF] = 0

		var line uint16 = 0
		var col uint16 = 0

		// Screen line loop
		for line = 0; line < height; line++ {
			spriteByte := chip8.memory[chip8.I+line]

			// Pixel from each line loop
			for col = 0; col < 8; col++ {
				spritePixel := spriteByte & (0x80 >> col)
				if spritePixel != 0 {
					if chip8.gfx[y+line][x+col] == 1 {
						chip8.registers[0xF] = 1
					}
					chip8.gfx[(y + line)][x+col] ^= 1
				}
			}
		}

		chip8.draw_screen = true

	case 0xE000:
		switch opcode & 0xFF {
		case 0x9E:
			if IsPressed(int(chip8.registers[x])) {
				chip8.pc += 2
			}
		case 0xA1:
			if !IsPressed(int(chip8.registers[x])) {
				chip8.pc += 2
			}

		}
	case 0xF000:
		switch opcode & 0xFF {
		case 0x07:
			chip8.registers[x] = chip8.delay_timer

		case 0x0A:
			isPressed := false
			for i := 0; i < len(GetKeys()); i++ {
				if IsPressed(i) {
					chip8.registers[x] = uint8(i)
					isPressed = true
				}
			}

			if isPressed {
				fmt.Println("Nothing was pressed !!!!")
				chip8.pc -= 2
			}

		case 0x15:
			chip8.delay_timer = chip8.registers[x]

		case 0x18:
			chip8.sound_timer = chip8.registers[x]

		case 0x1E:
			chip8.I += uint16(chip8.registers[x])

		case 0x29:
			digit := chip8.registers[x]

			chip8.I = 0x50 + uint16(digit*5)

		case 0x33:
			number := chip8.registers[x]

			chip8.memory[chip8.I+2] = number % 10
			number /= 10

			chip8.memory[chip8.I+1] = number % 10
			number /= 10

			chip8.memory[chip8.I] = number % 10

		case 0x55:
			var i uint16
			for i = 0; i <= x; i++ {
				memory_addr := chip8.I + i
				chip8.memory[memory_addr] = chip8.registers[i]
			}

		case 0x65:
			var i uint16
			for i = 0; i <= x; i++ {
				chip8.registers[i] = chip8.memory[chip8.I+i]
			}

		}

	}
}

func (chip8 *Chip8) EmulationCycle() {
	opcode := chip8.FetchOpcode()

	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	chip8.executeOpcodes(x, y, opcode)
	chip8.pc += 2

	chip8.updateTimers()
}
