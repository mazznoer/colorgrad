package colorgrad

import (
	"github.com/lucasb-eyer/go-colorful"
)

type linearGradient struct {
	colors []colorful.Color
	pos    []float64
	dmin   float64
	dmax   float64
	count  int
	mode   BlendMode
}

func (lg linearGradient) At(t float64) colorful.Color {
	if t < lg.dmin {
		return lg.colors[0]
	}
	if t > lg.dmax {
		return lg.colors[lg.count]
	}
	for i := 0; i < lg.count; i++ {
		p1 := lg.pos[i]
		p2 := lg.pos[i+1]
		if (p1 <= t) && (t <= p2) {
			t := (t - p1) / (p2 - p1)
			a := lg.colors[i]
			b := lg.colors[i+1]
			switch lg.mode {
			case BlendHcl:
				return a.BlendHcl(b, t).Clamped()
			case BlendHsv:
				return a.BlendHsv(b, t)
			case BlendLab:
				return a.BlendLab(b, t).Clamped()
			case BlendLinearRgb:
				return blendLrgb(a, b, t)
			case BlendLuv:
				return a.BlendLuv(b, t).Clamped()
			case BlendRgb:
				return a.BlendRgb(b, t)
			case BlendOklab:
				return blendOklab(a, b, t).Clamped()
			}
		}
	}
	return lg.colors[0]
}
