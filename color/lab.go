package color

import (
	"fmt"
	"math"
)

type LAB struct {
	L, A, B float64
}

func (c LAB) XYZ(cs *ColorSpace) XYZ {
	var fy = (c.L + 16.0) / 116.0
	var fx = 0.002*c.A + fy
	var fz = fy - 0.005*c.B

	var fx3 = fx * fx * fx
	var fz3 = fz * fz * fz

	var xr, yr, zr float64

	if fx > kE {
		xr = fx3
	} else {
		xr = (116.0*fx - 16.0) / kK
	}

	if c.L > kKE {
		yr = math.Pow((c.L+16.0)/116.0, 3.0)
	} else {
		yr = c.L / kK
	}
	if fz3 > kE {
		zr = fz3
	} else {
		zr = (116.0*fz - 16.0) / kK
	}

	return XYZ{
		X: xr * cs.RefWhite.X,
		Y: yr * cs.RefWhite.Y,
		Z: zr * cs.RefWhite.Z,
	}
}

func (c LAB) RGBA() (r, g, b, a uint32) {
	return c.XYZ(DefaultColorSpace).RGBA()
}

func (c LAB) String() string {
	return fmt.Sprintf("lab(%.2f, %.2f, %.2f)", c.L, c.A, c.B)
}
