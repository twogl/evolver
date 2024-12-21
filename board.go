package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
}

// |-----|-----|
// |---------|---------|

// 3 9 15 21 27
// 0 6 12 18 24

// 800 pixels
// 132  game dots
// 396x396 png

// drawCircle draws a filled circle at (cx, cy) with radius r and color clr.
func (g *Board) Draw(screen *ebiten.Image, game *Game) {

	for y := 6; y < game.screenHeight; y += 6 {
		for x := 0; x <= game.screenWidth; x++ {
			screen.Set(x, y, color.RGBA{1, 1, 1, 255})
		}
	}

	for x := 6; x < game.screenWidth; x += 6 {
		for y := 0; y <= game.screenHeight; y++ {
			screen.Set(x, y, color.RGBA{1, 1, 1, 255})
		}
	}

}
