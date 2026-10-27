package colorgrad

import (
	"math"
	"slices"
)

type BasisGradient struct {
	// Control points padded with one "ghost" point at each end:
	// ctrl[k+1] is color k, ctrl[0] and ctrl[len-1] are linear extrapolations.
	// Segment i (between colors i and i+1) therefore always uses
	// ctrl[i], ctrl[i+1], ctrl[i+2], ctrl[i+3], with no edge-case branches.
	ctrl        [][4]float64
	positions   []float64
	min, max    float64
	mode        BlendMode
	first, last Color
}

func (g *BasisGradient) At(t float64) Color {
	if math.IsNaN(t) {
		return Color{A: 1}
	}

	if t <= g.min {
		return g.first
	}

	if t >= g.max {
		return g.last
	}

	// Here positions[0] < t < positions[len-1], so the first index with
	// positions[idx] >= t is guaranteed to lie in [1, len-1].
	lo, hi := 1, len(g.positions)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if g.positions[mid] < t {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	// Segment is [lo-1, lo]; pos[lo-1] < t <= pos[lo], so the divisor is > 0.
	i := lo - 1
	u := (t - g.positions[i]) / (g.positions[lo] - g.positions[i])

	// Blend weights
	u2 := u * u
	u3 := u2 * u
	w0 := 1 - 3*u + 3*u2 - u3
	w1 := 4 - 6*u2 + 3*u3
	w2 := 1 + 3*u + 3*u2 - 3*u3
	w3 := u3

	p0 := &g.ctrl[i]
	p1 := &g.ctrl[i+1]
	p2 := &g.ctrl[i+2]
	p3 := &g.ctrl[i+3]

	a := (w0*p0[0] + w1*p1[0] + w2*p2[0] + w3*p3[0]) / 6
	b := (w0*p0[1] + w1*p1[1] + w2*p2[1] + w3*p3[1]) / 6
	c := (w0*p0[2] + w1*p1[2] + w2*p2[2] + w3*p3[2]) / 6
	d := (w0*p0[3] + w1*p1[3] + w2*p2[3] + w3*p3[3]) / 6

	switch g.mode {
	case BlendLinearRgb:
		return LinearRgb(a, b, c, d)
	case BlendLab:
		return Lab(a, b, c, d).Clamp()
	case BlendOklab:
		return Oklab(a, b, c, d).Clamp()
	default:
		return Color{R: a, G: b, B: c, A: d}
	}
}

func newBasisGradient(colors []Color, positions []float64, mode BlendMode) Gradient {
	pos := slices.Clone(positions)
	n := len(colors)
	base := convertColors(colors, mode)

	ctrl := make([][4]float64, n+2)
	copy(ctrl[1:], base)

	for j := 0; j < 4; j++ {
		ctrl[0][j] = 2*base[0][j] - base[1][j]
		ctrl[n+1][j] = 2*base[n-1][j] - base[n-2][j]
	}

	g := &BasisGradient{
		ctrl:      ctrl,
		positions: pos,
		min:       pos[0],
		max:       pos[n-1],
		mode:      mode,
		first:     colors[0],
		last:      colors[n-1],
	}

	return Gradient{
		Core: g,
		Min:  g.min,
		Max:  g.max,
	}
}
