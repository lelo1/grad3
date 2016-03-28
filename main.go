package main

import (
	"github.com/aiju/gl"
	"github.com/neagix/Go-SDL/sdl"
	"time"
	"math"
	"fmt"
 )

func cylinder(m *Material) *Mesh {
	l := []float64{}
	s := 2*math.Pi/64
	for a := 0.0; a < 2*math.Pi; a += s {
		l = append(l,
			0, 0, 0,
			0, math.Cos(a), math.Sin(a),
			0, math.Cos(a+s), math.Sin(a+s))
	}
	return NewMesh(m, l...)
}
func cube(m *Material) *Object {
	return NewObject(NewMesh(m,
		-1, -1, -1,
		-1, 1, -1,
		1, 1, -1,

		-1, -1, -1,
		1, 1, -1,
		1, -1, -1,

		-1, -1, 1,
		-1, 1, 1,
		1, 1, 1,

		-1, -1, 1,
		1, 1, 1,
		1, -1, 1,
		
		-1, -1, -1,
		-1, 1, -1,
		-1, 1, 1,

		-1, -1, -1,
		-1 , -1, 1,
		-1 ,1 ,1,	

		1, -1, -1,
		1, 1, -1,
		1, 1, 1,

		1, -1, -1,
		1 , -1, 1,
		1 ,1 ,1,	
																																						
		-1, 1, -1,
		1, 1, -1,
		1, 1, 1,

		-1, 1, -1,
		-1, 1, 1,
		1, 1, 1,

		-1, -1, -1,
		1, -1, -1,
		1, -1, 1,

		-1, -1, -1,
		-1, -1, 1,
		1, -1, 1,
	))
}

var timer, camx, camy, camz float64
var keys map[uint32] bool
var cube1, cube2,bullet, pyr1 *Object
var cyl *Mesh
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
		NewBullet(camx, camy-0.3, camz,0.3)
		reload = 10
	}else if(reload > 0){
		reload--
	}
}

func draw() {
	modelview = gl.Mul4(gl.Frustum(45, 800./600, 0.01, 100), gl.Translate(-camx, -camy, -camz))
	Clear(Color{0, 0, 0, 1})
	cube1.Render(gl.Mul4(gl.Translate(4*math.Cos(timer/50), 0, 4*math.Sin(timer/50)), gl.RotX(timer), gl.RotY(0)))
	cube1.Render(gl.Mul4(gl.Translate(0, 4*math.Sin(timer/50), 4*math.Cos(timer/50)), gl.RotX(0), gl.RotY(timer)))
	cube2.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(timer), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	pyr1.Render(gl.Mul4(gl.Translate(6, 6, 6), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	pyr1.Render(gl.Mul4(gl.Translate(-6, 6, -6), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	pyr1.Render(gl.Mul4(gl.Translate(6, 6, -6), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))
	pyr1.Render(gl.Mul4(gl.Translate(0, 2.2, 0), gl.RotX(0), gl.RotY(timer), gl.Scale(0.5,0.5,0.5)))	
	cyl.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(timer), gl.RotY(timer),gl.Scale(1.7,1.7,1.7)))
	DrawBullets()
	sdl.WM_SetCaption(fmt.Sprintf("x %.4f y %.4f z %.4f", camx , camy, camz), "")
}

func main() {
	gfxInit(800, 600)
	tick := time.Tick(time.Second / 50)

	red := SolidColor(Color{1, 0, 0, 1})
	green := SolidColor(Color{0, 1, 0, 1})
	blue := SolidColor(Color{0, 0, 1, 1})
	yellow := SolidColor(Color{1, 1, 0, 1})
	cube1 = NewObject(NewMesh(red,
		-1, -1, -1,
		-1, 1, -1,
		1, 1, -1,

		-1, -1, -1,
		1, 1, -1,
		1, -1, -1,

		-1, -1, 1,
		-1, 1, 1,
		1, 1, 1,

		-1, -1, 1,
		1, 1, 1,
		1, -1, 1,
	), NewMesh(green,
		-1, -1, -1,
		-1, 1, -1,
		-1, 1, 1,

		-1, -1, -1,
		-1 , -1, 1,
		-1 ,1 ,1,	

		1, -1, -1,
		1, 1, -1,
		1, 1, 1,

		1, -1, -1,
		1 , -1, 1,
		1 ,1 ,1,	
																																						
	), NewMesh(blue,
		-1, 1, -1,
		1, 1, -1,
		1, 1, 1,

		-1, 1, -1,
		-1, 1, 1,
		1, 1, 1,

		-1, -1, -1,
		1, -1, -1,
		1, -1, 1,

		-1, -1, -1,
		-1, -1, 1,
		1, -1, 1,
	))
	pyr1 = NewObject(NewMesh(red,
		-1, -1, -1,
		1, -1, -1,
		1, -1, 1,
		
		-1, -1, -1,
		-1, -1, 1,
		1, -1, 1,
		
	), NewMesh(blue,
		-1,-1,-1,
		0 ,1 ,0,
		-1, -1, 1,
	), NewMesh(yellow,
		-1, -1, -1,
		0, 1, 0,
		1, -1, -1,
	), NewMesh(green,
		1, -1, 1,
		1, -1, -1,
		0, 1, 0,
	),NewMesh(yellow,
		-1, -1, 1,
		1, -1, 1,
		0, 1, 0,
	))
	//))
	cube2 = cube(yellow)
 	bullet = cube(blue)
 	cyl = cylinder(yellow)
	


	timer = 0.0
	camx = 0.0
	camy = 0.0
	camz = 8.0
	keys = make(map[uint32]bool)
	for { 
		select {
		case ev := <- sdl.Events:
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
		case <-tick:
			keyboard()
			draw()
			MoveBullets()
			sdl.GL_SwapBuffers()
			timer++
		}
	}
}
