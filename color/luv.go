package color

import (
	"fmt"
	"math"
)

type LUV struct {
	L, U, V float64
}

func (c LUV) XYZ(cs *ColorSpace) XYZ {
	xyz := XYZ{}
	if c.L > kKE {
		xyz.Y = math.Pow((c.L+16.0)/116.0, 3.0)
	} else {
		xyz.Y = c.L / kK
	}
	var u0 = (4.0 * cs.RefWhite.X) / (cs.RefWhite.X + 15.0*cs.RefWhite.Y + 3.0*cs.RefWhite.Z)
	var v0 = (9.0 * cs.RefWhite.Y) / (cs.RefWhite.X + 15.0*cs.RefWhite.Y + 3.0*cs.RefWhite.Z)
	var a = (((52.0 * c.L) / (c.U + 13.0*c.L*u0)) - 1.0) / 3.0
	var b = -5.0 * xyz.Y
	var d = xyz.Y * (((39.0 * c.L) / (c.V + 13.0*c.L*v0)) - 5.0)

	xyz.X = (d - b) / (a - (-1.0 / 3.0))
	xyz.Z = xyz.X*a + b

	return xyz
}

func (c LUV) RGBA() (r, g, b, a uint32) {
	return c.XYZ(DefaultColorSpace).RGBA()
}

func (c LUV) String() string {
	return fmt.Sprintf("luv(%.2f, %.2f, %.2f)", c.L, c.U, c.V)
}
