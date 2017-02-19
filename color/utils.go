package color

import "math"

var (
	kE  = 216.0 / 24389.0
	kK  = 24389.0 / 27.0
	kKE = 8.0
)

//    0 1 2
//  0 0 1 2
//  1 3 4 5
//  2 6 7 8

func determinant(m [9]float64) float64 {
	return m[0]*(m[8]*m[4]-m[7]*m[5]) - m[3]*(m[8]*m[1]-m[7]*m[2]) + m[6]*(m[5]*m[1]-m[4]*m[2])
}

func invert(m [9]float64) [9]float64 {
	var scale = 1.0 / determinant(m)
	return [9]float64{
		+scale * (m[8]*m[4] - m[7]*m[5]),
		-scale * (m[8]*m[1] - m[7]*m[2]),
		+scale * (m[5]*m[1] - m[4]*m[2]),

		-scale * (m[8]*m[3] - m[6]*m[5]),
		+scale * (m[8]*m[0] - m[6]*m[2]),
		-scale * (m[5]*m[0] - m[3]*m[2]),

		+scale * (m[7]*m[3] - m[6]*m[4]),
		-scale * (m[7]*m[0] - m[6]*m[1]),
		+scale * (m[4]*m[0] - m[3]*m[1]),
	}
}

func transpose(m [9]float64) {
	m[1], m[3] = m[3], m[1]
	m[2], m[6] = m[6], m[2]
	m[5], m[7] = m[7], m[5]
}

func com(cs *ColorSpace, l float64) (c float64) {
	if cs.Gamma > 0 {
		if l >= 0 {
			c = math.Pow(l, 1.0/float64(cs.Gamma))
		} else {
			c = -math.Pow(-l, 1.0/float64(cs.Gamma))
		}
	} else if cs.Gamma < 0 {
		s := 1.0
		if l < 0 {
			s = -1.0
			l = -l
		}
		if l <= 0.0031308 {
			c = l * 12.92
		} else {
			c = 1.055*math.Pow(l, 1.0/2.4) - 0.055
		}
		c *= s
	} else {
		s := 1.0
		if l < 0 {
			s = -1.0
			l = -l
		}
		if l <= 216.0/24389.0 {
			c = l * 24389.0 / 2700.0
		} else {
			c = 1.16*math.Pow(l, 1.0/3.0) - 0.16
		}
		c *= s
	}
	return
}

func invcom(cs *ColorSpace, c float64) (l float64) {
	if cs.Gamma > 0 {
		if c > 0 {
			l = math.Pow(c, float64(cs.Gamma))
		} else {
			l = -math.Pow(-c, float64(cs.Gamma))
		}
	} else if cs.rgbModel.Gamma < 0 {
		s := 1.0
		if c < 0 {
			s = -1.0
			c = -c
		}
		if c <= 0.04045 {
			l = c / 12.92
		} else {
			l = math.Pow((c+0.055)/1.055, 2.4)
		}
		l *= s
	} else {
		s := 1.0
		if c < 0 {
			s = -1.0
			c = -c
		}
		if c <= 0.08 {
			l = 2700.0 * c / 24389.0
		} else {
			l = (((1000000.0*c+480000.0)*c+76800.0)*c + 4096.0) / 1560896.0
		}
		l *= s
	}
	return
}
