package color

type Adaptation byte

const (
	Bradford Adaptation = iota
	VonKries
	XYZScaling
	NoAdaptation
)

func (a Adaptation) Matrix() [9]float64 {
	matrix := [9]float64{}
	switch a {
	case Bradford:
		matrix = [9]float64{
			0.8951, -0.7502, 0.0389,
			0.2664, 1.7135, -0.0685,
			-0.1614, 0.0367, 1.0296,
		}
	case VonKries:
		matrix = [9]float64{
			0.40024, -0.22630, 0.0,
			0.70760, 1.16532, 0.0,
			-0.08081, 0.04570, 0.91822,
		}
	default:
		matrix = [9]float64{
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
		}
	}
	return matrix
}

func (a Adaptation) String() string {
	switch a {
	case Bradford:
		return "Bradford"
	case VonKries:
		return "Von Kries"
	case XYZScaling:
		return "XYZ Scaling"
	case NoAdaptation:
		return "NoAdaptation"
	}
	return "Unknown"
}
