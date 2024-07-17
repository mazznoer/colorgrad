package colorgrad

import (
	"math"
)

type linearGradient struct {
	colors     [][4]float64
	pos        []float64
	dmin       float64
	dmax       float64
	mode       BlendMode
	firstColor Color
	lastColor  Color
}

func (lg linearGradient) At(t float64) Color {
	if t <= lg.dmin {
		return lg.firstColor
	}

	if t >= lg.dmax {
		return lg.lastColor
	}

	if math.IsNaN(t) {
		return Color{R: 0, G: 0, B: 0, A: 0}
	}

	low := 0
	high := len(lg.pos)

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
	t = (t - p1) / (p2 - p1)
	a, b, c, d := linearInterpolate(lg.colors[low-1], lg.colors[low], t)

	switch lg.mode {
	case BlendRgb:
		return Color{R: a, G: b, B: c, A: d}
	case BlendLinearRgb:
		return LinearRgb(a, b, c, d)
	case BlendOklab:
		return Oklab(a, b, c, d) //.Clamped()
	}

	return Color{R: 0, G: 0, B: 0, A: 0}
}

func newLinearGradient(colors []Color, pos []float64, mode BlendMode) Gradient {
	gradbase := linearGradient{
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
