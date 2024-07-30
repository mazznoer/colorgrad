package colorgrad

import (
	"math"

	"github.com/mazznoer/csscolorparser"
)

// Reference: https://github.com/d3/d3-scale-chromatic

const deg2rad = math.Pi / 180
const pi1_3 = math.Pi / 3
const pi2_3 = math.Pi * 2 / 3

// Sinebow

type sinebowGradient struct{}

func Sinebow() Gradient {
	return Gradient{
		grad: sinebowGradient{},
		dmin: 0,
		dmax: 1,
	}
}

func (sg sinebowGradient) At(t float64) Color {
	t = (0.5 - t) * math.Pi
	return Color{
		R: math.Pow(math.Sin(t), 2),
		G: math.Pow(math.Sin(t+pi1_3), 2),
		B: math.Pow(math.Sin(t+pi2_3), 2),
		A: 1,
	}
}

// Turbo

type turboGradient struct{}

func Turbo() Gradient {
	return Gradient{
		grad: turboGradient{},
		dmin: 0,
		dmax: 1,
	}
}

func (tg turboGradient) At(t float64) Color {
	t = math.Max(0, math.Min(1, t))
	r := math.Round(34.61 + t*(1172.33-t*(10793.56-t*(33300.12-t*(38394.49-t*14825.05)))))
	g := math.Round(23.31 + t*(557.33+t*(1225.33-t*(3574.96-t*(1073.77+t*707.56)))))
	b := math.Round(27.2 + t*(3211.1-t*(15327.97-t*(27814-t*(22569.18-t*6838.66)))))
	return Color{
		R: clamp01(r / 255),
		G: clamp01(g / 255),
		B: clamp01(b / 255),
		A: 1,
	}
}

// Cividis

type cividisGradient struct{}

func Cividis() Gradient {
	return Gradient{
		grad: cividisGradient{},
		dmin: 0,
		dmax: 1,
	}
}

func (cg cividisGradient) At(t float64) Color {
	t = math.Max(0, math.Min(1, t))
	r := math.Round(-4.54 - t*(35.34-t*(2381.73-t*(6402.7-t*(7024.72-t*2710.57)))))
	g := math.Round(32.49 + t*(170.73+t*(52.82-t*(131.46-t*(176.58-t*67.37)))))
	b := math.Round(81.24 + t*(442.36-t*(2482.43-t*(6167.24-t*(6614.94-t*2475.67)))))
	return Color{
		R: clamp01(r / 255),
		G: clamp01(g / 255),
		B: clamp01(b / 255),
		A: 1,
	}
}

// Cubehelix

type cubehelix struct {
	h, s, l float64
}

func (c cubehelix) toColor() Color {
	h := (c.h + 120) * deg2rad
	l := c.l
	a := c.s * l * (1 - l)
	cosh := math.Cos(h)
	sinh := math.Sin(h)
	r := (l - a*math.Min(0.14861*cosh-1.78277*sinh, 1.0))
	g := (l - a*math.Min(0.29227*cosh+0.90649*sinh, 1.0))
	b := l + a*(1.97294*cosh)
	return Color{
		R: clamp01(r),
		G: clamp01(g),
		B: clamp01(b),
		A: 1,
	}
}

func (c cubehelix) interpolate(c2 cubehelix, t float64) cubehelix {
	return cubehelix{
		h: c.h + t*(c2.h-c.h),
		s: c.s + t*(c2.s-c.s),
		l: c.l + t*(c2.l-c.l),
	}
}

// Cubehelix gradient

type cubehelixGradient struct {
	start, end cubehelix
}

func CubehelixDefault() Gradient {
	gradbase := cubehelixGradient{
		start: cubehelix{300, 0.5, 0.0},
		end:   cubehelix{-240, 0.5, 1.0},
	}
	return Gradient{
		grad: gradbase,
		dmin: 0,
		dmax: 1,
	}
}

func Warm() Gradient {
	gradbase := cubehelixGradient{
		start: cubehelix{-100, 0.75, 0.35},
		end:   cubehelix{80, 1.50, 0.8},
	}
	return Gradient{
		grad: gradbase,
		dmin: 0,
		dmax: 1,
	}
}

func Cool() Gradient {
	gradbase := cubehelixGradient{
		start: cubehelix{260, 0.75, 0.35},
		end:   cubehelix{80, 1.50, 0.8},
	}
	return Gradient{
		grad: gradbase,
		dmin: 0,
		dmax: 1,
	}
}

func (cg cubehelixGradient) At(t float64) Color {
	return cg.start.interpolate(cg.end, clamp01(t)).toColor()
}

