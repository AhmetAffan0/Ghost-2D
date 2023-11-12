package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	sprite1 *ebiten.Image
	sprite2 *ebiten.Image
	sprite3 *ebiten.Image
	sprite4 *ebiten.Image
)

func init() {
	img, _, err := ebitenutil.NewImageFromFile("Sprite-0001.png")
	if err != nil {
		panic(err)
	}
	sprite1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprite-0002.png")
	if err != nil {
		panic(err)
	}
	sprite2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprite-0003.png")
	if err != nil {
		panic(err)
	}
	sprite3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprite-0004.png")
	if err != nil {
		panic(err)
	}
	sprite4 = img
}

type char struct {
	x  int
	y  int
	vx int
	vy int
}

const (
	groundY = 250
	unit    = 10
)

func (c *char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.vx > 0 {
		c.vx -= 5
	} else if c.vx < 0 {
		c.vx += 5
	}
	if c.vy > 0 {
		c.vy -= 5
	} else if c.vy < 0 {
		c.vy += 5
	}
}

func (c *char) draw(screen *ebiten.Image) {
	s := sprite1

	if c.vx > 0 {
		s = sprite2
	} else if c.vx < 0 {
		s = sprite3
	}
	if c.vy > 0 {
		s = sprite1
	} else if c.vy < 0 {
		s = sprite4
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

type Game struct {
	ghost *char
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ghost.draw(screen)

	msg := fmt.Sprintf("TPS: %0.2f\nFloating Ghost.", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 1000, 950
}

func (g *Game) Update() error {
	if g.ghost == nil {
		g.ghost = &char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.ghost.vy -= 2 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.ghost.vy += 2 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.ghost.vx -= 2 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.ghost.vx += 2 * unit
	}
	g.ghost.update()
	return nil
}

func main() {
	ebiten.SetWindowSize(1000, 950)
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
