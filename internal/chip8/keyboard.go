package chip8

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var keys [16]bool

var fontset = [80]uint8{
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

func HandleKeyboard() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch eventType := event.(type) {
		case *sdl.QuitEvent:
			os.Exit(0)

		case *sdl.KeyboardEvent:
			if eventType.Type == sdl.KEYUP {
				switch eventType.Keysym.Sym {
				case sdl.K_1:
					keys[0x1] = false
				case sdl.K_2:
					keys[0x2] = false
				case sdl.K_3:
					keys[0x3] = false
				case sdl.K_4:
					keys[0xC] = false
				case sdl.K_q:
					keys[0x4] = false
				case sdl.K_w:
					keys[0x5] = false
				case sdl.K_e:
					keys[0x6] = false
				case sdl.K_r:
					keys[0xD] = false
				case sdl.K_a:
					keys[0x7] = false
				case sdl.K_s:
					keys[0x8] = false
				case sdl.K_d:
					keys[0x9] = false
				case sdl.K_f:
					keys[0xE] = false
				case sdl.K_z:
					keys[0xA] = false
				case sdl.K_x:
					keys[0x0] = false
				case sdl.K_c:
					keys[0xB] = false
				case sdl.K_v:
					keys[0xF] = false
				}
			} else if eventType.Type == sdl.KEYDOWN {
				switch eventType.Keysym.Sym {
				case sdl.K_1:
					keys[0x1] = true
				case sdl.K_2:
					keys[0x2] = true
				case sdl.K_3:
					keys[0x3] = true
				case sdl.K_4:
					keys[0xC] = true
				case sdl.K_q:
					keys[0x4] = true
				case sdl.K_w:
					keys[0x5] = true
				case sdl.K_e:
					keys[0x6] = true
				case sdl.K_r:
					keys[0xD] = true
				case sdl.K_a:
					keys[0x7] = true
				case sdl.K_s:
					keys[0x8] = true
				case sdl.K_d:
					keys[0x9] = true
				case sdl.K_f:
					keys[0xE] = true
				case sdl.K_z:
					keys[0xA] = true
				case sdl.K_x:
					keys[0x0] = true
				case sdl.K_c:
					keys[0xB] = true
				case sdl.K_v:
					keys[0xF] = true

				}
			}
		}
	}
}

func IsPressed(index uint8) bool {
	return keys[index]
}

func GetKeys() [16]bool {
	return keys
}
