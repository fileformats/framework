package color

import (
	"fmt"
	"image/color"
)

type CMYK struct {
	C, M, Y, K float64
}

func (c CMYK) XYZ(cs *ColorSpace) XYZ {
	r, g, b := color.CMYKToRGB(uint8(c.C), uint8(c.M), uint8(c.Y), uint8(c.K))
	return RGB{float64(r), float64(g), float64(b), 255.0}.XYZ(cs)
}

func (c CMYK) RGBA() (r, g, b, a uint32) {
	rr, gg, bb := color.CMYKToRGB(uint8(c.C), uint8(c.M), uint8(c.Y), uint8(c.K))
	return uint32(rr), uint32(gg), uint32(bb), 255
}

func (c CMYK) String() string {
	return fmt.Sprintf("hsl(%.2f, %.2f, %.2f, %.2f)", c.C, c.M, c.Y, c.K)
}