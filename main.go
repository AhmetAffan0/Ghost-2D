package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	player1SpriteFull1 *ebiten.Image
	player1SpriteFull2 *ebiten.Image
	player1SpriteFull3 *ebiten.Image
	player1SpriteFull4 *ebiten.Image

	player1SpriteHalf1 *ebiten.Image
	player1SpriteHalf2 *ebiten.Image
	player1SpriteHalf3 *ebiten.Image
	player1SpriteHalf4 *ebiten.Image

	player1SpriteLow1 *ebiten.Image
	player1SpriteLow2 *ebiten.Image
	player1SpriteLow3 *ebiten.Image
	player1SpriteLow4 *ebiten.Image

	player2SpriteFull1 *ebiten.Image
	player2SpriteFull2 *ebiten.Image
	player2SpriteFull3 *ebiten.Image
	player2SpriteFull4 *ebiten.Image

	player2SpriteHalf1 *ebiten.Image
	player2SpriteHalf2 *ebiten.Image
	player2SpriteHalf3 *ebiten.Image
	player2SpriteHalf4 *ebiten.Image

	player2SpriteLow1 *ebiten.Image
	player2SpriteLow2 *ebiten.Image
	player2SpriteLow3 *ebiten.Image
	player2SpriteLow4 *ebiten.Image

	healBarFull *ebiten.Image
	healBarHalf *ebiten.Image
	healBarLow  *ebiten.Image

	ebitenBullet *ebiten.Image
)

func init() {
	rand.Seed(time.Now().UnixNano())

	img, _, err := ebitenutil.NewImageFromFile("Sprites/Sprite-0001.png")
	if err != nil {
		panic(err)
	}
	player1SpriteFull1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0002.png")
	if err != nil {
		panic(err)
	}
	player1SpriteFull2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0003.png")
	if err != nil {
		panic(err)
	}
	player1SpriteFull3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0004.png")
	if err != nil {
		panic(err)
	}
	player1SpriteFull4 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0006.png")
	if err != nil {
		panic(err)
	}
	player2SpriteFull1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0007.png")
	if err != nil {
		panic(err)
	}
	player2SpriteFull2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0008.png")
	if err != nil {
		panic(err)
	}
	player2SpriteFull3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0009.png")
	if err != nil {
		panic(err)
	}
	player2SpriteFull4 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0015.png")
	if err != nil {
		panic(err)
	}
	player1SpriteHalf1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0016.png")
	if err != nil {
		panic(err)
	}
	player1SpriteHalf2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0017.png")
	if err != nil {
		panic(err)
	}
	player1SpriteHalf3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0018.png")
	if err != nil {
		panic(err)
	}
	player1SpriteHalf4 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0019.png")
	if err != nil {
		panic(err)
	}
	player2SpriteHalf1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0020.png")
	if err != nil {
		panic(err)
	}
	player2SpriteHalf2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0021.png")
	if err != nil {
		panic(err)
	}
	player2SpriteHalf3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0022.png")
	if err != nil {
		panic(err)
	}
	player2SpriteHalf4 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0023.png")
	if err != nil {
		panic(err)
	}
	player1SpriteLow1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0024.png")
	if err != nil {
		panic(err)
	}
	player1SpriteLow2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0025.png")
	if err != nil {
		panic(err)
	}
	player1SpriteLow3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0026.png")
	if err != nil {
		panic(err)
	}
	player1SpriteLow4 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0027.png")
	if err != nil {
		panic(err)
	}
	player2SpriteLow1 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0028.png")
	if err != nil {
		panic(err)
	}
	player2SpriteLow2 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0029.png")
	if err != nil {
		panic(err)
	}
	player2SpriteLow3 = img

	img, _, err = ebitenutil.NewImageFromFile("Sprites/Sprite-0030.png")
	if err != nil {
		panic(err)
	}
	player2SpriteLow4 = img
}

type Pos struct {
	X int
	Y int
}

type player1 struct {
	x  int
	y  int
	vx int
	vy int
	s  *ebiten.Image
	b  *ebiten.Image
}

const (
	groundY = 250
	unit    = 10
)

