package main

import (
	"fmt"
	"os"
)

func main() {

}

func getRomData(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		error_message := fmt.Sprintf("Could not load ROM, something went wrong: %s", err.Error())
		panic(error_message)
	}

	defer file.Close()

}
