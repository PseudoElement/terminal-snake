package main

import (
	"log"

	"github.com/pseudoelement/terminal-snake/src/game"
)

func main() {
	g := game.NewDemoGame()
	if err := g.Run(); err != nil {
		g.Quit()
		log.Println(err)
	}
}
