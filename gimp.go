package colorgrad

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

// References:
// https://gitlab.gnome.org/GNOME/gimp/-/blob/master/devel-docs/ggr.txt
// https://gitlab.gnome.org/GNOME/gimp/-/blob/master/app/core/gimpgradient.c
// https://gitlab.gnome.org/GNOME/gimp/-/blob/master/app/core/gimpgradient-load.c

const epsilon = 1e-10
const fracPi2 = math.Pi / 2

type blendingType int

const (
	linear blendingType = iota
	curved
	sinusoidal
	sphericalIncreasing
	sphericalDecreasing
	step
)

type coloringType int

const (
	rgb coloringType = iota
	hsvCcw
	hsvCw
)

type gimpSegment struct {
	// Left endpoint color
	lcolor colorful.Color
	// Right endpoint color
	rcolor colorful.Color
	// Left endpoint coordinate
	lpos float64
	// Midpoint coordinate
	mpos float64
	// Right endpoint coordinate
	rpos float64
	// Blending function type
	blending blendingType
	// Coloring type
	coloring coloringType
}

type gimpGradient struct {
	segments []gimpSegment
	dmin     float64
	dmax     float64
}

func (ggr gimpGradient) At(t float64) colorful.Color {
	if t <= ggr.dmin {
		return ggr.segments[0].lcolor
	}

	if t >= ggr.dmax {
		return ggr.segments[len(ggr.segments)-1].rcolor
	}

	if math.IsNaN(t) {
		return colorful.Color{R: 0, G: 0, B: 0}
	}

	low := 0
	high := len(ggr.segments)
	mid := 0

	for low < high {
		mid = (low + high) / 2
		if t > ggr.segments[mid].rpos {
			low = mid + 1
		} else if t < ggr.segments[mid].lpos {
			high = mid
		} else {
			break
		}
	}

	seg := ggr.segments[mid]
	seg_len := seg.rpos - seg.lpos

	var middle float64
	var pos float64

	if seg_len < epsilon {
		middle = 0.5
		pos = 0.5
	} else {
		middle = (seg.mpos - seg.lpos) / seg_len
		pos = (t - seg.lpos) / seg_len
	}

	var f float64

	switch seg.blending {
	case linear:
		f = calc_linear_factor(middle, pos)
	case curved:
		if middle < epsilon {
			return seg.rcolor
		} else if math.Abs(1-middle) < epsilon {
			return seg.lcolor
		} else {
			f = math.Exp(-math.Ln2 * math.Log10(pos) / math.Log10(middle))
		}
	case sinusoidal:
		x := calc_linear_factor(middle, pos)
		f = (math.Sin(-fracPi2+math.Pi*x) + 1) / 2
	case sphericalIncreasing:
		x := calc_linear_factor(middle, pos) - 1
		f = math.Sqrt(1 - x*x)
	case sphericalDecreasing:
		x := calc_linear_factor(middle, pos)
		f = 1 - math.Sqrt(1-x*x)
	case step:
		if pos >= middle {
			return seg.rcolor
		} else {
			return seg.lcolor
		}
	}

	switch seg.coloring {
	case rgb:
		return seg.lcolor.BlendRgb(seg.rcolor, f)
	case hsvCcw:
		return blendHsvCcw(seg.lcolor, seg.rcolor, f)
	case hsvCw:
		return blendHsvCw(seg.lcolor, seg.rcolor, f)
	}

	return ggr.segments[0].lcolor
}

func calc_linear_factor(middle, pos float64) float64 {
	if pos <= middle {
		if middle < epsilon {
			return 0
		} else {
			return 0.5 * pos / middle
		}
	} else {
		pos = pos - middle
		middle = 1 - middle

		if middle < epsilon {
			return 1
		} else {
			return 0.5 + 0.5*pos/middle
		}
	}
}

func blendHsvCcw(c1, c2 colorful.Color, t float64) colorful.Color {
	h1, s1, v1 := c1.Hsv()
	h2, s2, v2 := c2.Hsv()

	var hue float64

	if h1 < h2 {
		hue = h1 + ((h2 - h1) * t)
	} else {
		h := h1 + ((360 - (h1 - h2)) * t)

		if h > 360 {
			hue = h - 360
		} else {
			hue = h
		}
	}

	return colorful.Hsv(
		hue,
		s1+t*(s2-s1),
		v1+t*(v2-v1),
	)
}

