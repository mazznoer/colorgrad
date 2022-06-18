package colorgrad

import (
	"github.com/lucasb-eyer/go-colorful"
)

type sharpGradient struct {
	colors []colorful.Color
	pos    []float64
	last   int
	dmin   float64
	dmax   float64
}

func (sg sharpGradient) At(t float64) colorful.Color {
	if t <= sg.dmin {
		return sg.colors[0]
	}
	if t >= sg.dmax {
		return sg.colors[sg.last]
	}
	for i := 0; i < sg.last; i++ {
		p1 := sg.pos[i]
		p2 := sg.pos[i+1]
		if (p1 <= t) && (t <= p2) {
			if i%2 == 0 {
				return sg.colors[i]
			}
			t := (t - p1) / (p2 - p1)
			a := sg.colors[i]
			b := sg.colors[i+1]
			return a.BlendRgb(b, t)
		}
	}
	return sg.colors[0]
}

func newSharpGradient(colorsIn []colorful.Color, dmin, dmax float64, smoothness float64) Gradient {
	n := len(colorsIn)
	colors := make([]colorful.Color, n*2)
	i := 0
	for _, c := range colorsIn {
		colors[i] = c
		i++
		colors[i] = c
		i++
	}
	t := clamp01(smoothness) * (dmax - dmin) / float64(n) / 4
	p := linspace(dmin, dmax, uint(n+1))
	pos := make([]float64, n*2)
	i = 0
	j := 0
	for x := 0; x < int(n); x++ {
		pos[i] = p[j]
		if i > 0 {
			pos[i] += t
		}
		i++
		j++
		pos[i] = p[j]
		if i < len(colors)-1 {
			pos[i] -= t
		}
		i++
	}
	gradbase := sharpGradient{
		colors: colors,
		pos:    pos,
		last:   int(n*2 - 1),
		dmin:   dmin,
		dmax:   dmax,
	}
	return Gradient{
		grad: gradbase,
		dmin: dmin,
		dmax: dmax,
	}
}