// Rainbow

type rainbowGradient struct{}

func Rainbow() Gradient {
	return Gradient{
		grad: rainbowGradient{},
		dmin: 0,
		dmax: 1,
	}
}

func (rg rainbowGradient) At(t float64) Color {
	t = math.Max(0, math.Min(1, t))
	ts := math.Abs(t - 0.5)
	return cubehelix{
		h: 360*t - 100,
		s: 1.5 - 1.5*ts,
		l: 0.8 - 0.9*ts,
	}.toColor()
}

// --- Presets from color ramps

func preset(htmlColors []string) Gradient {
	var colors []Color
	for _, s := range htmlColors {
		c, err := csscolorparser.Parse(s)
		if err == nil {
			colors = append(colors, c)
		}
	}
	pos := linspace(0, 1, uint(len(colors)))
	return newBasisGradient(colors, pos, BlendRgb)
}

// Diverging

func BrBG() Gradient {
	colors := []string{"#543005", "#8c510a", "#bf812d", "#dfc27d", "#f6e8c3", "#f5f5f5", "#c7eae5", "#80cdc1", "#35978f", "#01665e", "#003c30"}
	return preset(colors)
}

func PRGn() Gradient {
	colors := []string{"#40004b", "#762a83", "#9970ab", "#c2a5cf", "#e7d4e8", "#f7f7f7", "#d9f0d3", "#a6dba0", "#5aae61", "#1b7837", "#00441b"}
	return preset(colors)
}

func PiYG() Gradient {
	colors := []string{"#8e0152", "#c51b7d", "#de77ae", "#f1b6da", "#fde0ef", "#f7f7f7", "#e6f5d0", "#b8e186", "#7fbc41", "#4d9221", "#276419"}
	return preset(colors)
}

func PuOr() Gradient {
	colors := []string{"#2d004b", "#542788", "#8073ac", "#b2abd2", "#d8daeb", "#f7f7f7", "#fee0b6", "#fdb863", "#e08214", "#b35806", "#7f3b08"}
	return preset(colors)
}

func RdBu() Gradient {
	colors := []string{"#67001f", "#b2182b", "#d6604d", "#f4a582", "#fddbc7", "#f7f7f7", "#d1e5f0", "#92c5de", "#4393c3", "#2166ac", "#053061"}
	return preset(colors)
}

func RdGy() Gradient {
	colors := []string{"#67001f", "#b2182b", "#d6604d", "#f4a582", "#fddbc7", "#ffffff", "#e0e0e0", "#bababa", "#878787", "#4d4d4d", "#1a1a1a"}
	return preset(colors)
}

func RdYlBu() Gradient {
	colors := []string{"#a50026", "#d73027", "#f46d43", "#fdae61", "#fee090", "#ffffbf", "#e0f3f8", "#abd9e9", "#74add1", "#4575b4", "#313695"}
	return preset(colors)
}

func RdYlGn() Gradient {
	colors := []string{"#a50026", "#d73027", "#f46d43", "#fdae61", "#fee08b", "#ffffbf", "#d9ef8b", "#a6d96a", "#66bd63", "#1a9850", "#006837"}
	return preset(colors)
}

func Spectral() Gradient {
	colors := []string{"#9e0142", "#d53e4f", "#f46d43", "#fdae61", "#fee08b", "#ffffbf", "#e6f598", "#abdda4", "#66c2a5", "#3288bd", "#5e4fa2"}
	return preset(colors)
}

// Sequential (Single Hue)

func Blues() Gradient {
	colors := []string{"#f7fbff", "#deebf7", "#c6dbef", "#9ecae1", "#6baed6", "#4292c6", "#2171b5", "#08519c", "#08306b"}
	return preset(colors)
}

func Greens() Gradient {
	colors := []string{"#f7fcf5", "#e5f5e0", "#c7e9c0", "#a1d99b", "#74c476", "#41ab5d", "#238b45", "#006d2c", "#00441b"}
	return preset(colors)
}

func Greys() Gradient {
	colors := []string{"#ffffff", "#f0f0f0", "#d9d9d9", "#bdbdbd", "#969696", "#737373", "#525252", "#252525", "#000000"}
	return preset(colors)
}

func Oranges() Gradient {
	colors := []string{"#fff5eb", "#fee6ce", "#fdd0a2", "#fdae6b", "#fd8d3c", "#f16913", "#d94801", "#a63603", "#7f2704"}
	return preset(colors)
}

