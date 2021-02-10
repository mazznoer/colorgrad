package colorgrad

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/mazznoer/csscolorparser"
)

func basis(t1, v0, v1, v2, v3 float64) float64 {
	t2 := t1 * t1
	t3 := t2 * t1
	return ((1-3*t1+3*t2-t3)*v0 + (4-6*t2+3*t3)*v1 + (1+3*t1+3*t2-3*t3)*v2 + t3*v3) / 6
}

type spline struct {
	values []float64
}

func (s spline) at(t float64) float64 {
	n := len(s.values) - 1
	var i int
	if t <= 0 {
		t = 0
		i = 0
	} else if t >= 1 {
		t = 1
		i = n - 1
	} else {
		i = int(t * float64(n))
	}
	v1 := s.values[i]
	v2 := s.values[i+1]
	var v0 float64
	if i > 0 {
		v0 = s.values[i-1]
	} else {
		v0 = 2*v1 - v2
	}
	var v3 float64
	if i < n-1 {
		v3 = s.values[i+2]
	} else {
		v3 = 2*v2 - v1
	}
	return basis((t-float64(i)/float64(n))*float64(n), v0, v1, v2, v3)
}

type splineGradient struct {
	r spline
	g spline
	b spline
}

func (sg splineGradient) At(t float64) colorful.Color {
	if math.IsNaN(t) {
		t = 0
	}
	return colorful.Color{
		R: sg.r.at(t),
		G: sg.g.at(t),
		B: sg.b.at(t),
	}
}

func presetSpline(htmlColors []string) Gradient {
	var colors []csscolorparser.Color
	for _, s := range htmlColors {
		c, err := csscolorparser.Parse(s)
		if err == nil {
			colors = append(colors, c)
		}
	}
	n := len(colors)
	reds := make([]float64, n)
	greens := make([]float64, n)
	blues := make([]float64, n)
	for i, c := range colors {
		reds[i] = c.R
		greens[i] = c.G
		blues[i] = c.B
	}
	gradbase := splineGradient{
		r: spline{values: reds},
		g: spline{values: greens},
		b: spline{values: blues},
	}
	return Gradient{
		grad: gradbase,
		dmin: 0,
		dmax: 1,
	}
}
