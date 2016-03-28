package main

import (
	"github.com/aiju/gl"
	"github.com/neagix/Go-SDL/sdl"
	"fmt"
	"os"
	"math"
)

type Color struct {
	R, G, B, A float64
}

type Material struct {
	shader *gl.Program
	color Color
}

type Mesh struct {
	vert []float64
	normals []float64
	buf *gl.Buffer
	mat *Material
}

type Object struct {
	mesh []*Mesh
}

var modelview gl.Mat4

func gfxInit(w, h int) {
	sdl.Init(sdl.INIT_VIDEO)
	sdl.SetVideoMode(w, h, 32, sdl.OPENGL|sdl.DOUBLEBUF|sdl.HWSURFACE)
	gl.Init()
	gl.Enable(gl.DEPTH_TEST)
	gl.Viewport(0, 0, w, h)
	shaderInit()
}

func NewShader(a, b string) *gl.Program {
	p, err := gl.MakeProgram([]string{a}, []string{b})
	if err != nil {
		fmt.Println(err)
		sdl.Quit()
		os.Exit(0)
	}
	return p
}

func SolidColor(col Color) *Material {
	return &Material{
		shader: defaultShader,
		color: col,
	}
}

func normal(v []float64) []float64 {
	var a, b, c [3]float64
	a[0], a[1], a[2] = v[0] - v[3], v[1] - v[4], v[2] - v[5]
	b[0], b[1], b[2] = v[0] - v[6], v[1] - v[7], v[2] - v[8]
	c[0] = a[1] * b[2] - a[2] * b[1]
	c[1] = a[2] * b[0] - a[0] * b[2]
	c[2] = a[0] * b[1] - a[1] * b[0]
	l := c[0] * c[0] + c[1] * c[1] + c[2] * c[2]
	if l != 0 {
		l = math.Pow(l, -0.5)
		c[0], c[1], c[2] = c[0] * l, c[1] * l, c[2] * l
	}
	return c[:]
}

func interleave(a []float64, na int, b []float64, nb int) []float64 {
	if len(a) / na != len(b) / nb || len(a) % na != 0 || len(b) % nb != 0 {
		panic("interleave: slices of invalid length")
	}
	n := len(a) / na
	nr := na + nb
	r := make([]float64, n * nr)
	for i := 0; i < n; i++ {
		copy(r[i*nr:i*nr+na], a[i*na:(i+1)*na])
		copy(r[i*nr+na:(i+1)*nr], b[i*nb:(i+1)*nb])
	}
	return r
}

func NewMesh(mat *Material, vert ...float64) *Mesh {
	if len(vert) % 9 != 0 {
		panic("NewMesh: number of vertices not divisible by 9")
	}
	n := make([]float64, len(vert))
	for i := 0; i < len(vert); i += 9 {
		no := normal(vert[i:i+9])
		copy(n[i:i+3], no)
		copy(n[i+3:i+6], no)
		copy(n[i+6:i+9], no)
	}
	arr := interleave(vert, 3, n, 3)
	fmt.Println(len(arr))
	return &Mesh{
		vert: vert,
		normals: n,
		buf: gl.NewBuffer(gl.ARRAY_BUFFER, arr, gl.STATIC_DRAW),
		mat: mat,
	}
}

func NewObject(mesh ...*Mesh) *Object {
	return &Object{mesh}
}

func (m *Mesh) Render(mat gl.Mat4) {
	s := m.mat.shader
	s.Use()
	s.EnableAttrib("position", m.buf, 0, 3, 6, false)
	s.EnableAttrib("normal", m.buf, 3, 3, 6, false)
	s.SetUniform("matrix", gl.Mul4(modelview, mat))
	c := m.mat.color
	s.SetUniform("color", [4]float64{c.R, c.G, c.B, c.A})
	gl.DrawArrays(gl.TRIANGLES, 0, len(m.vert)/3)
	s.DisableAttrib("position")
	s.DisableAttrib("normal")
	s.Unuse()
}

func (o *Object) Render(mat gl.Mat4) {
	for _, a := range o.mesh {
		a.Render(mat)
	}
}

func Clear(col Color){
	gl.ClearColor(col.R, col.G, col.B, col.A)
	gl.Clear(gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT)
}
