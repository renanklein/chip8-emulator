package chip8

import "azul3d.org/engine/keyboard"

var keys [16]bool

var fontset = [80]uint16{
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

func KeyPressed(keyCode int) {
	switch keyCode {
	case int(keyboard.One):
		keys[1] = true
	case int(keyboard.Two):
		keys[2] = true
	case int(keyboard.Three):
		keys[3] = true
	case int(keyboard.Four):
		keys[0xA] = true
	case int(keyboard.Q):
		keys[4] = true
	case int(keyboard.W):
		keys[5] = true
	case int(keyboard.E):
		keys[6] = true
	case int(keyboard.R):
		keys[0xD] = true
	case int(keyboard.A):
		keys[7] = true
	case int(keyboard.S):
		keys[8] = true
	case int(keyboard.D):
		keys[9] = true
	case int(keyboard.F):
		keys[0xE] = true
	case int(keyboard.Z):
		keys[0xA] = true
	case int(keyboard.X):
		keys[0] = true
	case int(keyboard.C):
		keys[0xB] = true
	case int(keyboard.V):
		keys[0xF] = true
	}

}

func KeyReleased(keycode int) {
	switch keycode {
	case int(keyboard.One):
		keys[1] = false
	case int(keyboard.Two):
		keys[2] = false
	case int(keyboard.Three):
		keys[3] = false
	case int(keyboard.Four):
		keys[0xA] = false
	case int(keyboard.Q):
		keys[4] = false
	case int(keyboard.W):
		keys[5] = false
	case int(keyboard.E):
		keys[6] = false
	case int(keyboard.R):
		keys[0xD] = false
	case int(keyboard.A):
		keys[7] = false
	case int(keyboard.S):
		keys[8] = false
	case int(keyboard.D):
		keys[9] = false
	case int(keyboard.F):
		keys[0xE] = false
	case int(keyboard.Z):
		keys[0xA] = false
	case int(keyboard.X):
		keys[0] = false
	case int(keyboard.C):
		keys[0xB] = false
	case int(keyboard.V):
		keys[0xF] = false
	}
}

func IsPressed(index int) bool {
	return keys[index]
}

func GetKeys() [16]bool {
	return keys
}
