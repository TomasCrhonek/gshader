package main

import (
	"fmt"

	"github.com/TomasCrhonek/gshader/colormap"
	"github.com/TomasCrhonek/gshader/shader"
)

func listShadersColormaps() {
	fmt.Printf("First ones are default.\n\n")
	fmt.Printf("Shaders  : %v\n", shaderList)
	fmt.Printf("ColorMaps: %v\n", colorList)
}

type shaderInfo struct {
	name   string
	shader shader.Shader
}

func (s shaderInfo) String() string {
	return s.name
}

var shaderList []shaderInfo

func registerShader(s shader.Shader, n string) {
	shaderList = append(shaderList, shaderInfo{name: n, shader: s})
}

func getShader(n string) shader.Shader {
	for _, s := range shaderList {
		if n == s.name {
			return s.shader
		}
	}

	return nil
}

type colorInfo struct {
	name  string
	color colormap.ColorMap
}

func (c colorInfo) String() string {
	return c.name
}

var colorList []colorInfo

func registerColormap(c colormap.ColorMap, n string) {
	colorList = append(colorList, colorInfo{name: n, color: c})
}

func getColormap(n string) colormap.ColorMap {
	for _, c := range colorList {
		if n == c.name {
			return c.color
		}
	}

	return nil
}
