package main

import (
	"github.com/aiju/gl"
	"github.com/asig/Go-SDL/sdl"
	"time"
	"math"
	"fmt"
 )

  
var timer, camx, camy, camz float64
var keys map[uint32] bool
var reload int
const border = 10

func keyboard(){
	if keys['w'] {
		if camz > -border {
			camz -= 0.1
		}
	}
	if keys['s']{
		if camz < border{
			camz += 0.1
		}
	}
	if keys['a']{
		if camx > -border{
			camx -= 0.1
		}
	}
	if keys['d']{
		if camx < border{
			camx += 0.1
		}
	}
	if keys[sdl.K_LSHIFT]{
		if camy < border{
			camy += 0.1
		}
	}
	if keys[sdl.K_LCTRL]{
		if camy > -border{
			camy -= 0.1
		}
	}
	if keys[sdl.K_SPACE] && reload == 0 {
		NewBullet(camx, camy-0.3, camz,0.3,0)
		reload = 10
	}else if(reload > 0){
		reload--
	}
}

var ring *Mesh

func draw() {
	modelview = gl.Translate(-camx, -camy, -camz)
	Clear(Color{0, 0, 0, 1})
	_ = math.Cos(0)
	//cube1.Render(gl.Mul4(gl.Translate(4*math.Cos(timer/50), 0, 4*math.Sin(timer/50)), gl.RotX(timer), gl.RotY(0)))
	//cube1.Render(gl.Mul4(gl.Translate(0, 4*math.Sin(timer/50), 4*math.Cos(timer/50)), gl.RotX(0), gl.RotY(timer)))
	//cube2.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(timer), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	//pyr1.Render(gl.Mul4(gl.Translate(6, 6, 6), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	//pyr1.Render(gl.Mul4(gl.Translate(-6, 6, -6), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	//pyr1.Render(gl.Mul4(gl.Translate(6, 6, -6), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	//pyr1.Render(gl.Mul4(gl.Translate(0, 2.2, 0), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))	
	//cyl.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(timer), gl.RotY(timer),gl.Scale(1.7,1.7,1.7)))
	//ring.Render(gl.RotX(timer))
	//affe.Render(gl.RotX(90))
	test.Render(gl.Mul4(gl.RotX(0),gl.RotY(timer)))
	//perfCyl.Render(gl.RotX(timer))
	DrawBullets()
	DrawShips()
	sdl.WM_SetCaption(fmt.Sprintf("x %.4f y %.4f z %.4f", camx , camy, camz), "")
}

func main() {
	gfxInit(800, 600)
	projection = gl.Frustum(45, 800./600, 0.01, 100)
	tick := time.Tick(time.Second / 50)
	NewShip(3,4,6,0)
	
	light = [3]float64{2, 0, 0}
	ambient = grey(0.1)
	diffuse = grey(0.5)
	specular = grey(0.5)

	modelsInit()
	ring = ReadModel("ring.ply", red)

	timer = 0.0
	camx = 0.0
	camy = 0.0
	camz = 8.0
	keys = make(map[uint32]bool)
	for {
		<- tick
		for _, ev := range sdl.PollEvents() {
			switch ev := ev.(type) {
			case sdl.QuitEvent:
				return
			case sdl.KeyboardEvent:
				switch ev.Type {
				case sdl.KEYDOWN:
					keys[ev.Keysym.Sym] = true
				case sdl.KEYUP:
					keys[ev.Keysym.Sym] = false
				}
			}
		}
		keyboard()
		light = [3]float64{camx, camy+3, camz}
		draw()
		MoveBullets()
		SpawnShips()
		MoveShips()
		ShipShoot()
		sdl.GL_SwapBuffers()
		timer++
	}
}
