package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Screen struct {
	width    int
	height   int
	scale    int
	window   *sdl.Window
	renderer *sdl.Renderer
}

func Initialize(height int, width int, scale int) Screen {
	screen := Screen{}

	screen.width = width
	screen.height = height
	screen.scale = scale

	sdl.Init(uint32(sdl.INIT_EVERYTHING))

	window, _ := sdl.CreateWindow("Chip 8 Emulator", 0, 0, int32(width)*int32(scale), int32(height)*int32(scale), uint32(sdl.WINDOW_SHOWN))

	renderer, _ := sdl.CreateRenderer(window, -1, uint32(sdl.RENDERER_ACCELERATED))

	screen.window = window
	screen.renderer = renderer

	return screen
}

func (screen *Screen) Clear(c8 Chip8) {
	for i := 0; i < len(c8.gfx); i++ {
		for j := 0; j < len(c8.gfx[i]); j++ {
			c8.gfx[i][j] = 0x0
		}
	}
}

func (screen *Screen) Render(c8 Chip8, modif int) {
	screen.renderer.SetDrawColor(255, 255, 0, 255)
	screen.renderer.Clear()

	for j := 0; j < len(c8.gfx); j++ {
		for i := 0; i < len(c8.gfx[j]); i++ {
			if c8.gfx[j][i] != 0 {
				screen.renderer.SetDrawColor(255, 255, 0, 255)
			} else {
				screen.renderer.SetDrawColor(255, 0, 0, 255)
			}

			screen.renderer.FillRect(&sdl.Rect{
				Y: int32(j) * int32(modif),
				X: int32(i) * int32(modif),
				W: int32(modif),
				H: int32(modif),
			})
		}
	}

	screen.renderer.Present()
}
