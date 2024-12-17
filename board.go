package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
}

// drawCircle draws a filled circle at (cx, cy) with radius r and color clr.
func (g *Board) Draw(screen *ebiten.Image) {

	for y := 10; y < 600; y += 10 {
		for x := 0; x <= 600; x++ {
			screen.Set(x, y, color.RGBA{1, 1, 1, 255})
		}
	}

	for x := 10; x < 600; x += 10 {
		for y := 0; y <= 600; y++ {
			screen.Set(x, y, color.RGBA{1, 1, 1, 255})
		}
	}

}
