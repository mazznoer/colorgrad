package colorgrad

import (
	"math"

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
