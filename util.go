package colorgrad

import "github.com/lucasb-eyer/go-colorful"

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
