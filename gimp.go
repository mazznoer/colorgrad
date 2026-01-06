package colorgrad

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
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
	lcolor Color
	// Right endpoint color
	rcolor Color
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
	min      float64
	max      float64
}

func (ggr gimpGradient) At(t float64) Color {
	if t <= ggr.min {
		return ggr.segments[0].lcolor
	}

	if t >= ggr.max {
		return ggr.segments[len(ggr.segments)-1].rcolor
	}

	if math.IsNaN(t) {
		return Color{A: 1}
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
		return blendRgb(seg.lcolor, seg.rcolor, f)
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

func blendHsvCcw(c1, c2 Color, t float64) Color {
	hsvA := col2hsv(c1)
	hsvB := col2hsv(c2)

	var hue float64

	if hsvA[0] < hsvB[0] {
		hue = hsvA[0] + ((hsvB[0] - hsvA[0]) * t)
	} else {
		h := hsvA[0] + ((360 - (hsvA[0] - hsvB[0])) * t)

		if h > 360 {
			hue = h - 360
		} else {
			hue = h
		}
	}

	return Hsv(
		hue,
		hsvA[1]+t*(hsvB[1]-hsvA[1]),
		hsvA[2]+t*(hsvB[2]-hsvA[2]),
		hsvA[3]+t*(hsvB[3]-hsvA[3]),
	)
}

func blendHsvCw(c1, c2 Color, t float64) Color {
	hsvA := col2hsv(c1)
	hsvB := col2hsv(c2)

	var hue float64

	if hsvB[0] < hsvA[0] {
		hue = hsvA[0] - ((hsvA[0] - hsvB[0]) * t)
	} else {
		h := hsvA[0] - ((360 - (hsvB[0] - hsvA[0])) * t)

		if h < 0 {
			hue = h + 360
		} else {
			hue = h
		}
	}

	return Hsv(
		hue,
		hsvA[1]+t*(hsvB[1]-hsvA[1]),
		hsvA[2]+t*(hsvB[2]-hsvA[2]),
		hsvA[3]+t*(hsvB[3]-hsvA[3]),
	)
}

func ParseGgr(r io.Reader, fg, bg Color) (Gradient, string, error) {
	zgrad := Gradient{
		Core: zeroGradient{},
		Min:  0,
		Max:  1,
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
		min:      0,
		max:      1,
	}

	return Gradient{
		Core: gradbase,
		Min:  0,
		Max:  1,
	}, name, nil
}

func parseSegment(s string, fg, bg Color) (gimpSegment, bool) {
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

	var lcolor Color

	switch int(d[13]) {
	case 0:
		lcolor = Color{R: d[3], G: d[4], B: d[5], A: d[6]}
	case 1:
		lcolor = fg
	case 2:
		lcolor = Rgb(fg.R, fg.G, fg.B, 0)
	case 3:
		lcolor = bg
	case 4:
		lcolor = Rgb(bg.R, bg.G, bg.B, 0)
	default:
		return gimpSegment{}, false
	}

	var rcolor Color

	switch int(d[14]) {
	case 0:
		rcolor = Color{R: d[7], G: d[8], B: d[9], A: d[10]}
	case 1:
		rcolor = fg
	case 2:
		rcolor = Rgb(fg.R, fg.G, fg.B, 0)
	case 3:
		rcolor = bg
	case 4:
		rcolor = Rgb(bg.R, bg.G, bg.B, 0)
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
