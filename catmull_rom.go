package colorgrad

import (
	"math"
)

// Adapted from https://qroph.github.io/2018/07/30/smooth-paths-using-catmull-rom-splines.html

func toCatmullRomSegments(values []float64) [][4]float64 {
	alpha := 0.5
	tension := 0.0
	n := len(values)

	vals := make([]float64, n+2)
	vals[0] = 2*values[0] - values[1]
	for i, v := range values {
		vals[i+1] = v
	}
	vals[n+1] = 2*values[n-1] - values[n-2]

	segments := [][4]float64{}

	for i := 1; i < len(vals)-2; i++ {
		v0 := vals[i-1]
		v1 := vals[i]
		v2 := vals[i+1]
		v3 := vals[i+2]
		t0 := 0.0
		t1 := t0 + math.Pow(math.Abs(v0-v1), alpha)
		t2 := t1 + math.Pow(math.Abs(v1-v2), alpha)
		t3 := t2 + math.Pow(math.Abs(v2-v3), alpha)
		m1 := (1. - tension) * (t2 - t1) * ((v0-v1)/(t0-t1) - (v0-v2)/(t0-t2) + (v1-v2)/(t1-t2))
		m2 := (1. - tension) * (t2 - t1) * ((v1-v2)/(t1-t2) - (v1-v3)/(t1-t3) + (v2-v3)/(t2-t3))
		if math.IsNaN(m1) {
			m1 = 0
		}
		if math.IsNaN(m2) {
			m2 = 0
		}
		a := 2*v1 - 2*v2 + m1 + m2
		b := -3*v1 + 3*v2 - 2*m1 - m2
		c := m1
		d := v1
		segments = append(segments, [4]float64{a, b, c, d})
	}
	return segments
}

type catmullRomGradient struct {
	segments  [][4][4]float64
	positions []float64
	min       float64
	max       float64
	mode      BlendMode
	first     Color
	last      Color
}

func newCatmullRomGradient(colors []Color, positions []float64, space BlendMode) Gradient {
	n := len(colors)
	a := make([]float64, n)
	b := make([]float64, n)
	c := make([]float64, n)
	d := make([]float64, n)
	for i, col := range colors {
		var arr [4]float64
		switch space {
		case BlendRgb:
			arr = [4]float64{col.R, col.G, col.B, col.A}
		case BlendLinearRgb:
			arr = col2linearRgb(col)
		case BlendOklab:
			arr = col2oklab(col)
		}
		a[i] = arr[0]
		b[i] = arr[1]
		c[i] = arr[2]
		d[i] = arr[3]
	}
	s1 := toCatmullRomSegments(a)
	s2 := toCatmullRomSegments(b)
	s3 := toCatmullRomSegments(c)
	s4 := toCatmullRomSegments(d)
	segments := make([][4][4]float64, len(s1))
	for i, v1 := range s1 {
		segments[i] = [4][4]float64{
			v1,
			s2[i],
			s3[i],
			s4[i],
		}
	}
	min := positions[0]
	max := positions[n-1]
	gradbase := catmullRomGradient{
		segments:  segments,
		positions: positions,
		min:       min,
		max:       max,
		mode:      space,
		first:     colors[0],
		last:      colors[len(colors)-1],
	}
	return Gradient{
		grad: gradbase,
		dmin: min,
		dmax: max,
	}
}

func (g catmullRomGradient) At(t float64) Color {
	if math.IsNaN(t) {
		return Color{A: 1}
	}

	if t <= g.min {
		return g.first
	}

	if t >= g.max {
		return g.last
	}

	low := 0
	high := len(g.positions)

	for low < high {
		mid := (low + high) / 2
		if g.positions[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low == 0 {
		low = 1
	}

	pos0 := g.positions[low-1]
	pos1 := g.positions[low]
	seg_a := g.segments[low-1][0]
	seg_b := g.segments[low-1][1]
	seg_c := g.segments[low-1][2]
	seg_d := g.segments[low-1][3]

	t1 := (t - pos0) / (pos1 - pos0)
	t2 := t1 * t1
	t3 := t2 * t1

	a := seg_a[0]*t3 + seg_a[1]*t2 + seg_a[2]*t1 + seg_a[3]
	b := seg_b[0]*t3 + seg_b[1]*t2 + seg_b[2]*t1 + seg_b[3]
	c := seg_c[0]*t3 + seg_c[1]*t2 + seg_c[2]*t1 + seg_c[3]
	d := seg_d[0]*t3 + seg_d[1]*t2 + seg_d[2]*t1 + seg_d[3]

	switch g.mode {
	case BlendRgb:
		return Color{R: a, G: b, B: c, A: d}
	case BlendLinearRgb:
		return LinearRgb(a, b, c, d)
	case BlendOklab:
		return Oklab(a, b, c, d).Clamp()
	}

	return Color{}
}
