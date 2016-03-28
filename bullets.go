package main

import "github.com/aiju/gl"

type Bullet struct {
	ex bool
	x, y, z, speed float64
}

var bullets []Bullet

func NewBullet(x, y, z,speed float64){
	b := Bullet{true, x, y, z,speed}
	for i := range bullets {
		if !bullets[i].ex {
			bullets[i] = b
			return
		}
	}
	bullets = append(bullets, b)
}

func DrawBullets() {
	for _, b := range bullets {
		if b.ex {
			bullet.Render(gl.Mul4(gl.Translate(b.x, b.y, b.z),gl.RotX(timer), gl.RotY(2*timer),gl.Scale(0.2,0.2,0.2)))
		}
	}
}
func MoveBullets(){
	for i, b := range bullets {
		if b.ex {
			bullets[i].z -= b.speed
		}
		if b.z < -border - 50 {
			bullets[i].ex = false
		}
	}
}