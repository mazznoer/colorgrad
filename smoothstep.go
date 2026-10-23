package colorgrad

import (
	"math"
	"slices"
)

type SmoothstepGradient struct {
	colors      [][4]float64
	positions   []float64
	min, max    float64
	mode        BlendMode
	first, last Color
}

func (sg *SmoothstepGradient) At(t float64) Color {
	if math.IsNaN(t) {
		return Color{A: 1}
	}

	if t <= sg.min {
		return sg.first
	}

	if t >= sg.max {
		return sg.last
	}

	lo, hi := 1, len(sg.positions)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if sg.positions[mid] < t {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	p1 := sg.positions[lo-1]
	p2 := sg.positions[lo]
	t = (t - p1) / (p2 - p1)
	a, b, c, d := smoothstepInterpolate(sg.colors[lo-1], sg.colors[lo], t)

	switch sg.mode {
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

// colors & positions guaranteed to have the same length (minimum 2)
func newSmoothstepGradient(colors []Color, positions []float64, mode BlendMode) Gradient {
	pos := slices.Clone(positions)

	g := &SmoothstepGradient{
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

func smoothstepInterpolate(a, b [4]float64, t float64) (i, j, k, l float64) {
	i = (b[0]-a[0])*(3.0-t*2.0)*t*t + a[0]
	j = (b[1]-a[1])*(3.0-t*2.0)*t*t + a[1]
	k = (b[2]-a[2])*(3.0-t*2.0)*t*t + a[2]
	l = (b[3]-a[3])*(3.0-t*2.0)*t*t + a[3]
	return
}
