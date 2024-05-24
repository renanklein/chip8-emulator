package chip8

type Screen struct {
	width  int
	height int
	scale  int
	pixels [][]int
}

func (screen *Screen) Clear() {
	for y := 0; y < screen.height; y++ {
		for x := 0; x < screen.width; x++ {
			screen.pixels[x][y] = 0
		}
	}
}
