package main

import (
	"github.com/aiju/gl"
	"github.com/neagix/Go-SDL/sdl"
	"time"
	"math"
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

func main() {
	gfxInit(800, 600)
	tick := time.Tick(time.Second / 50)

	red := SolidColor(Color{1, 0, 0, 1})
	green := SolidColor(Color{0, 1, 0, 1})
	blue := SolidColor(Color{0, 0, 1, 1})
	yellow := SolidColor(Color{1, 1, 0, 1})
	cube := NewObject(NewMesh(red,
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
	cube2 := NewObject(NewMesh(yellow,
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
 	cyl := cylinder(yellow)
	_  = cube2



	time := 0.0
	camx := 0.0
	camy := 0.0
	camz := 8.0
	keys := make(map[uint32]bool)
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
			if keys['w'] {
				camz -= 0.1
			}
			if keys['s']{
				camz += 0.1
			}
			if keys['a']{
				camx -= 0.1
			}
			if keys['d']{
				camx += 0.1
			}
			if keys[sdl.K_LSHIFT]{
				camy += 0.1
			}
			if keys[sdl.K_LCTRL]{
				camy -= 0.1
			}
		
			modelview = gl.Mul4(gl.Frustum(45, 800./600, 0.01, 100), gl.Translate(-camx, -camy, -camz))
			Clear(Color{0, 0, 0, 1})
			_ = math.Sin(0)
			cube.Render(gl.Mul4(gl.Translate(4*math.Cos(time/50), 0, 4*math.Sin(time/50)), gl.RotX(time), gl.RotY(0)))
			cube.Render(gl.Mul4(gl.Translate(0, 4*math.Sin(time/50), 4*math.Cos(time/50)), gl.RotX(0), gl.RotY(time)))
			//cube2.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(time), gl.RotY(time), gl.Scale(0.5,0.5,0.5)))
			cyl.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(time), gl.RotY(time),gl.Scale(2,2,2)))
			
			sdl.GL_SwapBuffers()
			time++
		}
	}
}
