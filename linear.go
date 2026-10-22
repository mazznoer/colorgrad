package colorgrad

import (
	"iter"
	"math"
	"slices"
)

type LinearGradient struct {
	colors      [][4]float64
	positions   []float64
	min, max    float64
	mode        BlendMode
	first, last Color
}

func (lg *LinearGradient) At(t float64) Color {
	if math.IsNaN(t) {
		return Color{A: 1}
	}

	if t <= lg.min {
		return lg.first
	}

	if t >= lg.max {
		return lg.last
	}

	lo, hi := 0, len(lg.positions)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if lg.positions[mid] < t {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	p1 := lg.positions[lo-1]
	p2 := lg.positions[lo]
	t = (t - p1) / (p2 - p1)
	a, b, c, d := linearInterpolate(lg.colors[lo-1], lg.colors[lo], t)

	switch lg.mode {
	case BlendLinearRgb:
		return LinearRgb(a, b, c, d)
	case BlendLab:
		return Lab(a, b, c, d).Clamp()
	case BlendOklab:
		return Oklab(a, b, c, d).Clamp()
	default:
		return Color{R: a, G: b, B: c, A: d}
	}
}

func (lg *LinearGradient) Mode() BlendMode {
	return lg.mode
}

func (lg *LinearGradient) Stops() iter.Seq2[float64, [4]float64] {
	return func(yield func(float64, [4]float64) bool) {
		for i, t := range lg.positions {
			if !yield(t, lg.colors[i]) {
				return
			}
		}
	}
}

func newLinearGradient(colors []Color, positions []float64, mode BlendMode) Gradient {
	pos := slices.Clone(positions)

	g := &LinearGradient{
		colors:    convertColors(colors, mode),
		positions: pos,
		min:       pos[0],
		max:       pos[len(pos)-1],
		mode:      mode,
		first:     colors[0],
		last:      colors[len(pos)-1],
	}

	return Gradient{
		Core: g,
		Min:  g.min,
		Max:  g.max,
	}
}
