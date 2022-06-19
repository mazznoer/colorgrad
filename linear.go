package colorgrad

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

type linearGradient struct {
	colors     [][3]float64
	pos        []float64
	dmin       float64
	dmax       float64
	count      int
	mode       BlendMode
	firstColor colorful.Color
	lastColor  colorful.Color
}

func (lg linearGradient) At(t float64) colorful.Color {
	if t <= lg.dmin {
		return lg.firstColor
	}
	if t >= lg.dmax {
		return lg.lastColor
	}
	for i := 0; i < lg.count; i++ {
		p1 := lg.pos[i]
		p2 := lg.pos[i+1]
		if (p1 <= t) && (t <= p2) {
			t := (t - p1) / (p2 - p1)
			x, y, z := linearInterpolate(lg.colors[i], lg.colors[i+1], t)
			switch lg.mode {
			case BlendHcl:
				hue := interpAngle(lg.colors[i][0], lg.colors[i+1][0], t)
				return colorful.Hcl(hue, y, z).Clamped()
			case BlendHsv:
				hue := interpAngle(lg.colors[i][0], lg.colors[i+1][0], t)
				return colorful.Hsv(hue, y, z)
			case BlendLab:
				return colorful.Lab(x, y, z).Clamped()
			case BlendLinearRgb:
				return colorful.LinearRgb(x, y, z)
			case BlendLuv:
				return colorful.Luv(x, y, z).Clamped()
			case BlendRgb:
				return colorful.Color{R: x, G: y, B: z}
			case BlendOklab:
				a, b, c := oklabToLrgb(x, y, z)
				return colorful.LinearRgb(a, b, c).Clamped()
			}
		}
	}
	return colorful.Color{R: 0, G: 0, B: 0}
}

func newLinearGradient(colors []colorful.Color, pos []float64, mode BlendMode) Gradient {
	gradbase := linearGradient{
		colors:     convertColors(colors, mode),
		pos:        pos,
		dmin:       pos[0],
		dmax:       pos[len(pos)-1],
		count:      len(colors) - 1,
		mode:       mode,
		firstColor: colors[0],
		lastColor:  colors[len(colors)-1],
	}

	return Gradient{
		grad: gradbase,
		dmin: pos[0],
		dmax: pos[len(pos)-1],
	}
}

func convertColors(colorsIn []colorful.Color, mode BlendMode) [][3]float64 {
	colors := make([][3]float64, len(colorsIn))
	for i, col := range colorsIn {
		var c1, c2, c3 float64
		switch mode {
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
		colors[i] = [3]float64{c1, c2, c3}
	}
	return colors
}

func linearInterpolate(a, b [3]float64, t float64) (x, y, z float64) {
	x = a[0] + t*(b[0]-a[0])
	y = a[1] + t*(b[1]-a[1])
	z = a[2] + t*(b[2]-a[2])
	return
}

func interpAngle(a, b, t float64) float64 {
	delta := math.Mod(((math.Mod(b-a, 360))+540), 360) - 180
	return math.Mod((a + t*delta + 360), 360)
}
