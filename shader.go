package main

import (
	"github.com/aiju/gl"
)

var vertexShader = `
#version 110
attribute vec3 position;
attribute vec3 normal;
uniform mat4 matrix;
uniform mat4 projection;
varying vec3 positiont, normalt;
void main() {
	positiont = vec3(matrix * vec4(position, 1));
	normalt = vec3(matrix * vec4(normal, 0));
	gl_Position = projection * vec4(positiont, 1);
}
`

var fragmentShader = `
#version 110
varying vec3 positiont, normalt;
uniform vec3 light;
uniform vec3 specular, diffuse, ambient;
uniform vec4 color;
uniform float specexp;
void main() {
	vec3 k, a, d, s, n, nn, r;
	float x;

	k = color.xyz;
	a = ambient * k;
	
	n = normalize(light - positiont);
	nn = normalize(normalt);
	d = max(dot(n, nn), 0.0) * diffuse * k;
	
	r = reflect(-n, nn);
	x = max(dot(r, normalize(-positiont)), 0.0);
	s = pow(x, specexp) * specular * k;
	gl_FragColor = vec4(a + d + s, color.w);
}
`

var defaultShader *gl.Program

func shaderInit() {
	defaultShader = NewShader(vertexShader, fragmentShader)
}
