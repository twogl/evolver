package main

import (
	"image/color"
	"log"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	circles        []*Circle
	board          *Board
	screenWidth    int
	screenHeight   int
	updateDelay    time.Duration
	lastUpdateTime time.Time
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

	//g.board.Draw(screen)
	for _, c := range g.circles {
		c.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth + 100, g.screenHeight + 400
}

func main() {
	// Create the game
	game := &Game{
		circles:      []*Circle{},
		board:        &Board{},
		screenWidth:  600,
		screenHeight: 600,
	}
	game.circles = loadCircles()

	// Run the game
	ebiten.SetWindowSize(game.screenWidth+100, game.screenHeight+400)
	ebiten.SetWindowTitle("Ebitengine Button Example")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func loadCircles() []*Circle {
	res := []*Circle{}

	for i := 0; i < 300; i++ {
		r := uint8(rand.IntN(256))
		g := uint8(rand.IntN(256))
		b := uint8(rand.IntN(256))
		x := 10*rand.IntN(60) - 5
		y := 10*rand.IntN(60) - 5
		u := rand.IntN(300)
		c := &Circle{
			updateDelay:    time.Duration(u) * time.Millisecond,
			lastUpdateTime: time.Now(),
			clr:            color.RGBA{r, g, b, 100},
			radius:         4,
			posX:           x,
			posY:           y,
		}
		res = append(res, c)
	}

	return res
}
