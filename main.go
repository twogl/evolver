package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand/v2"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Open the image file
	file, err := os.Open("photo.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()
	fmt.Println("bounds: ", bounds)
	fmt.Println(bounds.Max.X)
	fmt.Println(bounds.Max.Y)

	//pixel := getAveragePixel(img)

	game := &Game{
		board:        &Board{},
		circles:      []*Circle{},
		screenWidth:  800,
		screenHeight: 800,
	}

	for x := 1; x < bounds.Max.X; x += 3 {
		for y := 1; y < bounds.Max.Y; y += 3 {
			pixel := getAveragePixel(x, y, img)
			x := (x - 1) / 3
			y := (y - 1) / 3
			dot := createCircle(x, y, pixel, game)
			game.circles = append(game.circles, dot)
		}
	}

	// Run the game
	ebiten.SetWindowSize(game.screenWidth, game.screenHeight)
	ebiten.SetWindowTitle("Ebitengine Button Example")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func createCircle(x, y int, pixel color.RGBA, game *Game) *Circle {
	u := rand.IntN(30)
	c := &Circle{
		updateDelay:    time.Duration(u) * time.Millisecond,
		lastUpdateTime: time.Now(),
		clr:            pixel,
		radius:         3,
		// posX:           3 + 6*x,
		// posY:           3 + 6*y,
		endX: 3 + 6*x,
		endY: 3 + 6*y,
	}
	position := rand.IntN(4)

	switch position {
	case 0:
		c.posX = 3
		c.posY = c.endY
	case 1:
		c.posX = c.endX
		c.posY = 3
	case 2:
		c.posX = 795
		c.posY = c.endY
	case 3:
		c.posX = c.endX
		c.posY = 795
	default:
		panic("unecpected position value")
	}

	return c
}

func loadCircles(game *Game) []*Circle {
	res := []*Circle{}

	for i := 0; i < 300; i++ {
		r := uint8(rand.IntN(256))
		g := uint8(rand.IntN(256))
		b := uint8(rand.IntN(256))
		x := 3 + 6*rand.IntN(133)
		y := 3 + 6*rand.IntN(133)
		u := rand.IntN(300)
		c := &Circle{
			updateDelay:    time.Duration(u) * time.Millisecond,
			lastUpdateTime: time.Now(),
			clr:            color.RGBA{r, g, b, 100},
			radius:         3,
			posX:           x,
			posY:           y,
		}
		res = append(res, c)
	}

	return res
}

var steps [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func getAveragePixel(x, y int, img image.Image) color.RGBA {
	col := img.At(x, y)
	red, green, blue, _ := col.RGBA()

	for i := 0; i < len(steps); i++ {
		col = img.At(x+steps[i][0], y+steps[i][1])
		r, g, b, _ := col.RGBA()
		red += r
		green += g
		blue += b
	}

	return color.RGBA{
		R: uint8((red / 9) / 257),
		G: uint8((green / 9) / 257),
		B: uint8((blue / 9) / 257),
		A: 255,
	}
	// Convert the color to RGBA

}
