package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1024
	screenHeight = 512

	gridDim = 30
	PI      = math.Pi
	PI2     = PI / 2
	PI3     = 3 * PI / 2
	DR      = 0.0174533 //1 deg in rad
)

var (
	p player

	gridX = 8
	gridY = 8
	gridS = 64
	grid  = []int{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}
)

type Game struct {
}

type player struct {
	x, y, width, height float64
	pdx, pdy, ang       float64 //player delta x/y and angle
}

func init() {
	p = player{x: screenWidth/2 - screenWidth/4, y: screenHeight / 2, width: 10, height: 10,
		ang: 0}
	p.pdx = math.Cos(p.ang) * 2 //player dist
	p.pdy = math.Sin(p.ang) * 2
}

func degToRad(a int) float64 {
	return float64(a) * PI / 180.0
}

func (g *Game) Update() error {
	g.handleMovement()
	return nil
}

func (g *Game) handleMovement() {
	xo := 0 //offset
	yo := 0
	if p.pdx < 0 {
		xo = -20
	} else {
		xo = 20
	}
	if p.pdy < 0 {
		yo = -20
	} else {
		yo = 20
	}

	//rotate the player
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.ang -= 0.05

		if p.ang < 0 {
			p.ang += 2 * PI
		}
		p.pdx = math.Cos(p.ang) * 2
		p.pdy = math.Sin(p.ang) * 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.ang += 0.05

		if p.ang > 2*PI {
			p.ang -= 2 * PI
		}
		p.pdx = math.Cos(p.ang) * 2
		p.pdy = math.Sin(p.ang) * 2
	}

	//collision wall
	var ipx int = int(p.x / 64)
	var ipx_add_xo int = int(p.x+float64(xo)) / 64
	var ipx_sub_xo int = int(p.x-float64(xo)) / 64

	var ipy int = int(p.y / 64)
	var ipy_add_yo int = int(p.y+float64(yo)) / 64
	var ipy_sub_yo int = int(p.y-float64(yo)) / 64

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if grid[int(ipy*(gridX)+ipx_add_xo)] == 0 {
			p.x += p.pdx
		}
		if grid[int(ipy_add_yo*(gridX)+ipx)] == 0 {
			p.y += p.pdy
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if grid[int(ipy*(gridX)+ipx_sub_xo)] == 0 {
			p.x -= p.pdx
		}
		if grid[int(ipy_sub_yo*(gridX)+ipx)] == 0 {
			p.y -= p.pdy
		}
	}

}

func (g *Game) Draw(screen *ebiten.Image) {

	//floor
	ebitenutil.DrawRect(screen, screenWidth/2, screenHeight/3,
		screenWidth/2, screenHeight/2+screenHeight/3,
		color.RGBA{0, 100, 255, 255})

	//ceiling
	ebitenutil.DrawRect(screen, screenWidth/2, 0,
		screenWidth/2, screenHeight/3,
		color.RGBA{100, 100, 100, 255})

	drawMap(screen)

	drawRays(screen)

	ebitenutil.DrawRect(screen, float64(p.x-p.width/2), float64(p.y-p.height/2), float64(p.width), float64(p.height), color.RGBA{255, 0, 0, 255}) //player
	ebitenutil.DrawLine(screen, p.x, p.y, p.x+p.pdx*5, p.y+p.pdy*5, color.RGBA{R: 0, G: 255, B: 255, A: 255})                                     //player looking

}

func drawMap(screen *ebiten.Image) {
	var x, y int
	var xo, yo float64

	//create grid 64*64 is one square
	for y = 0; y < gridY; y++ {
		for x = 0; x < gridX; x++ {
			xo = float64(x) * float64(gridS)
			yo = float64(y) * float64(gridS)
			if grid[y*gridX+x] == 1 {
				ebitenutil.DrawRect(screen, float64(xo), float64(yo), float64(gridS), float64(gridS), color.White)
			} else {
				ebitenutil.DrawRect(screen, float64(xo), float64(yo), float64(gridS), float64(gridS), color.RGBA{20, 20, 20, 255})
			}
		}
	}
}

