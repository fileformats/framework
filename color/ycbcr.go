package color

import (
	"fmt"
	"image/color"
)

type YCbCr struct {
	Y, Cb, Cr float64
}

func (c YCbCr) XYZ(cs *ColorSpace) XYZ {
	r, g, b := color.YCbCrToRGB(uint8(c.Y), uint8(c.Cb), uint8(c.Cr))
	return RGB{float64(r), float64(g), float64(b), 255.0}.XYZ(cs)
}

func (c YCbCr) RGBA() (r, g, b, a uint32) {
	rr, gg, bb := color.YCbCrToRGB(uint8(c.Y), uint8(c.Cb), uint8(c.Cr))
	return uint32(rr), uint32(gg), uint32(bb), 255
}

func (c YCbCr) String() string {
	return fmt.Sprintf("yCbCr(%.2f, %.2f, %.2f)", c.Y, c.Cb, c.Cr)
}
