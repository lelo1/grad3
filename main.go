package main

import (
	"github.com/aiju/gl"
	"github.com/neagix/Go-SDL/sdl"
	"time"
	"math"
)

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
	



	time := 0.0
	for {
		select {
		case ev := <- sdl.Events:
			if _, ok := ev.(sdl.QuitEvent); ok {
				return
			}
		case <-tick:
			modelview = gl.Mul4(gl.Frustum(45, 800./600, 0.01, 100), gl.Translate(0, 0, -8))
			Clear(Color{0, 0, 0, 1})
			_ = math.Sin(0)
			cube.Render(gl.Mul4(gl.Translate(4*math.Cos(time/50), 0, 4*math.Sin(time/50)), gl.RotX(time), gl.RotY(0)))
			cube.Render(gl.Mul4(gl.Translate(0, 4*math.Sin(time/50), 4*math.Cos(time/50)), gl.RotX(0), gl.RotY(time)))
			cube2.Render(gl.Mul4(gl.Translate(0, 0, 0), gl.RotX(time), gl.RotY(time), gl.Scale(0.5,0.5,0.5)))
			sdl.GL_SwapBuffers()
			time++
		}
	}
}
