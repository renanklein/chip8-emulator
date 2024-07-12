package main

import (
	"fmt"
	"os"

	"github.com/renanklein/chip8-emulator/internal/chip8"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	DISPLAY_WIDTH  = 64
	DISPLAY_HEIGHT = 32
	SCALE          = 30
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide a rom to be executed")
	}

	game_data := getRomData(os.Args[1])

	c8 := chip8.Chip8{}
	sc := chip8.Initialize(DISPLAY_HEIGHT, DISPLAY_WIDTH, SCALE)

	c8.Initialize(game_data)

	quit := false

	for !quit {
		c8.EmulationCycle()
		if c8.ShouldDraw() {
			sc.Render(c8, SCALE)
		}

		chip8.HandleKeyboard()
		sdl.Delay(1000 / 60)
	}
}

func getRomData(filename string) []byte {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		error_message := fmt.Sprintf("Could not load ROM, something went wrong: %s", err.Error())
		panic(error_message)
	}

	stat, errStat := file.Stat()

	if errStat != nil {
		panic("Something went wrong on reading file data")
	}

	buffer := make([]byte, stat.Size())

	file.Read(buffer)

	return buffer
}
