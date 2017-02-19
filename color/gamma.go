package color

type Gamma float64

const (
	Gamma_10   Gamma = 1.0
	Gamma_18   Gamma = 1.8
	Gamma_22   Gamma = 2.2
	Gamma_sRGB Gamma = -2.2
	Gamma_L    Gamma = 0.0
)

func (g Gamma) String() string {
	switch g {
	case Gamma_10:
		return "gamma(1.0)"
	case Gamma_18:
		return "gamma(1.8)"
	case Gamma_22:
		return "gamma(2.2)"
	case Gamma_sRGB:
		return "gamma(sRGB)"
	case Gamma_L:
		return "gamma(L)"
	}
	return "gamma(Unknown)"
}
