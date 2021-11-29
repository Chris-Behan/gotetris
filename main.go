package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	fmt.Println("Starting gotetris!")
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("gotetris")
	game := Game{}
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
