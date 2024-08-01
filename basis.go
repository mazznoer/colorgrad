package colorgrad

import (
	"math"
)

// https://github.com/d3/d3-interpolate/blob/master/src/basis.js

type basisGradient struct {
	colors    [][4]float64
	positions []float64
	min       float64
	max       float64
	mode      BlendMode
	first     Color
	last      Color
}

func (lg basisGradient) At(t float64) Color {
	if t <= lg.min {
		return lg.first
	}

	if t >= lg.max {
		return lg.last
	}

	if math.IsNaN(t) {
		return Color{A: 1}
	}

	low := 0
	high := len(lg.positions)
	n := high - 1

	for low < high {
		mid := (low + high) / 2
		if lg.positions[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low == 0 {
		low = 1
	}

	p1 := lg.positions[low-1]
	p2 := lg.positions[low]
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

	a := xx(val0[0], val1[0], 0)
	b := xx(val0[1], val1[1], 1)
	c := xx(val0[2], val1[2], 2)
	d := xx(val0[3], val1[3], 3)

	switch lg.mode {
	case BlendRgb:
		return Color{R: a, G: b, B: c, A: d}
	case BlendLinearRgb:
		return LinearRgb(a, b, c, d)
	case BlendOklab:
		return Oklab(a, b, c, d).Clamp()
	}

	return Color{}
}

func newBasisGradient(colors []Color, positions []float64, mode BlendMode) Gradient {
	gradbase := basisGradient{
		colors:    convertColors(colors, mode),
		positions: positions,
		min:       positions[0],
		max:       positions[len(positions)-1],
		mode:      mode,
		first:     colors[0],
		last:      colors[len(colors)-1],
	}

	return Gradient{
		grad: gradbase,
		dmin: positions[0],
		dmax: positions[len(positions)-1],
	}
}

func basis(t1, v0, v1, v2, v3 float64) float64 {
	t2 := t1 * t1
	t3 := t2 * t1
	return ((1-3*t1+3*t2-t3)*v0 + (4-6*t2+3*t3)*v1 + (1+3*t1+3*t2-3*t3)*v2 + t3*v3) / 6
}
