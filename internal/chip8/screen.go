package chip8

import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type Screen struct {
	width    int
	height   int
	scale    int
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
}

func Initialize(height int, width int, scale int) Screen {
	screen := Screen{}

	screen.width = width
	screen.height = height
	screen.scale = scale

	sdl.Init(uint32(sdl.INIT_VIDEO))

	window, _ := sdl.CreateWindow("Chip 8 Emulator", 0, 0, int32(width)*int32(scale), int32(height)*int32(scale), uint32(sdl.WINDOW_SHOWN))

	renderer, _ := sdl.CreateRenderer(window, -1, uint32(sdl.RENDERER_ACCELERATED))

	texture, _ := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA8888), int(sdl.TEXTUREACCESS_STREAMING), int32(width), int32(height))

	screen.window = window
	screen.renderer = renderer
	screen.texture = texture

	return screen
}

func (screen *Screen) Clear(c8 Chip8) {
	for y := 0; y < screen.height; y++ {
		for x := 0; x < screen.width; x++ {
			c8.gfx[(y*32)+x] = 0
		}
	}
}

func (screen *Screen) Render(c8 Chip8) {
	videoPitch := screen.width * 2

	screen.texture.Update(nil, unsafe.Pointer(&c8.gfx), videoPitch)
	screen.renderer.Clear()
	screen.renderer.Copy(screen.texture, nil, nil)
	screen.renderer.Present()
}
