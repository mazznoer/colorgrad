package colorgrad

import "math"

const (
	maxDepth = 8
	seeds    = 36
)

type gradFn func(t float64) Color

func linearize(grad gradFn, dmin, dmax, threshold float64) Gradient {
	if math.IsNaN(threshold) || math.IsInf(threshold, 0) {
		threshold = 0.007
	} else {
		threshold = clamp(threshold, 0.001, 0.035)
	}

	t0 := dmin
	c0 := grad(t0).Clamp()
	sf := newStreamFilter(threshold, t0, c0, 500)

	// Adaptive sampling
	for i := 1; i <= seeds; i++ {
		t1 := dmin + (dmax-dmin)*(float64(i)/float64(seeds))
		c1 := grad(t1).Clamp()

		// Subdivide interval [t0, t1]
		subdivide(grad, t0, t1, c0, c1, threshold, 0, sf)

		// Push right seed boundary
		sf.push(t1, c1)

		t0 = t1
		c0 = c1
	}

	positions := sf.finish()
	colors := make([]Color, len(positions))
	for i := range positions {
		colors[i] = grad(positions[i]).Clamp()
	}

	return newLinearGradient(colors, positions, BlendRgb)
}

func subdivide(
	grad gradFn,
	t0, t1 float64,
	c0, c1 Color,
	threshold float64,
	depth int,
	sf *streamFilter,
) {
	if depth > maxDepth {
		return
	}

	mid := (t0 + t1) / 2.0
	cMidActual := grad(mid).Clamp()
	cMidLinear := blendRgb(c0, c1, 0.5)

	if colorDiff(cMidActual, cMidLinear) > threshold {
		// Left branch (t0 -> mid)
		subdivide(grad, t0, mid, c0, cMidActual, threshold, depth+1, sf)

		// In-order midpoint push
		sf.push(mid, cMidActual)

		// Right branch (mid -> t1)
		subdivide(grad, mid, t1, cMidActual, c1, threshold, depth+1, sf)
	}
}

// streamFilter is a streaming pruner that maintains a 3-point sliding window
// with cached colors to perform on-the-fly collinear pruning and deduplication.
type streamFilter struct {
	threshold float64
	committed []float64
	colorA    Color

	// Pending candidate
	hasCandidate bool
	candT        float64
	candColor    Color
}

func newStreamFilter(threshold, lo float64, initialColor Color, capacity int) *streamFilter {
	committed := make([]float64, 0, capacity)
	committed = append(committed, lo)
	return &streamFilter{
		threshold: threshold,
		committed: committed,
		colorA:    initialColor,
	}
}

func (sf *streamFilter) push(c float64, colorC Color) {
	a := sf.committed[len(sf.committed)-1]

	// Skip duplicates close to last committed point
	if math.Abs(c-a) < epsilon {
		return
	}

	if sf.hasCandidate {
		b, colorB := sf.candT, sf.candColor

		// Deduplicate candidate against incoming point
		if math.Abs(c-b) < epsilon {
			sf.candT, sf.candColor = c, colorC
			return
		}

		// Evaluate if candidate B is collinear between A and C using cached colors
		lerpFactor := (b - a) / (c - a)
		cCurrLinear := blendRgb(sf.colorA, colorC, lerpFactor)

		if colorDiff(colorB, cCurrLinear) > sf.threshold {
			// B is essential -> commit B and update last committed color A
			sf.committed = append(sf.committed, b)
			sf.colorA = colorB
		}
		// If B is collinear, it drops automatically when C replaces it
	}

	sf.hasCandidate = true
	sf.candT, sf.candColor = c, colorC
}

func (sf *streamFilter) finish() []float64 {
	if sf.hasCandidate {
		a := sf.committed[len(sf.committed)-1]
		if math.Abs(sf.candT-a) >= epsilon {
			sf.committed = append(sf.committed, sf.candT)
		}
	}
	return sf.committed
}

// colorDiff is the Euclidean distance between two colors in RGBA space,
// normalized to [0.0, 1.0].
func colorDiff(a, b Color) float64 {
	dr := a.R - b.R
	dg := a.G - b.G
	db := a.B - b.B
	da := a.A - b.A
	return math.Sqrt(dr*dr+dg*dg+db*db+da*da) / 2.0
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
