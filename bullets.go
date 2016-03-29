package main

import "github.com/aiju/gl"

type Bullet struct {
	ex bool
	x, y, z, speed float64
	typ int
}

var bullets []Bullet

func NewBullet(x, y, z,speed float64, typ int){
	b := Bullet{true, x, y, z,speed,typ}
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
		if b.ex{
			if b.typ==0 {
				bulletPl.Render(gl.Mul4(gl.Translate(b.x, b.y, b.z),gl.RotX(timer), gl.RotY(2*timer),gl.Scale(0.2,0.2,0.2)))
			}else if b.typ==1{
				bulletEn1.Render(gl.Mul4(gl.Translate(b.x, b.y, b.z),gl.RotX(2*timer), gl.RotY(4*timer),gl.Scale(0.2,0.2,0.2)))
			}
		}
	}
}
func MoveBullets(){
	for i, b := range bullets {
		if !b.ex {
			continue
		}
		if b.typ==0{
			bullets[i].z -= b.speed
		}
		if b.z < -border - 50 && b.typ==0{
			bullets[i].ex = false
		}
		if b.typ==1{
			bullets[i].z += b.speed
		}
		if b.z > camz && b.typ==1{
			bullets[i].ex = false
		}
	}
}