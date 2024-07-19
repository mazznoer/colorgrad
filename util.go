package colorgrad

import (
	"math"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
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

func interpAngle(a, b, t float64) float64 {
	delta := math.Mod(((math.Mod(b-a, 360))+540), 360) - 180
	return math.Mod((a + t*delta + 360), 360)
}

func blendRgb(a, b Color, t float64) Color {
	return Color{
		R: a.R + t*(b.R-a.R),
		G: a.G + t*(b.G-a.G),
		B: a.B + t*(b.B-a.B),
		A: a.A + t*(b.A-a.A),
	}
}

func blendLrgb(a, b colorful.Color, t float64) colorful.Color {
	r1, g1, b1 := a.LinearRgb()
	r2, g2, b2 := b.LinearRgb()
	return colorful.LinearRgb(
		r1+t*(r2-r1),
		g1+t*(g2-g1),
		b1+t*(b2-b1),
	)
}

func blendOklab(ca, cb colorful.Color, t float64) colorful.Color {
	r1, g1, b1 := ca.LinearRgb()
	r2, g2, b2 := cb.LinearRgb()
	l1, a1, b1 := lrgbToOklab(r1, g1, b1)
	l2, a2, b2 := lrgbToOklab(r2, g2, b2)
	r, g, b := oklabToLrgb(
		l1+t*(l2-l1),
		a1+t*(a2-a1),
		b1+t*(b2-b1),
	)
	return colorful.LinearRgb(r, g, b)
}

func lrgbToOklab(r, g, b float64) (L, A, B float64) {
	l := math.Cbrt(0.4121656120*r + 0.5362752080*g + 0.0514575653*b)
	m := math.Cbrt(0.2118591070*r + 0.6807189584*g + 0.1074065790*b)
	s := math.Cbrt(0.0883097947*r + 0.2818474174*g + 0.6302613616*b)
	L = 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	A = 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	B = 0.0259040371*l + 0.7827717662*m - 0.8086757660*s
	return
}

func oklabToLrgb(l, a, b float64) (R, G, B float64) {
	l_ := math.Pow(l+0.3963377774*a+0.2158037573*b, 3)
	m_ := math.Pow(l-0.1055613458*a-0.0638541728*b, 3)
	s_ := math.Pow(l-0.0894841775*a-1.2914855480*b, 3)
	R = 4.0767245293*l_ - 3.3072168827*m_ + 0.2307590544*s_
	G = -1.2681437731*l_ + 2.6093323231*m_ - 0.3411344290*s_
	B = -0.0041119885*l_ - 0.7034763098*m_ + 1.7068625689*s_
	return
}
