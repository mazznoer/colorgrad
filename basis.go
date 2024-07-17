//go:build ignore
package colorgrad

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// https://github.com/d3/d3-interpolate/blob/master/src/basis.js

type basisGradient struct {
	colors     [][3]float64
	pos        []float64
	dmin       float64
	dmax       float64
	mode       BlendMode
	firstColor colorful.Color
	lastColor  colorful.Color
}

func (lg basisGradient) At(t float64) colorful.Color {
	if t <= lg.dmin {
		return lg.firstColor
	}

	if t >= lg.dmax {
		return lg.lastColor
	}

	if math.IsNaN(t) {
		return colorful.Color{R: 0, G: 0, B: 0}
	}

	low := 0
	high := len(lg.pos)
	n := high - 1

	for low < high {
		mid := (low + high) / 2
		if lg.pos[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low == 0 {
		low = 1
	}

	p1 := lg.pos[low-1]
	p2 := lg.pos[low]
	val0 := lg.colors[low-1]
	val1 := lg.colors[low]
	i := low - 1
	t = (t - p1) / (p2 - p1)

	xx := func(v1, v2 float64, j int) float64 {
		v0 := 2*v1 - v2
		if i > 0 {
			v0 = lg.colors[i-1][j]
		}

		v3 := 2*v2 - v1
		if i < n-1 {
			v3 = lg.colors[i+2][j]
		}

		return basis(t, v0, v1, v2, v3)
	}

	x := xx(val0[0], val1[0], 0)
	y := xx(val0[1], val1[1], 1)
	z := xx(val0[2], val1[2], 2)

	switch lg.mode {
	case BlendHcl:
		hue := interpAngle(lg.colors[low-1][0], lg.colors[low][0], t)
		return colorful.Hcl(hue, y, z).Clamped()
	case BlendHsv:
		hue := interpAngle(lg.colors[low-1][0], lg.colors[low][0], t)
		return colorful.Hsv(hue, y, z)
	case BlendLab:
		return colorful.Lab(x, y, z).Clamped()
	case BlendLinearRgb:
		return colorful.LinearRgb(x, y, z)
	case BlendLuv:
		return colorful.Luv(x, y, z).Clamped()
	case BlendRgb:
		return colorful.Color{R: x, G: y, B: z}
	case BlendOklab:
		a, b, c := oklabToLrgb(x, y, z)
		return colorful.LinearRgb(a, b, c).Clamped()
	}

	return colorful.Color{R: 0, G: 0, B: 0}
}

func newBasisGradient(colors []colorful.Color, pos []float64, mode BlendMode) Gradient {
	gradbase := basisGradient{
		colors:     convertColors(colors, mode),
		pos:        pos,
		dmin:       pos[0],
		dmax:       pos[len(pos)-1],
		mode:       mode,
		firstColor: colors[0],
		lastColor:  colors[len(colors)-1],
	}

	return Gradient{
		grad: gradbase,
		dmin: pos[0],
		dmax: pos[len(pos)-1],
	}
}

func basis(t1, v0, v1, v2, v3 float64) float64 {
	t2 := t1 * t1
	t3 := t2 * t1
	return ((1-3*t1+3*t2-t3)*v0 + (4-6*t2+3*t3)*v1 + (1+3*t1+3*t2-3*t3)*v2 + t3*v3) / 6
}
