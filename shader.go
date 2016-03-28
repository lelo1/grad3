package main

import (
	"github.com/aiju/gl"
)

var vertexShader = `
#version 110
attribute vec3 position;
attribute vec3 normal;
uniform mat4 matrix;
void main() {
	gl_Position = matrix * vec4(position.xyz, 1);
}
`

var fragmentShader = `
#version 110
uniform vec4 color;
void main() {
	gl_FragColor = color;
}
`

var defaultShader *gl.Program

func shaderInit() {
	defaultShader = NewShader(vertexShader, fragmentShader)
}
