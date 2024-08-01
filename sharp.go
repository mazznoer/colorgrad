package colorgrad

import (
	"math"
)

type sharpGradient struct {
	colors    []Color
	positions []float64
	last      int
	min       float64
	max       float64
}

func (sg sharpGradient) At(t float64) Color {
	if t <= sg.min {
		return sg.colors[0]
	}

	if t >= sg.max {
		return sg.colors[sg.last]
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

	i := low - 1
	p1 := sg.positions[i]
	p2 := sg.positions[low]

	if i%2 == 0 {
		return sg.colors[i]
	}

	t = (t - p1) / (p2 - p1)
	a := sg.colors[i]
	b := sg.colors[low]
	return blendRgb(a, b, t)
}

func newSharpGradient(colorsIn []Color, dmin, dmax float64, smoothness float64) Gradient {
	n := len(colorsIn)
	colors := make([]Color, n*2)
	i := 0
	for _, c := range colorsIn {
		colors[i] = c
		i++
		colors[i] = c
		i++
	}
	t := clamp01(smoothness) * (dmax - dmin) / float64(n) / 4
	p := linspace(dmin, dmax, uint(n+1))
	positions := make([]float64, n*2)
	i = 0
	j := 0
	for x := 0; x < int(n); x++ {
		positions[i] = p[j]
		if i > 0 {
			positions[i] += t
		}
		i++
		j++
		positions[i] = p[j]
		if i < len(colors)-1 {
			positions[i] -= t
		}
		i++
	}
	gradbase := sharpGradient{
		colors:    colors,
		positions: positions,
		last:      int(n*2 - 1),
		min:       dmin,
		max:       dmax,
	}
	return Gradient{
		grad: gradbase,
		dmin: dmin,
		dmax: dmax,
	}
}
