package color

import "fmt"

type Gray struct {
	Y float64
}

func (c Gray) XYZ(cs *ColorSpace) XYZ {
	return RGB{c.Y, c.Y, c.Y, 255.0}.XYZ(cs)
}

func (c Gray) RGBA() (r, g, b, a uint32) {
	return uint32(c.Y), uint32(c.Y), uint32(c.Y), 255
}

func (c Gray) String() string {
	return fmt.Sprintf("gray(%.2f)", c.Y)
}