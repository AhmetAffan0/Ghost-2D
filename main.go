package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"math/rand"
	"time"

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
	rand.Seed(time.Now().UnixNano())
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
	s  *ebiten.Image
}

const (
	groundY = 250
	unit    = 10
)

var (
	random = rand.Intn(2)
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
	c.s = sprite1

	if c.vx > 0 {
		c.s = sprite2
	} else if c.vx < 0 {
		c.s = sprite3
	}
	if c.vy > 0 {
		c.s = sprite1
	} else if c.vy < 0 {
		c.s = sprite4
	}

	c.x += 9000
	c.x %= 9000

	c.y += 8000
	c.y %= 8000
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(c.s, op)

	opb := &ebiten.DrawImageOptions{}
	opb.GeoM.Scale(2, 2)
	opb.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit-900)
	screen.DrawImage(c.s, opb)

	opg := &ebiten.DrawImageOptions{}
	opg.GeoM.Scale(2, 2)
	opg.GeoM.Translate(float64(c.x)/unit+1000, float64(c.y)/unit)
	screen.DrawImage(c.s, opg)

	opc := &ebiten.DrawImageOptions{}
	opc.GeoM.Scale(2, 2)
	opc.GeoM.Translate(float64(c.x)/unit+1500, float64(c.y)/unit-900)
	screen.DrawImage(c.s, opc)

}

type Game struct {
	ghost *char
	yes   bool
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.ghost.draw(screen)

	op := &ebiten.DrawImageOptions{}

	rect := ebiten.NewImage(50, 50)
	rect.Fill(color.Black)

	op.GeoM.Scale(4.2, 3.8)
	screen.DrawImage(rect, op)

	var RectPos1, RectPos2 []int

	switch random {
	case 0:
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(500, 500)
		op.GeoM.Scale(1, 1)
		screen.DrawImage(rect, op)

		if rect.Bounds().Min.X+3000+50*unit <= g.ghost.x-50*unit && g.ghost.x-50*unit < rect.Bounds().Max.X+4450+50*unit && rect.Bounds().Min.Y-1000+groundY*unit <= g.ghost.y-groundY*unit && g.ghost.y-groundY*unit < rect.Bounds().Max.Y+500+groundY*unit {
			time.Sleep(2 * time.Second)
			g.yes = true
		}

		RectPos1 = append(RectPos1, rect.Bounds().Min.X+3000+50*unit, rect.Bounds().Max.X+4450+50*unit, rect.Bounds().Min.Y-1000+groundY*unit, rect.Bounds().Max.Y+500+groundY*unit)
	case 1:
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(200, 600)
		op.GeoM.Scale(1, 1)
		screen.DrawImage(rect, op)

		if rect.Bounds().Min.X+50*unit <= g.ghost.x-50*unit && g.ghost.x-50*unit < rect.Bounds().Max.X+1500+50*unit && rect.Bounds().Min.Y+groundY*unit <= g.ghost.y-groundY*unit && g.ghost.y-groundY*unit < rect.Bounds().Max.Y+1450+groundY*unit {
			time.Sleep(2 * time.Second)
			g.yes = true
		}

		RectPos2 = append(RectPos2, rect.Bounds().Min.X+50*unit, rect.Bounds().Max.X+1500+50*unit, rect.Bounds().Min.Y+groundY*unit, rect.Bounds().Max.Y+1450+groundY*unit)
	}

	msg1 := fmt.Sprintf("\n\n\n\n\n\n\n\n\n\nRect 1 Pos: %d", RectPos1[:])
	msg2 := fmt.Sprintf("\n\n\n\n\n\n\n\n\n\n\nRect 2 Pos: %d", RectPos2[:])

	msg := fmt.Sprintf("TPS: %0.2f\nDo NOT Touch To The Black Square.\nFirst Ghost X: %d\nFirst Ghost Y: %d\nSecond Ghost X: %d\nSecond Ghost Y: %d\nThird Ghost X: %d\nThird Ghost Y: %d\nFourth Ghost X: %d\nFourth Ghost Y: %d", ebiten.ActualTPS(), g.ghost.x-50*unit, g.ghost.y-groundY*unit, g.ghost.x/unit, g.ghost.y/unit-900, g.ghost.x/unit+1000, g.ghost.y/unit, g.ghost.x/unit+1000, g.ghost.y/unit-900)
	ebitenutil.DebugPrint(screen, msg)
	ebitenutil.DebugPrint(screen, msg1)
	ebitenutil.DebugPrint(screen, msg2)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 900, 800
}

// 8500, -550
func (g *Game) Update() error {
	if g.ghost == nil {
		g.ghost = &char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.ghost.vy = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.ghost.vy = 5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.ghost.vx = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.ghost.vx = 5 * unit
	}

	if g.yes {
		panic("err")
	}

	fmt.Println(g.ghost.x-50*unit, g.ghost.y-groundY*unit)
	g.ghost.update()
	return nil
}

func main() {
	ebiten.SetWindowSize(900, 800)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
