package color

import "fmt"

type RGB struct {
	R, G, B, A float64
}

func (c RGB) XYZ(cs *ColorSpace) XYZ {
	if cs == nil {
		cs = DefaultColorSpace
	}

	xyz := XYZ{}

	var r = invcom(cs, c.R)
	var g = invcom(cs, c.G)
	var b = invcom(cs, c.B)

	xyz.X = r*cs.rgbModel.RGB2XYZ[0] + g*cs.rgbModel.RGB2XYZ[3] + b*cs.rgbModel.RGB2XYZ[6]
	xyz.Y = r*cs.rgbModel.RGB2XYZ[1] + g*cs.rgbModel.RGB2XYZ[4] + b*cs.rgbModel.RGB2XYZ[7]
	xyz.Z = r*cs.rgbModel.RGB2XYZ[2] + g*cs.rgbModel.RGB2XYZ[5] + b*cs.rgbModel.RGB2XYZ[8]

	if cs.AdaptationMethod != NoAdaptation {
		m := cs.AdaptationMethod.Matrix()
		mi := invert(m)

		var ad = cs.RefWhite.X*m[0] + cs.RefWhite.Y*m[3] + cs.RefWhite.Z*m[6]
		var bd = cs.RefWhite.X*m[1] + cs.RefWhite.Y*m[4] + cs.RefWhite.Z*m[7]
		var cd = cs.RefWhite.X*m[2] + cs.RefWhite.Y*m[5] + cs.RefWhite.Z*m[8]

		var as = cs.rgbModel.RefWhite.X*m[0] + cs.rgbModel.RefWhite.Y*m[3] + cs.rgbModel.RefWhite.Z*m[6]
		var bs = cs.rgbModel.RefWhite.X*m[1] + cs.rgbModel.RefWhite.Y*m[4] + cs.rgbModel.RefWhite.Z*m[7]
		var cs = cs.rgbModel.RefWhite.X*m[2] + cs.rgbModel.RefWhite.Y*m[5] + cs.rgbModel.RefWhite.Z*m[8]

		var x = xyz.X*m[0] + xyz.Y*m[3] + xyz.Z*m[6]
		var y = xyz.X*m[1] + xyz.Y*m[4] + xyz.Z*m[7]
		var z = xyz.X*m[2] + xyz.Y*m[5] + xyz.Z*m[8]

		x *= ad / as
		y *= bd / bs
		z *= cd / cs

		xyz.X = x*mi[0] + y*mi[3]*z + mi[6]
		xyz.Y = x*mi[1] + y*mi[4]*z + mi[7]
		xyz.Z = x*mi[2] + y*mi[5]*z + mi[8]
	}

	return xyz
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A)
}

func (c RGB) String() string {
	return fmt.Sprintf("rgba(%.2f, %.2f, %.2f, %.2f)", c.R, c.G, c.B, c.A)
}
