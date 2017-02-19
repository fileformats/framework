package color

import "image/color"

type csColor interface {
	XYZ(cs *ColorSpace) XYZ
}

type ColorSpace struct {
	RefWhite         RefWhite
	RGBModel         RGBModel
	AdaptationMethod Adaptation
	Gamma            Gamma

	rgbModel computedRGBModel
}

var DefaultColorSpace = NewColorSpace(D50, SRGB, Bradford, Gamma_sRGB)

func NewColorSpace(refWhite RefWhite, rgbModel RGBModel, adaptation Adaptation, gamma Gamma) *ColorSpace {
	var cs = &ColorSpace{
		RefWhite:         refWhite,
		RGBModel:         rgbModel,
		AdaptationMethod: adaptation,
		Gamma:            gamma,

		rgbModel: rgbModel.getRGBModel(),
	}
	return cs
}

func (cs *ColorSpace) SetRGBModel(model RGBModel) {
	cs.RGBModel = model
	cs.rgbModel = model.getRGBModel()
}

func (cs *ColorSpace) ToCMYK(c color.Color) CMYK {
	rgb := cs.ToRGB(c)
	cc, m, y, k := color.RGBToCMYK(uint8(rgb.R), uint8(rgb.G), uint8(rgb.B))
	return CMYK{float64(cc), float64(m), float64(y), float64(k)}
}

func (cs *ColorSpace) ToGray(c color.Color) Gray {
	if csc, ok := c.(csColor); ok {
		c = csc.XYZ(cs).rgb(cs)
	}
	r, g, b, _ := c.RGBA()
	return Gray{(float64(r) + float64(g) + float64(b)) / 3}
}

func (cs *ColorSpace) ToLAB(c color.Color) LAB {
	if csc, ok := c.(csColor); ok {
		return csc.XYZ(cs).lab(cs)
	}
	r, g, b, a := c.RGBA()
	rgb := RGB{float64(r), float64(g), float64(b), float64(a)}
	return rgb.XYZ(cs).lab(cs)
}

func (cs *ColorSpace) ToLUV(c color.Color) LUV {
	if csc, ok := c.(csColor); ok {
		return csc.XYZ(cs).luv(cs)
	} else {
		r, g, b, a := c.RGBA()
		rgb := RGB{float64(r), float64(g), float64(b), float64(a)}
		return rgb.XYZ(cs).luv(cs)
	}
}

func (cs *ColorSpace) ToRGB(c color.Color) RGB {
	if csc, ok := c.(csColor); ok {
		return csc.XYZ(cs).rgb(cs)
	}
	r, g, b, a := c.RGBA()
	return RGB{float64(r), float64(g), float64(b), float64(a)}
}

func (cs *ColorSpace) ToXYZ(c color.Color) XYZ {
	if csc, ok := c.(csColor); ok {
		return csc.XYZ(cs)
	}
	r, g, b, a := c.RGBA()
	return RGB{float64(r), float64(g), float64(b), float64(a)}.XYZ(cs)
}

func (cs *ColorSpace) ToYCbCr(c color.Color) YCbCr {
	rgb := cs.ToRGB(c)
	y, cb, cr := color.RGBToYCbCr(uint8(rgb.R), uint8(rgb.G), uint8(rgb.B))
	return YCbCr{float64(y), float64(cb), float64(cr)}
}
