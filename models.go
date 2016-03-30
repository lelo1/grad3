package main

import (
	"math"
 )


var cube1, cube2,bulletPl, pyr1,modShip1,bulletEn1 *Object
var cyl,affe,perfCyl *Mesh

var red, green,blue,yellow *Material 

func modelsInit() {
	red = SolidColor(Color{1, 0, 0, 1})
	green = SolidColor(Color{0, 1, 0, 1})
	blue = SolidColor(Color{0, 0, 1, 1})
	yellow = SolidColor(Color{1, 1, 0, 1})

	cube2 = cube(yellow)
 	bulletPl = cube(blue)
 	bulletEn1 = cube(red)
	cyl = cylinder(yellow)
	affe = ReadModel("Affe.ply",blue)
	perfCyl = ReadModel("perfCyl.ply",green)
	spezModel()
}

//Modelle



//Funktionen
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
		1, 1, 1,
		-1, 1, 1,

		-1, -1, 1,
		1, -1, 1,
		1, 1, 1,
		
		-1, -1, -1,
		-1, 1, 1,
		-1, 1, -1,

		-1, -1, -1,
		-1 , -1, 1,
		-1 ,1 ,1,	

		1, -1, -1,
		1, 1, -1,
		1, 1, 1,

		1, -1, -1,
		1 ,1 ,1,
		1 , -1, 1,	
																																						
		-1, 1, -1,
		1, 1, 1,
		1, 1, -1,

		-1, 1, -1,
		-1, 1, 1,
		1, 1, 1,

		-1, -1, -1,
		1, -1, -1,
		1, -1, 1,

		-1, -1, -1,
		1, -1, 1,
		-1, -1, 1,
	))
}
func pyr(m *Material) *Object {
	return NewObject(NewMesh(m,
		-1, -1, -1,
		1, -1, -1,
		1, -1, 1,
		
		-1, -1, -1,
		-1, -1, 1,
		1, -1, 1,
		
	), NewMesh(m,
		-1,-1,-1,
		0 ,1 ,0,
		-1, -1, 1,
	), NewMesh(m,
		-1, -1, -1,
		0, 1, 0,
		1, -1, -1,
	), NewMesh(m,
		1, -1, 1,
		1, -1, -1,
		0, 1, 0,
	),NewMesh(m,
		-1, -1, 1,
		1, -1, 1,
		0, 1, 0,
	))
}
	
//  Spezial
func spezModel(){
	cube1 = NewObject(NewMesh(red,
		-1, -1, -1,
		-1, 1, -1,
		1, 1, -1,

		-1, -1, -1,
		1, 1, -1,
		1, -1, -1,

		-1, -1, 1,
		1, 1, 1,
		-1, 1, 1,

		-1, -1, 1,
		1, -1, 1,
		1, 1, 1,
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
	modShip1 = NewObject(NewMesh(red,
		-1, -1, -1,
		1, -1, -1,
		1, -1, 1,
		
		-1, -1, -1,
		-1, -1, 1,
		1, -1, 1,
		
	), NewMesh(blue,
		-1,-1,-1,
		-1, -1, 1,
		0 ,1 ,0,
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
}




 