package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
)

type Game struct {
	grid           [40][10]string
	verticalLine   *ebiten.Image
	horizontalLine *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBoard(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	fmt.Println("Starting gotetris!")
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("gotetris")
	game := Game{}
	generateAssets(&game)
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}

func generateAssets(g *Game) {
	white := color.RGBA{255, 255, 255, 255}

	verticalLine := ebiten.NewImage(2, SCREEN_HEIGHT)
	verticalLine.Fill(white)
	g.verticalLine = verticalLine

	// We want grid to be approx 1/3 of screen, but also want grid squares to match up
	// 800/3 = 266.6666667. But there will be 10 vertical rows so 260 works better, width of 10 between each row
	gridWith := 260
	horizontalLine := ebiten.NewImage(gridWith, 2)
	horizontalLine.Fill(white)
	g.horizontalLine = horizontalLine
}

func drawBoard(g *Game, screen *ebiten.Image) {
	// We want grid to be approx 1/3 of screen, but also want grid squares to match up
	// 800/3 = 266.6666667. But there will be 10 vertical rows so 260 works better, width of 10 between each row
	width := SCREEN_WIDTH / 3
	gridCols := len(g.grid[0])
	gap := width / gridCols
	for i := 0; i <= gridCols; i += 1 {
		lineOp := &ebiten.DrawImageOptions{}
		lineOp.GeoM.Translate(float64(i*gap+width), 0.0)
		screen.DrawImage(g.verticalLine, lineOp)
	}

	gridRows := len(g.grid)
	for j := 0; j < gridRows; j += 1 {
		lineOp := &ebiten.DrawImageOptions{}
		lineOp.GeoM.Translate(float64(width), float64(j*gap))
		screen.DrawImage(g.horizontalLine, lineOp)
	}

}

// centerMargin returns the margin you would use if you wanted
// a line of length "inside" to be centered within a line of
// length "outside".
func centerMargin(outside, inside int) int {
	return (outside - inside) / 2
}
