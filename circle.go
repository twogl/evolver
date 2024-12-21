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
	endX, endY     int
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

func (c *Circle) Update(game *Game) {

	if time.Since(c.lastUpdateTime) >= c.updateDelay {

		if c.posX < c.endX {
			c.posX += 6
		} else if c.posX > c.endX {
			c.posX -= 6
		} else if c.posY < c.endY {
			c.posY += 6
		} else if c.posY > c.endY {
			c.posY -= 6
		}

		c.lastUpdateTime = time.Now()
	}
}
