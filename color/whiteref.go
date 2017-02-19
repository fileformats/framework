package color

type RefWhite struct {
	X, Y, Z float64
}

var (
	A   = RefWhite{1.09850, 1.0, 0.35585}
	B   = RefWhite{0.99072, 1.0, 0.85223}
	C   = RefWhite{0.98074, 1.0, 1.18232}
	D50 = RefWhite{0.96422, 1.0, 0.82521}
	D55 = RefWhite{0.95682, 1.0, 0.92149}
	D65 = RefWhite{0.95047, 1.0, 1.08883}
	D75 = RefWhite{0.94972, 1.0, 1.22638}
	E   = RefWhite{1.0, 1.0, 1.0}
	F2  = RefWhite{0.99186, 1.0, 0.67393}
	F7  = RefWhite{0.95041, 1.0, 1.08747}
	F11 = RefWhite{1.00962, 1.0, 0.64350}
)
