//go:build ignore
package colorgrad

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// Adapted from https://qroph.github.io/2018/07/30/smooth-paths-using-catmull-rom-splines.html

type catmullRomInterpolator struct {
	segments [][4]float64
	pos      []float64
}

func newCatmullRomInterpolator(values, pos []float64) catmullRomInterpolator {
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

	return catmullRomInterpolator{
		segments,
		pos,
	}
}

func (cr catmullRomInterpolator) at(t float64) float64 {
	low := 0
	high := len(cr.pos)

	for low < high {
		mid := (low + high) / 2
		if cr.pos[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low == 0 {
		low = 1
	}

	p1 := cr.pos[low-1]
	p2 := cr.pos[low]
	seg := cr.segments[low-1]
	t1 := (t - p1) / (p2 - p1)
	t2 := t1 * t1
	t3 := t2 * t1
	return seg[0]*t3 + seg[1]*t2 + seg[2]*t1 + seg[3]
}

type catmullRomGradient struct {
	a    catmullRomInterpolator
	b    catmullRomInterpolator
	c    catmullRomInterpolator
	dmin float64
	dmax float64
	mode BlendMode
}

func (s catmullRomGradient) At(t float64) colorful.Color {
	if math.IsNaN(t) {
		return colorful.Color{R: 0, G: 0, B: 0}
	}
	t = math.Max(s.dmin, math.Min(s.dmax, t))
	switch s.mode {
	case BlendLinearRgb:
		return colorful.LinearRgb(s.a.at(t), s.b.at(t), s.c.at(t))
	case BlendLab:
		return colorful.Lab(s.a.at(t), s.b.at(t), s.c.at(t)).Clamped()
	case BlendLuv:
		return colorful.Luv(s.a.at(t), s.b.at(t), s.c.at(t)).Clamped()
	case BlendHcl:
		return colorful.Hcl(s.a.at(t), s.b.at(t), s.c.at(t)).Clamped()
	case BlendHsv:
		return colorful.Hsv(s.a.at(t), s.b.at(t), s.c.at(t))
	case BlendOklab:
		r, g, b := oklabToLrgb(s.a.at(t), s.b.at(t), s.c.at(t))
		return colorful.LinearRgb(r, g, b).Clamped()
	default:
		return colorful.Color{R: s.a.at(t), G: s.b.at(t), B: s.c.at(t)}
	}
}

func newCatmullRomGradient(colors []colorful.Color, pos []float64, space BlendMode) Gradient {
	n := len(colors)
	a := make([]float64, n)
	b := make([]float64, n)
	c := make([]float64, n)
	for i, col := range colors {
		var c1, c2, c3 float64
		switch space {
		case BlendLinearRgb:
			c1, c2, c3 = col.LinearRgb()
		case BlendLab:
			c1, c2, c3 = col.Lab()
		case BlendLuv:
			c1, c2, c3 = col.Luv()
		case BlendHcl:
			c1, c2, c3 = col.Hcl()
		case BlendHsv:
			c1, c2, c3 = col.Hsv()
		case BlendOklab:
			lr, lg, lb := col.LinearRgb()
			c1, c2, c3 = lrgbToOklab(lr, lg, lb)
		case BlendRgb:
			c1, c2, c3 = col.R, col.G, col.B
		}
		a[i] = c1
		b[i] = c2
		c[i] = c3
	}
	dmin := pos[0]
	dmax := pos[n-1]
	gradbase := catmullRomGradient{
		a:    newCatmullRomInterpolator(a, pos),
		b:    newCatmullRomInterpolator(b, pos),
		c:    newCatmullRomInterpolator(c, pos),
		dmin: dmin,
		dmax: dmax,
		mode: space,
	}
	return Gradient{
		grad: gradbase,
		dmin: dmin,
		dmax: dmax,
	}
}
