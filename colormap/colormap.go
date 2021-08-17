package colormap

import "image/color"

type ColorMap func(color.Gray) color.Color
