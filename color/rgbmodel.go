package color

type RGBModel byte

const (
	AdobeRGB1998 RGBModel = iota
	AppleRGB
	BestRGB
	BetaRGB
	BruceRGB
	CIERGB
	ColorMatchRGB
	DonRGB4
	ECIRGBv2
	EktaSpacePS5
	NTSCRGB
	PALSECAMRGB
	ProPhotoRGB
	SMPTECRGB
	SRGB
	WideGamutRGB
)

var models = []string{
	"Adobe RGB (1998)",
	"Apple RGB",
	"Best RGB",
	"Beta RGB",
	"Bruce RGB",
	"CIER RGB",
	"Color Match RGB",
	"Don RGB 4",
	"ECI RGB v2",
	"Ekta Space PS5",
	"NTSC RGB",
	"PAL/SECAM RGB",
	"ProPhoto RGB",
	"SMPTE-C RGB",
	"sRGB",
	"Wide Gamut RGB",
}

func (m RGBModel) String() string {
	if int(m) < 0 || int(m) > len(models)-1 {
		return "Unknown"
	}
	return models[int(m)]
}

type computedRGBModel struct {
	RefWhite   RefWhite
	Gamma      float64
	GammaIndex int
	RGB2XYZ    [9]float64
	XYZ2RGB    [9]float64
}

func (mod RGBModel) getRGBModel() computedRGBModel {
	model := computedRGBModel{
		RefWhite: RefWhite{Y: 1.0},
	}

	var xr, yr, xg, yg, xb, yb float64

	switch mod {
	case AdobeRGB1998:
		xr = 0.64
		yr = 0.33
		xg = 0.21
		yg = 0.71
		xb = 0.15
		yb = 0.06
		model.RefWhite.X = 0.95047
		model.RefWhite.Z = 1.08883
		model.Gamma = 2.2
		model.GammaIndex = 2
	case AppleRGB:
		xr = 0.625
		yr = 0.340
		xg = 0.280
		yg = 0.595
		xb = 0.155
		yb = 0.070
		model.RefWhite.X = 0.95047
		model.RefWhite.Z = 1.08883
		model.Gamma = 1.8
		model.GammaIndex = 1
	case BestRGB:
		xr = 0.7347
		yr = 0.2653
		xg = 0.2150
		yg = 0.7750
		xb = 0.1300
		yb = 0.0350
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 2.2
		model.GammaIndex = 2
	case BetaRGB:
		xr = 0.6888
		yr = 0.3112
		xg = 0.1986
		yg = 0.7551
		xb = 0.1265
		yb = 0.0352
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 2.2
		model.GammaIndex = 2
	case BruceRGB:
		xr = 0.64
		yr = 0.33
		xg = 0.28
		yg = 0.65
		xb = 0.15
		yb = 0.06
		model.RefWhite.X = 0.95047
		model.RefWhite.Z = 1.08883
		model.Gamma = 2.2
		model.GammaIndex = 2
	case CIERGB:
		xr = 0.735
		yr = 0.265
		xg = 0.274
		yg = 0.717
		xb = 0.167
		yb = 0.009
		model.RefWhite.X = 1.00000
		model.RefWhite.Z = 1.00000
		model.Gamma = 2.2
		model.GammaIndex = 2
	case ColorMatchRGB:
		xr = 0.630
		yr = 0.340
		xg = 0.295
		yg = 0.605
		xb = 0.150
		yb = 0.075
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 1.8
		model.GammaIndex = 1
	case DonRGB4:
		xr = 0.696
		yr = 0.300
		xg = 0.215
		yg = 0.765
		xb = 0.130
		yb = 0.035
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 2.2
		model.GammaIndex = 2
	case ECIRGBv2:
		xr = 0.67
		yr = 0.33
		xg = 0.21
		yg = 0.71
		xb = 0.14
		yb = 0.08
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 0.0
		model.GammaIndex = 4
	case EktaSpacePS5:
		xr = 0.695
		yr = 0.305
		xg = 0.260
		yg = 0.700
		xb = 0.110
		yb = 0.005
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 2.2
		model.GammaIndex = 2
	case NTSCRGB:
		xr = 0.67
		yr = 0.33
		xg = 0.21
		yg = 0.71
		xb = 0.14
		yb = 0.08
		model.RefWhite.X = 0.98074
		model.RefWhite.Z = 1.18232
		model.Gamma = 2.2
		model.GammaIndex = 2
	case PALSECAMRGB:
		xr = 0.64
		yr = 0.33
		xg = 0.29
		yg = 0.60
		xb = 0.15
		yb = 0.06
		model.RefWhite.X = 0.95047
		model.RefWhite.Z = 1.08883
		model.Gamma = 2.2
		model.GammaIndex = 2
	case ProPhotoRGB:
		xr = 0.7347
		yr = 0.2653
		xg = 0.1596
		yg = 0.8404
		xb = 0.0366
		yb = 0.0001
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 1.8
		model.GammaIndex = 1
	case SMPTECRGB:
		xr = 0.630
		yr = 0.340
		xg = 0.310
		yg = 0.595
		xb = 0.155
		yb = 0.070
		model.RefWhite.X = 0.95047
		model.RefWhite.Z = 1.08883
		model.Gamma = 2.2
		model.GammaIndex = 2
	case SRGB:
		xr = 0.64
		yr = 0.33
		xg = 0.30
		yg = 0.60
		xb = 0.15
		yb = 0.06
		model.RefWhite.X = 0.95047
		model.RefWhite.Z = 1.08883
		model.Gamma = -2.2
		model.GammaIndex = 3
	case WideGamutRGB:
		xr = 0.735
		yr = 0.265
		xg = 0.115
		yg = 0.826
		xb = 0.157
		yb = 0.018
		model.RefWhite.X = 0.96422
		model.RefWhite.Z = 0.82521
		model.Gamma = 2.2
		model.GammaIndex = 2
	}

	m := [9]float64{
		xr / yr, xg / yg, xb / yb,
		1.0, 1.0, 1.0,
		(1.0 - xr - yr) / yr, (1.0 - xg - yg) / yg, (1.0 - xb - yb) / yb,
	}
	mi := invert(m)

	sr := model.RefWhite.X*mi[0] + model.RefWhite.Y*m[1] + model.RefWhite.Z*model[2]
	sg := model.RefWhite.X*mi[3] + model.RefWhite.Y*m[4] + model.RefWhite.Z*model[5]
	sb := model.RefWhite.X*mi[6] + model.RefWhite.Y*m[7] + model.RefWhite.Z*model[8]

	model.RGB2XYZ = [9]float64{
		sr * m[0], sg * m[1], sb * m[2],
		sr * m[3], sg * m[4], sb * m[5],
		sr * m[6], sg * m[7], sb * m[8],
	}
	transpose(model.RGB2XYZ)
	model.XYZ2RGB = invert(model.RGB2XYZ)

	return model
}
