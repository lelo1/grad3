package main

import (
	"github.com/aiju/gl"
	"github.com/neagix/Go-SDL/sdl"
	"time"
)

func main() {
	gfxInit(800, 600)
	tick := time.Tick(time.Second / 50)

	red := SolidColor(Color{1, 0, 0, 1})
	green := SolidColor(Color{0, 1, 0, 1})
	blue := SolidColor(Color{0, 0, 1, 1})
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
	))
	_ = blue
	_ = green

	modelview = gl.Mul4(gl.Frustum(45, 800./600, 0.01, 100), gl.Translate(0, 0, -8))

	time := 0.0
	for {
		select {
		case ev := <- sdl.Events:
			if _, ok := ev.(sdl.QuitEvent); ok {
				return
			}
		case <-tick:
			Clear(Color{0, 0, 0, 1})
			cube.Render(gl.Mul4(gl.RotX(time), gl.RotY(2*time)))
			sdl.GL_SwapBuffers()
			time++
		}
	}
}
