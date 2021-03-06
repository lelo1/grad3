package main

import "github.com/aiju/gl"
import "math/rand"

type Ship struct {
	ex bool
	x, y, z, speed float64
}

var ships []Ship

func NewShip(x, y, z,speed float64){
	b := Ship{true, x, y, z,speed}
	for i := range ships {
		if !ships[i].ex {
			ships[i] = b
			return
		}
	}
	ships = append(ships, b)
}

func DrawShips() {
	for _, b := range ships {
		if b.ex {
			modShip1.Render(gl.Mul4(gl.Translate(b.x, b.y, b.z),gl.RotX(0),gl.RotZ(90), gl.RotY(-90),gl.Scale(0.5,0.5,0.5)))
		}
	}
}
func MoveShips(){
	for i, b := range ships {
		if b.ex {
			ships[i].z += b.speed
			if b.z > camz + 10 {
				ships[i].ex = false
			}
		}
	}
}
func SpawnShips(){
	rnd := rand.Intn(100)
	if(rnd< 10){
		randX:= rand.Float64() * 20 - 10
		randY:= rand.Float64() * 20 - 10
		NewShip(randX,randY,-30,0.2)
		
	}
	if rnd==99{
		NewShip(camx,camy,-30,0.2)
	}
}
func ShipShoot(){
	for i, b :=range ships{
		if b.ex{
			rnd :=rand.Intn(100)
			if rnd < 1{
				bx := ships[i].x
				by := ships[i].y
				bz := ships[i].z
				NewBullet(bx+0.5,by,bz,0.4,1)	
				NewBullet(bx-0.5,by,bz,0.4,1)
			}
		}
	}	
}





