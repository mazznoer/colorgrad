package colorgrad

import (
	"math"
	"strconv"
	"strings"
)

func linspace(min, max float64, n uint) []float64 {
	if n == 1 {
		return []float64{min}
	}
	d := max - min
	l := float64(n) - 1
	res := make([]float64, n)
	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}

// Map t from range [a, b] to range [0, 1]
func norm(t, a, b float64) float64 {
	return (t - a) * (1 / (b - a))
}

func modulo(x, y float64) float64 {
	return math.Mod(math.Mod(x, y)+y, y)
}

func clamp01(t float64) float64 {
	return math.Max(0, math.Min(1, t))
}

func parseFloat(s string) (float64, bool) {
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return f, err == nil
}

func col2hsv(col Color) [4]float64 {
	v := math.Max(col.R, math.Max(col.G, col.B))
	d := v - math.Min(col.R, math.Min(col.G, col.B))

	if math.Abs(d) < epsilon {
		return [4]float64{0, 0, v, col.A}
	}

	s := d / v
	dr := (v - col.R) / d
	dg := (v - col.G) / d
	db := (v - col.B) / d

	var h float64

	if math.Abs(col.R-v) < epsilon {
		h = db - dg
	} else if math.Abs(col.G-v) < epsilon {
		h = 2.0 + dr - db
	} else {
		h = 4.0 + dg - dr
	}

	h = math.Mod(h*60.0, 360.0)
	return [4]float64{normalizeAngle(h), s, v, col.A}
}

func normalizeAngle(t float64) float64 {
	t = math.Mod(t, 360.0)
	if t < 0.0 {
		t += 360.0
	}
	return t
}

func convertColors(colorsIn []Color, mode BlendMode) [][4]float64 {
	colors := make([][4]float64, len(colorsIn))
	for i, col := range colorsIn {
		switch mode {
		case BlendRgb:
			colors[i] = [4]float64{col.R, col.G, col.B, col.A}
		case BlendLinearRgb:
			colors[i] = col.LinearRgb()
		case BlendLab:
			colors[i] = col.Lab()
		case BlendOklab:
			colors[i] = col.Oklab()
		}
	}
	return colors
}

func linearInterpolate(a, b [4]float64, t float64) (i, j, k, l float64) {
	i = a[0] + t*(b[0]-a[0])
	j = a[1] + t*(b[1]-a[1])
	k = a[2] + t*(b[2]-a[2])
	l = a[3] + t*(b[3]-a[3])
	return
}

func blendRgb(a, b Color, t float64) Color {
	return Color{
		R: a.R + t*(b.R-a.R),
		G: a.G + t*(b.G-a.G),
		B: a.B + t*(b.B-a.B),
		A: a.A + t*(b.A-a.A),
	}
}
