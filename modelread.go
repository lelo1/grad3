package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func modelline(f *bufio.Scanner) []string {
	if !f.Scan() {
		return []string{}
	}
	return strings.Split(f.Text(), " ")
}

func ReadModel(filename string, mat *Material) *Mesh {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	f := bufio.NewScanner(fi)
	
	l := modelline(f)
	if l[0] != "ply" {
		panic(filename + " not ply format")
	}
	
	nvertex := 0
	nface := 0
	for {
		l = modelline(f)
		if l[0] == "end_header" {
			break
		}
		if l[0] == "element" && l[1] == "vertex"{
			nvertex, _ = strconv.Atoi(l[2])
		}
		if l[0] == "element" && l[1] == "face"{
			nface, _ = strconv.Atoi(l[2])
		}
	}
	vertices := make([][3]float64, nvertex)
	faces := make([][3]int, nface)
	normals := make([][3]float64, nvertex)
	for i := 0; i < nvertex; i++ {
		l = modelline(f)
		vertices[i][0],_ = strconv.ParseFloat(l[0], 64)
		vertices[i][1],_ = strconv.ParseFloat(l[1], 64)
		vertices[i][2],_ = strconv.ParseFloat(l[2], 64)
		normals[i][0],_ = strconv.ParseFloat(l[3], 64)
		normals[i][1],_ = strconv.ParseFloat(l[4], 64)
		normals[i][2],_ = strconv.ParseFloat(l[5], 64)
		
	}
	for i := 0;i < nface; i++{
		l = modelline(f)
		if l[0] == "3" {
			faces[i][0],_ = strconv.Atoi(l[1])
			faces[i][1],_ = strconv.Atoi(l[2])
			faces[i][2],_ = strconv.Atoi(l[3])
		}else {
			panic("Not a Triangle")
		}
	}  
	points := make([]float64, nface*9)
	pointsNormals := make([]float64, nface*9)
	for i := 0; i < nface;i++{
		for j := 0; j < 3; j++{
			for k:= 0; k < 3; k++{			
				points [i * 9 + 3 * j + k] = vertices[faces[i][j]][k]
				pointsNormals [i * 9 + 3 * j + k] = normals[faces[i][j]][k]
			}

}		 
	}
	_ = fmt.Print
	return NewMeshNorm(mat, points,pointsNormals)
}