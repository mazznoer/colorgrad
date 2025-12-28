package colorgrad

import (
	"math"
	"strings"

	"github.com/mazznoer/csscolorparser"
)

var eps = math.Nextafter(1.0, 2.0) - 1.0

const tau = math.Pi * 2

func parseCss(s string) ([]cssGradientStop, bool) {
	stops := []cssGradientStop{}

	for _, stop := range splitByComma(s) {
		if !prosesStop(&stops, splitBySpace(stop)) {
			return stops, false
		}
	}

	if len(stops) == 0 {
		return stops, false
	}

	if stops[0].color == nil {
		return stops, false
	}

	for i, stop := range stops {
		if i == 0 && stop.pos == nil {
			stops[i].pos = ptr(0.0)
		}

		if i == len(stops)-1 {
			if stop.pos == nil {
				stops[i].pos = ptr(1.0)
			}
			continue
		}

		if stop.color == nil {
			stops[i].color = ptrColor(blendRgb(*stops[i-1].color, *stops[i+1].color, 0.5))
		}
	}

	if *stops[0].pos > 0.0 {
		stops = append([]cssGradientStop{{ptr(0.0), stops[0].color}}, stops...)
	}

	if *stops[len(stops)-1].pos < 1.0 {
		stops = append(stops, cssGradientStop{ptr(1.0), stops[len(stops)-1].color})
	}

	for i, stop := range stops {
		if stop.pos == nil {
			for j := i + 1; j < len(stops); j++ {
				if stops[j].pos != nil {
					prev := *stops[i-1].pos
					next := *stops[j].pos
					stops[i].pos = ptr(prev + (next-prev)/float64(j-i+1))
					break
				}
			}
		}

		if i > 0 {
			stops[i].pos = ptr(math.Max(*stops[i].pos, *stops[i-1].pos))
		}
	}

	// Filter Stops

	prev := stops[0].pos
	last := len(stops) - 1
	newStops := make([]cssGradientStop, 0, len(stops))

	for i, s := range stops {
		var next *float64

		if i == last {
			next = stops[last].pos
		} else {
			next = stops[i+1].pos
		}

		if (*s.pos-*prev)+(*next-*s.pos) < eps {
			// skip 0 width stop
		} else {
			newStops = append(newStops, s)
		}
		prev = s.pos
	}

	return newStops, true
}

func ptr(f float64) *float64 {
	return &f
}

func ptrColor(c Color) *Color {
	return &c
}

type cssGradientStop struct {
	pos   *float64
	color *Color
}

func prosesStop(stops *[]cssGradientStop, arr []string) bool {
	switch len(arr) {
	case 1:
		col, err := csscolorparser.Parse(arr[0])
		if err == nil {
			*stops = append(*stops, cssGradientStop{nil, &col})
			return true
		}

		pos, ok := parsePos(arr[0])
		if ok {
			*stops = append(*stops, cssGradientStop{&pos, nil})
			return true
		}
		return false
	case 2:
		col, err := csscolorparser.Parse(arr[0])
		if err != nil {
			return false
		}

		pos, ok := parsePos(arr[1])
		if !ok {
			return false
		}

		*stops = append(*stops, cssGradientStop{&pos, &col})
	case 3:
		col, err := csscolorparser.Parse(arr[0])
		if err != nil {
			return false
		}

		pos1, ok1 := parsePos(arr[1])
		if !ok1 {
			return false
		}

		pos2, ok2 := parsePos(arr[2])
		if !ok2 {
			return false
		}

		*stops = append(*stops, cssGradientStop{&pos1, &col})
		*stops = append(*stops, cssGradientStop{&pos2, &col})
	default:
		return false
	}
	return true
}

func splitByComma(s string) []string {
	res := []string{}
	beg := 0
	inside := false

	for i := 0; i < len(s); i++ {
		if s[i] == ',' && !inside {
			res = append(res, s[beg:i])
			beg = i + 1
		} else if s[i] == '(' {
			inside = true
		} else if s[i] == ')' {
			inside = false
		}
	}
	return append(res, s[beg:])
}

func splitBySpace(s string) []string {
	res := []string{}
	beg := 0
	inside := false

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && !inside {
			if len(s[beg:i]) > 0 {
				res = append(res, s[beg:i])
			}
			beg = i + 1
		} else if s[i] == '(' {
			inside = true
		} else if s[i] == ')' {
			inside = false
		}
	}
	if len(s[beg:]) > 0 {
		res = append(res, s[beg:])
	}
	return res
}

func parsePos(s string) (float64, bool) {
	if strings.HasSuffix(s, "%") {
		f, ok := parseFloat(s[:len(s)-1])
		if ok {
			return f / 100, true
		}
		return 0, false
	}

	f, ok := parseFloat(s)
	return f, ok
}
