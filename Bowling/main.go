package main

import (
	"fmt"

	"github.com/mathieu-keller/go-katas/Bowling/game"
)

func main() {
	game := game.NewGame()
	for {
		fmt.Println("Enter the number of pins knocked down: ")
		var pins int
		fmt.Scanln(&pins)
		game.AddRoll(pins)
		fmt.Println(game.Frames())
		fmt.Println(game.TotalScore())
		if game.Over() {
			fmt.Println("Game end!")
			break
		}
	}
}
