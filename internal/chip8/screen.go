package chip8

type Screen struct {
	width  int
	height int
	scale  int
	pixels [][]int
}

func Initialize(height int, width int, scale int) Screen {
	screen := Screen{}

	screen.width = width
	screen.height = height
	screen.scale = scale
	screen.pixels = make([][]int, height)

	for i := 0; i < width; i++ {
		screen.pixels[i] = make([]int, screen.width)
	}

	return screen
}

func (screen *Screen) Clear() {
	for y := 0; y < screen.height; y++ {
		for x := 0; x < screen.width; x++ {
			screen.pixels[x][y] = 0
		}
	}
}

func (screen *Screen) Render() {

}
