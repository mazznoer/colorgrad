package colorgrad

import (
	"math"
)

type linearGradient struct {
	colors    [][4]float64
	positions []float64
	min       float64
	max       float64
	mode      BlendMode
	first     Color
	last      Color
}

func (lg linearGradient) At(t float64) Color {
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
	t = (t - p1) / (p2 - p1)
	a, b, c, d := linearInterpolate(lg.colors[low-1], lg.colors[low], t)

	switch lg.mode {
	case BlendRgb:
		return Color{R: a, G: b, B: c, A: d}
	case BlendLinearRgb:
		return LinearRgb(a, b, c, d)
	case BlendLab:
		return Lab(a, b, c, d).Clamp()
	case BlendOklab:
		return Oklab(a, b, c, d).Clamp()
	}

	return Color{}
}

func newLinearGradient(colors []Color, positions []float64, mode BlendMode) Gradient {
	gradbase := linearGradient{
		colors:    convertColors(colors, mode),
		positions: positions,
		min:       positions[0],
		max:       positions[len(positions)-1],
		mode:      mode,
		first:     colors[0],
		last:      colors[len(colors)-1],
	}

	return Gradient{
		Core: gradbase,
		Min:  positions[0],
		Max:  positions[len(positions)-1],
	}
}
