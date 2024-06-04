package chip8

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Screen struct {
	width  int
	height int
	scale  int
}

func Initialize(height int, width int, scale int) Screen {
	screen := Screen{}

	screen.width = width
	screen.height = height
	screen.scale = scale

	return screen
}

func (screen *Screen) Clear(c8 Chip8) {
	for y := 0; y < screen.height; y++ {
		for x := 0; x < screen.width; x++ {
			c8.gfx[(y*64)+x] = 0
		}
	}
}

func (screen *Screen) Render(c8 Chip8) {
	for y := 0; y < screen.height; y++ {
		for x := 0; x < screen.width; x++ {
			if c8.gfx[(y*64)+x] == 0 {
				rl.DrawPixel(int32(x), int32(x), rl.Black)
			} else {
				rl.DrawPixel(int32(x), int32(x), rl.White)
			}
		}
	}

}
