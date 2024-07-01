package main

import (
	"fmt"
	"os"

	"github.com/renanklein/chip8-emulator/internal/chip8"
)

const DISPLAY_WIDTH = 64
const DISPLAY_HEIGHT = 32
const SCALE = 10

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
			sc.Render(c8)
		}
	}
}

func getRomData(filename string) []byte {
	data, err := os.ReadFile(filename)

	if err != nil {
		error_message := fmt.Sprintf("Could not load ROM, something went wrong: %s", err.Error())
		panic(error_message)
	}

	return data
}