func drawRays(screen *ebiten.Image) {
	var mx, my, mp, dof int
	var rx, ry, ra, xo, yo, disT float64
	var shade float64

	ra = p.ang - DR*30 //back 30 deg to initial view
	if ra < 0 {
		ra += 2 * PI
	}
	if ra > 2*PI {
		ra -= 2 * PI
	}

	dof = 0

	for r := 0; r < 60; r++ {

		//---Horizontal hit line---
		dof = 0
		var disH float64 = 10000
		hx := p.x
		hy := p.y
		aTan := -1 / math.Tan(ra)
		if ra > PI { //looking up
			ry = float64((int(p.y)>>6)<<6) - 0.0001
			rx = (p.y-ry)*aTan + p.x
			yo = -64
			xo = -yo * aTan
		}
		if ra < PI {
			ry = float64((int(p.y)>>6)<<6) + 64
			rx = (p.y-ry)*aTan + p.x
			yo = 64
			xo = -yo * aTan
		}
		if ra == 0 || ra == PI {
			rx = p.x
			ry = p.y
			dof = 8
		}

		for dof < 8 {
			mx = int(rx) >> 6
			my = int(ry) >> 6
			mp = my*gridX + mx
			if mp > 0 && mp < gridX*gridY && grid[mp] > 0 { //hit wall
				hx = rx
				hy = ry
				disH = dist(p.x, p.y, hx, hy, ra)
				dof = 8
			} else {
				rx += xo
				ry += yo
				dof += 1
			}
		}

		//---vertical hit line---
		dof = 0
		var disV float64 = 10000
		vx := p.x
		vy := p.y
		nTan := -math.Tan(ra)
		if ra > PI2 && ra < PI3 {
			rx = float64((int(p.x)>>6)<<6) - 0.0001
			ry = (p.x-rx)*nTan + p.y
			xo = -64
			yo = -xo * nTan
		}
		if ra < PI2 || ra > PI3 {
			rx = float64((int(p.x)>>6)<<6) + 64
			ry = (p.x-rx)*nTan + p.y
			xo = 64
			yo = -xo * nTan
		}
		if ra == 0 || ra == PI {
			rx = p.x
			ry = p.y
			dof = 8
		}

		for dof < 8 {
			mx = int(rx) >> 6
			my = int(ry) >> 6
			mp = my*gridX + mx
			if mp > 0 && mp < gridX*gridY && grid[mp] > 0 { //hit wall
				vx = rx
				vy = ry
				disV = dist(p.x, p.y, vx, vy, ra)
				dof = 8
			} else {
				rx += xo
				ry += yo
				dof += 1
			}
		}
		shade = 1
		if disV < disH { //ray hit square side
			rx = vx
			ry = vy
			disT = disV

			shade = 0.7
		}
		if disH < disV {
			rx = hx
			ry = hy
			disT = disH

		}

		ebitenutil.DrawLine(screen, p.x, p.y, rx, ry, color.RGBA{R: 0, G: 255, B: 0, A: 255}) //draw ray

		//draw 3d scene
		draw3D(screen, disT, r, ra, rx, ry, shade)

		//next angle
		ra += DR
		//limit
		if ra < 0 {
			ra += 2 * PI
		}
		if ra > 2*PI {
			ra -= 2 * PI
		}
	}
}

func draw3D(screen *ebiten.Image, disT float64, r int, ra float64, rx float64, ry float64, shade float64) {

	ca := p.ang - ra
	if ca < 0 {
		ca += 2 * PI
	}
	if ca > 2*PI {
		ca -= 2 * PI
	}
	disT = disT * math.Cos(ca) //fix fisheye

	lineH := float64((gridS * 320)) / (disT) //altezza dei muri

	ty_step := 32 / lineH
	var ty_off float64

	if lineH > 320 {
		ty_off = (lineH - 320) / 2
		lineH = 320
	}

	lineO := 160 - lineH/2 //line offset

	//draw wall
	var ty float64 = ty_off * ty_step
	var tx int

	//apply texture for each side (horizontal / verical) and mirror it
	if shade == 1 {
		tx = int(rx/2) % 32
		if ra < degToRad(180) {
			tx = 31 - tx
		}
	} else {
		tx = int(ry/2) % 32
		if ra > degToRad(90) && ra < degToRad(270) {
			tx = 31 - tx
		}
	}

	for y := 0; y < int(lineH); y++ { //create 1 rectangle of 1 px for the entire hight
		c := int(float64(all_texture[int(ty)*32+tx]) * 255 * shade) //index of texture array (return 0 or 1) + dark shade
		ebitenutil.DrawRect(screen, float64(r*8+screenWidth/2),
			float64(y)+lineO,
			8, 1,
			color.RGBA{uint8(c), uint8(c), uint8(c), 255},
		)

		ty += ty_step
	}

}

func dist(ax float64, ay float64, bx float64, by float64, ang float64) float64 {
	return (math.Sqrt((bx-ax)*(bx-ax) + (by-ay)*(by-ay)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("raycast")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
