package color

import (
	"fmt"
	"image/color"
	"math"
)

type XYZ struct {
	X, Y, Z float64
}

func (c XYZ) XYZ(cs *ColorSpace) XYZ {
	return XYZ{c.X, c.Y, c.Z}
}

func (c XYZ) gray(cs *ColorSpace) Gray {
	rgb := c.rgb(cs)
	return Gray{(rgb.R + rgb.G + rgb.B) / 3}
}

func (c XYZ) ycbcr(cs *ColorSpace) YCbCr {
	r, g, b, _ := c.rgb(cs).RGBA()
	y, cb, cr := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))
	return YCbCr{float64(y), float64(cb), float64(cr)}
}

func (c XYZ) luv(cs *ColorSpace) LUV {
	var den = c.X + 15.0 + c.Y + 3*c.Z
	var up, vp, urp, vrp, l float64
	if den > 0 {
		up = (4.0 * c.X) / (c.X + 15.0*c.Y + 3.0*c.Z)
		vp = (9.0 * c.Y) / (c.X + 15.0*c.Y + 3.0*c.Z)
	}
	urp = (4.0 * cs.RefWhite.X) / (cs.RefWhite.X + 15.0*cs.RefWhite.Y + 3.0*cs.RefWhite.Z)
	vrp = (9.0 * cs.RefWhite.Y) / (cs.RefWhite.X + 15.0*cs.RefWhite.Y + 3.0*cs.RefWhite.Z)

	yr := c.Y / cs.RefWhite.Y
	if yr > kE {
		l = 116.0*math.Pow(yr, 1.0/3.0) - 16.0
	} else {
		l = kK * yr
	}
	return LUV{
		l,
		13.0 * l * (up - urp),
		13.0 * l * (vp - vrp),
	}
}

func (c XYZ) lab(cs *ColorSpace) LAB {
	var xr = c.X / cs.RefWhite.X
	var yr = c.Y / cs.RefWhite.Y
	var zr = c.Z / cs.RefWhite.Z
	var fx, fy, fz float64

	if xr > kE {
		fx = math.Pow(xr, 1.0/3.0)
	} else {
		fx = (kK*xr + 16.0) / 116.0
	}

	if yr > kE {
		fy = math.Pow(yr, 1.0/3.0)
	} else {
		fy = (kK*yr + 16.0) / 116.0
	}

	if zr > kE {
		fz = math.Pow(zr, 1.0/3.0)
	} else {
		fz = (kK*zr + 16.0) / 116.0
	}

	return LAB{
		L: 116.0*fy - 16.0,
		A: 500.0 * (fx - fy),
		B: 200.0 * (fy - fz),
	}
}

func (c XYZ) rgb(cs *ColorSpace) RGB {
	var x2 = c.X
	var y2 = c.Y
	var z2 = c.Z

	rgb := RGB{}

	if cs.AdaptationMethod != NoAdaptation {

		m := cs.AdaptationMethod.Matrix()
		mi := invert(m)

		var As = cs.RefWhite.X*m[0] + cs.RefWhite.Y*m[3] + cs.RefWhite.Z*m[6]
		var Bs = cs.RefWhite.X*m[1] + cs.RefWhite.Y*m[4] + cs.RefWhite.Z*m[7]
		var Cs = cs.RefWhite.X*m[2] + cs.RefWhite.Y*m[5] + cs.RefWhite.Z*m[8]

		var Ad = cs.rgbModel.RefWhite.X*m[0] + cs.rgbModel.RefWhite.Y*m[3] + cs.rgbModel.RefWhite.Z*m[6]
		var Bd = cs.rgbModel.RefWhite.X*m[1] + cs.rgbModel.RefWhite.Y*m[4] + cs.rgbModel.RefWhite.Z*m[7]
		var Cd = cs.rgbModel.RefWhite.X*m[2] + cs.rgbModel.RefWhite.Y*m[5] + cs.rgbModel.RefWhite.Z*m[8]

		var x1 = c.X*m[0] + c.Y*m[3] + c.Z*m[6]
		var y1 = c.X*m[1] + c.Y*m[4] + c.Z*m[7]
		var z1 = c.X*m[2] + c.Y*m[5] + c.Z*m[8]

		x1 *= Ad / As
		y1 *= Bd / Bs
		z1 *= Cd / Cs

		x2 = x1*mi[0] + y1*mi[3] + z1*mi[6]
		y2 = x1*mi[1] + y1*mi[4] + z1*mi[7]
		z2 = x1*mi[2] + y1*mi[5] + z1*mi[8]
	}

	rgb.R = com(cs, x2*cs.rgbModel.XYZ2RGB[0]+y2*cs.rgbModel.XYZ2RGB[3]+z2*cs.rgbModel.XYZ2RGB[6])
	rgb.G = com(cs, x2*cs.rgbModel.XYZ2RGB[1]+y2*cs.rgbModel.XYZ2RGB[4]+z2*cs.rgbModel.XYZ2RGB[7])
	rgb.B = com(cs, x2*cs.rgbModel.XYZ2RGB[2]+y2*cs.rgbModel.XYZ2RGB[5]+z2*cs.rgbModel.XYZ2RGB[8])
	rgb.A = 255.0

	return rgb
}

func (c XYZ) RGBA() (r, g, b, a uint32) {
	return c.rgb(DefaultColorSpace).RGBA()
}

func (c XYZ) String() string {
	return fmt.Sprintf("xyz(%.2f, %.2f, %.2f)", c.X, c.Y, c.Z)
}
