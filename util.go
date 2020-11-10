package colorgrad

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/colornames"
)

var reRgb, reRgba, reRgbPct, reRgbaPct, reHsl, reHsla *regexp.Regexp

func init() {
	// Regex taken from chroma.js (with some modification)
	reRgb = regexp.MustCompile(`^rgb\(\s*(-?\d+)\s*,\s*(-?\d+)\s*,\s*(-?\d+)\s*\)$`)
	reRgba = regexp.MustCompile(`^rgba\(\s*(-?\d+)\s*,\s*(-?\d+)\s*,\s*(-?\d+)\s*,\s*([01]|[01]?\.\d+)\)$`)
	reRgbPct = regexp.MustCompile(`^rgb\(\s*(-?\d+(?:\.\d+)?)%\s*,\s*(-?\d+(?:\.\d+)?)%\s*,\s*(-?\d+(?:\.\d+)?)%\s*\)$`)
	reRgbaPct = regexp.MustCompile(`^rgba\(\s*(-?\d+(?:\.\d+)?)%\s*,\s*(-?\d+(?:\.\d+)?)%\s*,\s*(-?\d+(?:\.\d+)?)%\s*,\s*([01]|[01]?\.\d+)\)$`)
	reHsl = regexp.MustCompile(`^hsl\(\s*(-?\d+(?:\.\d+)?)\s*,\s*(-?\d+(?:\.\d+)?)%\s*,\s*(-?\d+(?:\.\d+)?)%\s*\)$`)
	reHsla = regexp.MustCompile(`^hsla\(\s*(-?\d+(?:\.\d+)?)\s*,\s*(-?\d+(?:\.\d+)?)%\s*,\s*(-?\d+(?:\.\d+)?)%\s*,\s*([01]|[01]?\.\d+)\)$`)
}

func parseColor(str string) (colorful.Color, error) {
	str = strings.TrimSpace(strings.ToLower(str))

	// Predefined name / keyword
	c, ok := colornames.Map[str]
	if ok {
		col, _ := colorful.MakeColor(c)
		return col, nil
	}

	// Hexadecimal
	c2, err := colorful.Hex(str)
	if err == nil {
		return c2, nil
	}

	// RGB
	if reRgb.MatchString(str) {
		str = str[4 : len(str)-1]
		x := strings.Split(str, ",")
		return colorful.Color{
			R: parseFloat(x[0]) / 255,
			G: parseFloat(x[1]) / 255,
			B: parseFloat(x[2]) / 255,
		}, nil
	}

	// RGBA
	if reRgba.MatchString(str) {
		str = str[5 : len(str)-1]
		x := strings.Split(str, ",")
		return colorful.Color{
			R: parseFloat(x[0]) / 255,
			G: parseFloat(x[1]) / 255,
			B: parseFloat(x[2]) / 255,
		}, nil
	}

	// RGB percent
	if reRgbPct.MatchString(str) {
		str = str[4 : len(str)-1]
		x := strings.Split(str, ",")
		return colorful.Color{
			R: parseFloat(strings.Replace(x[0], "%", "", 1)) / 100,
			G: parseFloat(strings.Replace(x[1], "%", "", 1)) / 100,
			B: parseFloat(strings.Replace(x[2], "%", "", 1)) / 100,
		}, nil
	}

	// RGBA percent
	if reRgbaPct.MatchString(str) {
		str = str[5 : len(str)-1]
		x := strings.Split(str, ",")
		return colorful.Color{
			R: parseFloat(strings.Replace(x[0], "%", "", 1)) / 100,
			G: parseFloat(strings.Replace(x[1], "%", "", 1)) / 100,
			B: parseFloat(strings.Replace(x[2], "%", "", 1)) / 100,
		}, nil
	}

	// HSL
	if reHsl.MatchString(str) {
		str = str[4 : len(str)-1]
		x := strings.Split(str, ",")
		return colorful.Hsl(
			parseFloat(x[0]),
			parseFloat(strings.Replace(x[1], "%", "", 1))/100,
			parseFloat(strings.Replace(x[2], "%", "", 1))/100,
		), nil
	}

	// HSLA
	if reHsla.MatchString(str) {
		str = str[5 : len(str)-1]
		x := strings.Split(str, ",")
		return colorful.Hsl(
			parseFloat(x[0]),
			parseFloat(strings.Replace(x[1], "%", "", 1))/100,
			parseFloat(strings.Replace(x[2], "%", "", 1))/100,
		), nil
	}

	return colorful.Color{R: 0, G: 0, B: 0}, fmt.Errorf("Invalid color format")
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return f
}

func linspace(min, max float64, n uint) []float64 {
	d := max - min
	l := float64(n - 1)
	res := make([]float64, n)
	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}

// Algorithm taken from: https://github.com/gka/chroma.js

func blendLrgb(a, b colorful.Color, t float64) colorful.Color {
	return colorful.Color{
		R: math.Sqrt(math.Pow(a.R, 2)*(1-t) + math.Pow(b.R, 2)*t),
		G: math.Sqrt(math.Pow(a.G, 2)*(1-t) + math.Pow(b.G, 2)*t),
		B: math.Sqrt(math.Pow(a.B, 2)*(1-t) + math.Pow(b.B, 2)*t),
	}
}
