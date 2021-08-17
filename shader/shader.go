package shader

import "image/color"

type Shader func(int, int, int, int) color.Color