func (c *player1) update() {
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

func (c *player1) draw(screen *ebiten.Image) {
	c.s = player1SpriteFull1
	c.b = ebitenBullet
	if c.vx > 0 {
		c.s = player1SpriteFull2
	} else if c.vx < 0 {
		c.s = player1SpriteFull3
	}
	if c.vy > 0 {
		c.s = player1SpriteFull1
	} else if c.vy < 0 {
		c.s = player1SpriteFull4
	}

	c.x += 12000
	c.x = c.x % 12000

	c.y += 7600
	c.y %= 7600

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
	opg.GeoM.Translate(float64(c.x)/unit+1300, float64(c.y)/unit)
	screen.DrawImage(c.s, opg)

	opc := &ebiten.DrawImageOptions{}
	opc.GeoM.Scale(2, 2)
	opc.GeoM.Translate(float64(c.x)/unit+1500, float64(c.y)/unit-900)
	screen.DrawImage(c.s, opc)

}

type player2 struct {
	x  int
	y  int
	vx int
	vy int
	s  *ebiten.Image
}

func (c *player2) update() {
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

func (c *player2) draw(screen *ebiten.Image) {
	c.s = player2SpriteFull1

	if c.vx > 0 {
		c.s = player2SpriteFull2
	} else if c.vx < 0 {
		c.s = player2SpriteFull3
	}
	if c.vy > 0 {
		c.s = player2SpriteFull1
	} else if c.vy < 0 {
		c.s = player2SpriteFull4
	}

	c.x += 12000
	c.x = c.x % 12000

	c.y += 7600
	c.y %= 7600

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
	opg.GeoM.Translate(float64(c.x)/unit+1300, float64(c.y)/unit)
	screen.DrawImage(c.s, opg)

	opc := &ebiten.DrawImageOptions{}
	opc.GeoM.Scale(2, 2)
	opc.GeoM.Translate(float64(c.x)/unit+1500, float64(c.y)/unit-900)
	screen.DrawImage(c.s, opc)

}

type Game struct {
	player1       *player1
	player2       *player2
	yes           bool
	bullet        []Pos
	curtain       []Pos
	moveDirection int
	moveCurtain   int
}

const (
	dirNone = iota
	dirUp
	dirDown
	dirLeft
	dirRight

	curtDown
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	g.player1.draw(screen)
	g.player2.draw(screen)

	vector.DrawFilledRect(screen, float32(g.curtain[0].X), float32(g.curtain[0].Y), 1300, 860, color.Black, false)

	opti := &ebiten.DrawImageOptions{}
	rect := ebiten.NewImage(10, 10)
	rect.Fill(color.Black)
	opti.GeoM.Translate(float64(g.bullet[0].X), float64(g.bullet[0].Y))
	opti.GeoM.Scale(2, 2)
	screen.DrawImage(rect, opti)

	if g.moveCurtain == dirNone {
		ebitenutil.DebugPrint(screen, "Press Enter For Start: ")
	} else {
		msg := fmt.Sprintf(`TPS: %.2f FPS: %0.2f	
Do NOT Touch To The Black Square.
Player1 First Ghost X: %d Y: %d
Player1 Second Ghost X: %d Y: %d
Player1 Third Ghost X: %d Y: %d
Player1 Fourth Ghost X: %d Y: %d
Player2 First Ghost X: %d Y: %d
Player2 Second Ghost X: %d Y: %d
Player2 Third Ghost X: %d Y: %d
Player2 Fourth Ghost X: %d Y: %d`,
			ebiten.ActualTPS(), ebiten.ActualFPS(),
			g.player1.x-50*unit, g.player1.y-groundY*unit,
			g.player1.x/unit, g.player1.y/unit-900,
			g.player1.x/unit+1000, g.player1.y/unit,
			g.player1.x/unit+1000, g.player1.y/unit-900,
			g.player2.x-1000*unit, g.player2.y-groundY*unit,
			g.player2.x/unit, g.player2.y/unit-900,
			g.player2.x/unit+1000, g.player2.y/unit,
			g.player2.x/unit+1000, g.player2.y/unit-900)
		ebitenutil.DebugPrint(screen, msg)
	}

}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 1300, 860
}

// 8500, -550
func (g *Game) Update() error {
	if g.player1 == nil {
		g.player1 = &player1{x: 50 * unit, y: groundY * unit}
	}

	if g.player2 == nil {
		g.player2 = &player2{x: 1000 * unit, y: groundY * unit}
	}

	switch g.moveDirection {
	case dirUp:
		g.bullet[0].Y -= 10
	case dirDown:
		g.bullet[0].Y += 10
	case dirRight:
		g.bullet[0].X += 10
	case dirLeft:
		g.bullet[0].X -= 10
	}

	switch g.moveCurtain {
	case curtDown:
		g.curtain[0].Y += 20
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player1.vy = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player1.vy = 5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player1.vx = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player1.vx = 5 * unit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		if g.player1.s == player1SpriteFull1 {
			g.moveDirection = dirDown
		}
		if g.player1.s == player1SpriteFull2 {
			g.moveDirection = dirRight
		}
		if g.player1.s == player1SpriteFull3 {
			g.moveDirection = dirLeft
		}
		if g.player1.s == player1SpriteFull4 {
			g.moveDirection = dirUp
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyShiftRight) {
		if g.player2.s == player2SpriteFull1 {
			g.moveDirection = dirDown
		}
		if g.player2.s == player2SpriteFull2 {
			g.moveDirection = dirRight
		}
		if g.player2.s == player2SpriteFull3 {
			g.moveDirection = dirLeft
		}
		if g.player2.s == player2SpriteFull4 {
			g.moveDirection = dirUp
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.moveCurtain = curtDown
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player2.vy = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player2.vy = 5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player2.vx = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player2.vx = 5 * unit
	}

	if g.yes {
		panic("err")
	}

	g.player1.update()
	g.player2.update()
	return nil
}

func newGame() *Game {
	g := &Game{
		bullet:  make([]Pos, 1),
		curtain: make([]Pos, 1),
	}

	return g
}

func main() {
	ebiten.SetWindowSize(1300, 860)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
