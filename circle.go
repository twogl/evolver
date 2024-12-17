package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Circle struct {
	updateDelay    time.Duration
	lastUpdateTime time.Time
	clr            color.Color
	radius         int
	posX, posY     int
	vX, vY         int
}

// drawCircle draws a filled circle at (cx, cy) with radius r and color clr.
func (g *Circle) Draw(screen *ebiten.Image) {
	for y := -g.radius; y <= g.radius; y++ {
		for x := -g.radius; x <= g.radius; x++ {
			if x*x+y*y <= g.radius*g.radius { // Circle equation: x² + y² <= r²
				screen.Set(g.posX+x, g.posY+y, g.clr)
			}
		}
	}
}

func (g *Circle) Update(game *Game) {

	if time.Since(g.lastUpdateTime) >= g.updateDelay {
		g.posX += g.vX
		g.posY += g.vY

		// Bounce off walls
		if g.posX-g.radius < 0 || g.posX+g.radius > game.screenWidth {
			g.vX = -g.vX
			g.posX += g.vX
		}
		if g.posY-g.radius < 0 || g.posY+g.radius > game.screenHeight {
			g.vY = -g.vY
			g.posY += g.vY
		}

		g.lastUpdateTime = time.Now()
	}
}
