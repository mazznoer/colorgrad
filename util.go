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

func toLinear(x float64) float64 {
	if x >= 0.04045 {
		return math.Pow((x+0.055)/1.055, 2.4)
	}
	return x / 12.92
}

func col2linearRgb(col Color) [4]float64 {
	return [4]float64{
		toLinear(col.R),
		toLinear(col.G),
		toLinear(col.B),
		col.A,
	}
}

func col2oklab(col Color) [4]float64 {
	arr := col2linearRgb(col)
	l := math.Cbrt(0.4121656120*arr[0] + 0.5362752080*arr[1] + 0.0514575653*arr[2])
	m := math.Cbrt(0.2118591070*arr[0] + 0.6807189584*arr[1] + 0.1074065790*arr[2])
	s := math.Cbrt(0.0883097947*arr[0] + 0.2818474174*arr[1] + 0.6302613616*arr[2])
	return [4]float64{
		0.2104542553*l + 0.7936177850*m - 0.0040720468*s,
		1.9779984951*l - 2.4285922050*m + 0.4505937099*s,
		0.0259040371*l + 0.7827717662*m - 0.8086757660*s,
		col.A,
	}
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
			colors[i] = col2linearRgb(col)
		case BlendLab:
			colors[i] = col2lab(col)
		case BlendOklab:
			colors[i] = col2oklab(col)
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

// --- Lab

const (
	d65X = 0.95047
	d65Y = 1.0
	d65Z = 1.08883

	delta  = 6.0 / 29.0
	delta2 = delta * delta
	delta3 = delta2 * delta
)

func linearRGBToXYZ(r, g, b float64) [3]float64 {
	// Inverse sRGB matrix (D65)
	x := 0.4124564*r + 0.3575761*g + 0.1804375*b
	y := 0.2126729*r + 0.7151522*g + 0.0721750*b
	z := 0.0193339*r + 0.1191920*g + 0.9503041*b
	return [3]float64{x, y, z}
}

func xyzToLab(x, y, z float64) [3]float64 {
	labF := func(t float64) float64 {
		if t > delta3 {
			return math.Cbrt(t)
		}
		return (t / (3.0 * delta2)) + (4.0 / 29.0)
	}

	fx := labF(x / d65X)
	fy := labF(y / d65Y)
	fz := labF(z / d65Z)

	l := 116.0*fy - 16.0
	a := 500.0 * (fx - fy)
	b := 200.0 * (fy - fz)

	return [3]float64{l, a, b}
}

func col2lab(col Color) [4]float64 {
	c := col2linearRgb(col)
	x := linearRGBToXYZ(c[0], c[1], c[2])
	l := xyzToLab(x[0], x[1], x[2])
	return [4]float64{l[0], l[1], l[2], col.A}
}