func blendHsvCw(c1, c2 colorful.Color, t float64) colorful.Color {
	h1, s1, v1 := c1.Hsv()
	h2, s2, v2 := c2.Hsv()

	var hue float64

	if h2 < h1 {
		hue = h1 - ((h1 - h2) * t)
	} else {
		h := h1 - ((360 - (h2 - h1)) * t)

		if h < 0 {
			hue = h + 360
		} else {
			hue = h
		}
	}

	return colorful.Hsv(
		hue,
		s1+t*(s2-s1),
		v1+t*(v2-v1),
	)
}

func ParseGgr(r io.Reader, fg, bg colorful.Color) (Gradient, string, error) {
	zgrad := Gradient{
		grad: zeroGradient{},
		dmin: 0,
		dmax: 1,
	}

	segments := []gimpSegment{}
	var nseg int
	var name string
	xseg := 0
	i := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if i == 0 {
			if scanner.Text() != "GIMP Gradient" {
				return zgrad, name, fmt.Errorf("invalid header")
			}
		} else if i == 1 {
			if !strings.HasPrefix(scanner.Text(), "Name:") {
				return zgrad, name, fmt.Errorf("invalid header")
			}

			name = strings.TrimSpace(scanner.Text()[5:])
		} else if i == 2 {
			t, ok := parseFloat(scanner.Text())

			if ok {
				nseg = int(t)
			} else {
				return zgrad, name, fmt.Errorf("invalid header")
			}
		} else {
			if i >= nseg+3 {
				break
			}

			seg, ok := parseSegment(scanner.Text(), fg, bg)

			if ok {
				segments = append(segments, seg)
				xseg++
			} else {
				return zgrad, name, fmt.Errorf("invalid segment")
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return zgrad, name, err
	}

	if len(segments) == 0 {
		return zgrad, name, fmt.Errorf("segments %v", i)
	}

	if xseg < nseg {
		return zgrad, name, fmt.Errorf("wrong segments count, %v, %v", nseg, xseg)
	}

	gradbase := gimpGradient{
		segments: segments,
		dmin:     0,
		dmax:     1,
	}

	return Gradient{
		grad: gradbase,
		dmin: 0,
		dmax: 1,
	}, name, nil
}

func parseSegment(s string, fg, bg colorful.Color) (gimpSegment, bool) {
	params := strings.Fields(s)
	plen := len(params)

	if plen != 13 && plen != 15 {
		return gimpSegment{}, false
	}

	d := make([]float64, 15)

	for i, x := range params {
		t, ok := parseFloat(x)

		if ok {
			d[i] = t
			continue
		}

		return gimpSegment{}, false
	}

	if plen == 13 {
		d[13] = 0
		d[14] = 0
	}

	var blending blendingType

	switch int(d[11]) {
	case 0:
		blending = linear
	case 1:
		blending = curved
	case 2:
		blending = sinusoidal
	case 3:
		blending = sphericalIncreasing
	case 4:
		blending = sphericalDecreasing
	case 5:
		blending = step
	default:
		return gimpSegment{}, false
	}

	var coloring coloringType

	switch int(d[12]) {
	case 0:
		coloring = rgb
	case 1:
		coloring = hsvCcw
	case 2:
		coloring = hsvCw
	default:
		return gimpSegment{}, false
	}

	var lcolor colorful.Color

	switch int(d[13]) {
	case 0:
		lcolor = colorful.Color{R: d[3], G: d[4], B: d[5]}
	case 1:
		lcolor = fg
	case 2:
		lcolor = fg // TODO transparent
	case 3:
		lcolor = bg
	case 4:
		lcolor = bg // TODO transparent
	default:
		return gimpSegment{}, false
	}

	var rcolor colorful.Color

	switch int(d[14]) {
	case 0:
		rcolor = colorful.Color{R: d[7], G: d[8], B: d[9]}
	case 1:
		rcolor = fg
	case 2:
		rcolor = fg // TODO transparent
	case 3:
		rcolor = bg
	case 4:
		rcolor = bg // TODO transparent
	default:
		return gimpSegment{}, false
	}

	return gimpSegment{
		lcolor:   lcolor,
		rcolor:   rcolor,
		lpos:     d[0],
		mpos:     d[1],
		rpos:     d[2],
		blending: blending,
		coloring: coloring,
	}, true
}
