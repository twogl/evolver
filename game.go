package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board        *Board
	circles      []*Circle
	screenWidth  int
	screenHeight int
}

func (g *Game) Update() error {
	// Update the button state
	for _, c := range g.circles {
		c.Update(g)
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{190, 190, 190, 255})

	//g.board.Draw(screen, g)

	for _, c := range g.circles {
		c.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