func Purples() Gradient {
	colors := []string{"#fcfbfd", "#efedf5", "#dadaeb", "#bcbddc", "#9e9ac8", "#807dba", "#6a51a3", "#54278f", "#3f007d"}
	return preset(colors)
}

func Reds() Gradient {
	colors := []string{"#fff5f0", "#fee0d2", "#fcbba1", "#fc9272", "#fb6a4a", "#ef3b2c", "#cb181d", "#a50f15", "#67000d"}
	return preset(colors)
}

// Sequential (Multi-Hue)

func Viridis() Gradient {
	colors := []string{"#440154", "#482777", "#3f4a8a", "#31678e", "#26838f", "#1f9d8a", "#6cce5a", "#b6de2b", "#fee825"}
	return preset(colors)
}

func Inferno() Gradient {
	colors := []string{"#000004", "#170b3a", "#420a68", "#6b176e", "#932667", "#bb3654", "#dd513a", "#f3771a", "#fca50a", "#f6d644", "#fcffa4"}
	return preset(colors)
}

func Magma() Gradient {
	colors := []string{"#000004", "#140e37", "#3b0f70", "#641a80", "#8c2981", "#b63679", "#de4968", "#f66f5c", "#fe9f6d", "#fece91", "#fcfdbf"}
	return preset(colors)
}

func Plasma() Gradient {
	colors := []string{"#0d0887", "#42039d", "#6a00a8", "#900da3", "#b12a90", "#cb4678", "#e16462", "#f1834b", "#fca636", "#fccd25", "#f0f921"}
	return preset(colors)
}

func BuGn() Gradient {
	colors := []string{"#f7fcfd", "#e5f5f9", "#ccece6", "#99d8c9", "#66c2a4", "#41ae76", "#238b45", "#006d2c", "#00441b"}
	return preset(colors)
}

func BuPu() Gradient {
	colors := []string{"#f7fcfd", "#e0ecf4", "#bfd3e6", "#9ebcda", "#8c96c6", "#8c6bb1", "#88419d", "#810f7c", "#4d004b"}
	return preset(colors)
}

func GnBu() Gradient {
	colors := []string{"#f7fcf0", "#e0f3db", "#ccebc5", "#a8ddb5", "#7bccc4", "#4eb3d3", "#2b8cbe", "#0868ac", "#084081"}
	return preset(colors)
}

func OrRd() Gradient {
	colors := []string{"#fff7ec", "#fee8c8", "#fdd49e", "#fdbb84", "#fc8d59", "#ef6548", "#d7301f", "#b30000", "#7f0000"}
	return preset(colors)
}

func PuBuGn() Gradient {
	colors := []string{"#fff7fb", "#ece2f0", "#d0d1e6", "#a6bddb", "#67a9cf", "#3690c0", "#02818a", "#016c59", "#014636"}
	return preset(colors)
}

func PuBu() Gradient {
	colors := []string{"#fff7fb", "#ece7f2", "#d0d1e6", "#a6bddb", "#74a9cf", "#3690c0", "#0570b0", "#045a8d", "#023858"}
	return preset(colors)
}

func PuRd() Gradient {
	colors := []string{"#f7f4f9", "#e7e1ef", "#d4b9da", "#c994c7", "#df65b0", "#e7298a", "#ce1256", "#980043", "#67001f"}
	return preset(colors)
}

func RdPu() Gradient {
	colors := []string{"#fff7f3", "#fde0dd", "#fcc5c0", "#fa9fb5", "#f768a1", "#dd3497", "#ae017e", "#7a0177", "#49006a"}
	return preset(colors)
}

func YlGnBu() Gradient {
	colors := []string{"#ffffd9", "#edf8b1", "#c7e9b4", "#7fcdbb", "#41b6c4", "#1d91c0", "#225ea8", "#253494", "#081d58"}
	return preset(colors)
}

func YlGn() Gradient {
	colors := []string{"#ffffe5", "#f7fcb9", "#d9f0a3", "#addd8e", "#78c679", "#41ab5d", "#238443", "#006837", "#004529"}
	return preset(colors)
}

func YlOrBr() Gradient {
	colors := []string{"#ffffe5", "#fff7bc", "#fee391", "#fec44f", "#fe9929", "#ec7014", "#cc4c02", "#993404", "#662506"}
	return preset(colors)
}

func YlOrRd() Gradient {
	colors := []string{"#ffffcc", "#ffeda0", "#fed976", "#feb24c", "#fd8d3c", "#fc4e2a", "#e31a1c", "#bd0026", "#800026"}
	return preset(colors)
}
