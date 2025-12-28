package colorgrad

import (
	"math"
	"strings"

	"github.com/mazznoer/csscolorparser"
)

var eps = math.Nextafter(1.0, 2.0) - 1.0

const tau = math.Pi * 2

func parseCss(s string) ([]Stop, bool) {
	stops := []Stop{}

	for _, stop := range splitByComma(s) {
		if !prosesStop(&stops, splitBySpace(stop)) {
			return stops, false
		}
	}

	if len(stops) == 0 {
		return stops, false
	}

	if stops[0].Color == nil {
		return stops, false
	}

	for i, stop := range stops {
		if i == 0 && stop.Pos == nil {
			stops[i].Pos = ptr(0.0)
		}

		if i == len(stops)-1 {
			if stop.Pos == nil {
				stops[i].Pos = ptr(1.0)
			}
			continue
		}

		if stop.Color == nil {
			stops[i].Color = interpolateColor(*stops[i-1].Color, *stops[i+1].Color, 0.5)
		}
	}

	if *stops[0].Pos > 0.0 {
		stops = append([]Stop{{ptr(0.0), stops[0].Color}}, stops...)
	}

	if *stops[len(stops)-1].Pos < 1.0 {
		stops = append(stops, Stop{ptr(1.0), stops[len(stops)-1].Color})
	}

	for i, stop := range stops {
		if stop.Pos == nil {
			for j := i + 1; j < len(stops); j++ {
				if stops[j].Pos != nil {
					prev := *stops[i-1].Pos
					next := *stops[j].Pos
					stops[i].Pos = ptr(prev + (next-prev)/float64(j-i+1))
					break
				}
			}
		}

		if i > 0 {
			stops[i].Pos = ptr(math.Max(*stops[i].Pos, *stops[i-1].Pos))
		}
	}

	// Filter Stops

	prev := stops[0].Pos
	last := len(stops) - 1
	newStops := make([]Stop, 0, len(stops))

	for i, s := range stops {
		var next *float64

		if i == last {
			next = stops[last].Pos
		} else {
			next = stops[i+1].Pos
		}

		if (*s.Pos-*prev)+(*next-*s.Pos) < eps {
			// skip 0 width stop
		} else {
			newStops = append(newStops, s)
		}
		prev = s.Pos
	}

	return newStops, true
}

func ptr(f float64) *float64 {
	return &f
}

type Stop struct {
	Pos   *float64
	Color *Color
}

func prosesStop(stops *[]Stop, arr []string) bool {
	switch len(arr) {
	case 1:
		col, err := csscolorparser.Parse(arr[0])
		if err == nil {
			*stops = append(*stops, Stop{nil, &col})
			return true
		}

		pos, ok := parsePos(arr[0])
		if ok {
			*stops = append(*stops, Stop{&pos, nil})
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

		*stops = append(*stops, Stop{&pos, &col})
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

		*stops = append(*stops, Stop{&pos1, &col})
		*stops = append(*stops, Stop{&pos2, &col})
	default:
		return false
	}
	return true
}

func interpolateColor(a, b Color, t float64) *Color {
	return &Color{
		R: a.R + t*(b.R-a.R),
		G: a.G + t*(b.G-a.G),
		B: a.B + t*(b.B-a.B),
		A: a.A + t*(b.A-a.A),
	}
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
