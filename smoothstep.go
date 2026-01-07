package colorgrad

import (
	"math"
)

type smoothstepGradient struct {
	colors    [][4]float64
	positions []float64
	min       float64
	max       float64
	mode      BlendMode
	first     Color
	last      Color
}

func (sg smoothstepGradient) At(t float64) Color {
	if t <= sg.min {
		return sg.first
	}

	if t >= sg.max {
		return sg.last
	}

	if math.IsNaN(t) {
		return Color{A: 1}
	}

	low := 0
	high := len(sg.positions)

	for low < high {
		mid := (low + high) / 2
		if sg.positions[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low == 0 {
		low = 1
	}

	p1 := sg.positions[low-1]
	p2 := sg.positions[low]
	t = (t - p1) / (p2 - p1)
	a, b, c, d := smoothstepInterpolate(sg.colors[low-1], sg.colors[low], t)

	switch sg.mode {
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

func newSmoothstepGradient(colors []Color, positions []float64, mode BlendMode) Gradient {
	gradbase := smoothstepGradient{
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

func smoothstepInterpolate(a, b [4]float64, t float64) (i, j, k, l float64) {
	i = (b[0]-a[0])*(3.0-t*2.0)*t*t + a[0]
	j = (b[1]-a[1])*(3.0-t*2.0)*t*t + a[1]
	k = (b[2]-a[2])*(3.0-t*2.0)*t*t + a[2]
	l = (b[3]-a[3])*(3.0-t*2.0)*t*t + a[3]
	return
}
