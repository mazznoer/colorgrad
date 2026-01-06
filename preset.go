package colorgrad

import (
	"math"
)

// Reference: https://github.com/d3/d3-scale-chromatic

const deg2rad = math.Pi / 180
const pi1_3 = math.Pi / 3
const pi2_3 = math.Pi * 2 / 3

// Sinebow

type sinebowGradient struct{}

func Sinebow() Gradient {
	return Gradient{
		Core: sinebowGradient{},
		Min:  0,
		Max:  1,
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
		Core: turboGradient{},
		Min:  0,
		Max:  1,
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
		Core: cividisGradient{},
		Min:  0,
		Max:  1,
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
		Core: gradbase,
		Min:  0,
		Max:  1,
	}
}

func Warm() Gradient {
	gradbase := cubehelixGradient{
		start: cubehelix{-100, 0.75, 0.35},
		end:   cubehelix{80, 1.50, 0.8},
	}
	return Gradient{
		Core: gradbase,
		Min:  0,
		Max:  1,
	}
}

func Cool() Gradient {
	gradbase := cubehelixGradient{
		start: cubehelix{260, 0.75, 0.35},
		end:   cubehelix{80, 1.50, 0.8},
	}
	return Gradient{
		Core: gradbase,
		Min:  0,
		Max:  1,
	}
}

func (cg cubehelixGradient) At(t float64) Color {
	return cg.start.interpolate(cg.end, clamp01(t)).toColor()
}

// Rainbow

type rainbowGradient struct{}

func Rainbow() Gradient {
	return Gradient{
		Core: rainbowGradient{},
		Min:  0,
		Max:  1,
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

func u32ToColor(v uint32) Color {
	r := uint8(v >> 16)
	g := uint8(v >> 8)
	b := uint8(v)
	return Rgb8(r, g, b, 255)
}

func preset(data []uint32) Gradient {
	colors := make([]Color, len(data))
	for i, v := range data {
		colors[i] = u32ToColor(v)
	}
	pos := linspace(0, 1, uint(len(colors)))
	return newBasisGradient(colors, pos, BlendRgb)
}

// Diverging

func BrBG() Gradient {
	colors := []uint32{0x543005, 0x8c510a, 0xbf812d, 0xdfc27d, 0xf6e8c3, 0xf5f5f5, 0xc7eae5, 0x80cdc1, 0x35978f, 0x01665e, 0x003c30}
	return preset(colors)
}

func PRGn() Gradient {
	colors := []uint32{0x40004b, 0x762a83, 0x9970ab, 0xc2a5cf, 0xe7d4e8, 0xf7f7f7, 0xd9f0d3, 0xa6dba0, 0x5aae61, 0x1b7837, 0x00441b}
	return preset(colors)
}

func PiYG() Gradient {
	colors := []uint32{0x8e0152, 0xc51b7d, 0xde77ae, 0xf1b6da, 0xfde0ef, 0xf7f7f7, 0xe6f5d0, 0xb8e186, 0x7fbc41, 0x4d9221, 0x276419}
	return preset(colors)
}

func PuOr() Gradient {
	colors := []uint32{0x2d004b, 0x542788, 0x8073ac, 0xb2abd2, 0xd8daeb, 0xf7f7f7, 0xfee0b6, 0xfdb863, 0xe08214, 0xb35806, 0x7f3b08}
	return preset(colors)
}

func RdBu() Gradient {
	colors := []uint32{0x67001f, 0xb2182b, 0xd6604d, 0xf4a582, 0xfddbc7, 0xf7f7f7, 0xd1e5f0, 0x92c5de, 0x4393c3, 0x2166ac, 0x053061}
	return preset(colors)
}

func RdGy() Gradient {
	colors := []uint32{0x67001f, 0xb2182b, 0xd6604d, 0xf4a582, 0xfddbc7, 0xffffff, 0xe0e0e0, 0xbababa, 0x878787, 0x4d4d4d, 0x1a1a1a}
	return preset(colors)
}

func RdYlBu() Gradient {
	colors := []uint32{0xa50026, 0xd73027, 0xf46d43, 0xfdae61, 0xfee090, 0xffffbf, 0xe0f3f8, 0xabd9e9, 0x74add1, 0x4575b4, 0x313695}
	return preset(colors)
}

func RdYlGn() Gradient {
	colors := []uint32{0xa50026, 0xd73027, 0xf46d43, 0xfdae61, 0xfee08b, 0xffffbf, 0xd9ef8b, 0xa6d96a, 0x66bd63, 0x1a9850, 0x006837}
	return preset(colors)
}

func Spectral() Gradient {
	colors := []uint32{0x9e0142, 0xd53e4f, 0xf46d43, 0xfdae61, 0xfee08b, 0xffffbf, 0xe6f598, 0xabdda4, 0x66c2a5, 0x3288bd, 0x5e4fa2}
	return preset(colors)
}

// Sequential (Single Hue)

func Blues() Gradient {
	colors := []uint32{0xf7fbff, 0xdeebf7, 0xc6dbef, 0x9ecae1, 0x6baed6, 0x4292c6, 0x2171b5, 0x08519c, 0x08306b}
	return preset(colors)
}

func Greens() Gradient {
	colors := []uint32{0xf7fcf5, 0xe5f5e0, 0xc7e9c0, 0xa1d99b, 0x74c476, 0x41ab5d, 0x238b45, 0x006d2c, 0x00441b}
	return preset(colors)
}

func Greys() Gradient {
	colors := []uint32{0xffffff, 0xf0f0f0, 0xd9d9d9, 0xbdbdbd, 0x969696, 0x737373, 0x525252, 0x252525, 0x000000}
	return preset(colors)
}

func Oranges() Gradient {
	colors := []uint32{0xfff5eb, 0xfee6ce, 0xfdd0a2, 0xfdae6b, 0xfd8d3c, 0xf16913, 0xd94801, 0xa63603, 0x7f2704}
	return preset(colors)
}

func Purples() Gradient {
	colors := []uint32{0xfcfbfd, 0xefedf5, 0xdadaeb, 0xbcbddc, 0x9e9ac8, 0x807dba, 0x6a51a3, 0x54278f, 0x3f007d}
	return preset(colors)
}

func Reds() Gradient {
	colors := []uint32{0xfff5f0, 0xfee0d2, 0xfcbba1, 0xfc9272, 0xfb6a4a, 0xef3b2c, 0xcb181d, 0xa50f15, 0x67000d}
	return preset(colors)
}

// Sequential (Multi-Hue)

func Viridis() Gradient {
	colors := []uint32{0x440154, 0x482777, 0x3f4a8a, 0x31678e, 0x26838f, 0x1f9d8a, 0x6cce5a, 0xb6de2b, 0xfee825}
	return preset(colors)
}

func Inferno() Gradient {
	colors := []uint32{0x000004, 0x170b3a, 0x420a68, 0x6b176e, 0x932667, 0xbb3654, 0xdd513a, 0xf3771a, 0xfca50a, 0xf6d644, 0xfcffa4}
	return preset(colors)
}

func Magma() Gradient {
	colors := []uint32{0x000004, 0x140e37, 0x3b0f70, 0x641a80, 0x8c2981, 0xb63679, 0xde4968, 0xf66f5c, 0xfe9f6d, 0xfece91, 0xfcfdbf}
	return preset(colors)
}

func Plasma() Gradient {
	colors := []uint32{0x0d0887, 0x42039d, 0x6a00a8, 0x900da3, 0xb12a90, 0xcb4678, 0xe16462, 0xf1834b, 0xfca636, 0xfccd25, 0xf0f921}
	return preset(colors)
}

func BuGn() Gradient {
	colors := []uint32{0xf7fcfd, 0xe5f5f9, 0xccece6, 0x99d8c9, 0x66c2a4, 0x41ae76, 0x238b45, 0x006d2c, 0x00441b}
	return preset(colors)
}

func BuPu() Gradient {
	colors := []uint32{0xf7fcfd, 0xe0ecf4, 0xbfd3e6, 0x9ebcda, 0x8c96c6, 0x8c6bb1, 0x88419d, 0x810f7c, 0x4d004b}
	return preset(colors)
}

func GnBu() Gradient {
	colors := []uint32{0xf7fcf0, 0xe0f3db, 0xccebc5, 0xa8ddb5, 0x7bccc4, 0x4eb3d3, 0x2b8cbe, 0x0868ac, 0x084081}
	return preset(colors)
}

func OrRd() Gradient {
	colors := []uint32{0xfff7ec, 0xfee8c8, 0xfdd49e, 0xfdbb84, 0xfc8d59, 0xef6548, 0xd7301f, 0xb30000, 0x7f0000}
	return preset(colors)
}

func PuBuGn() Gradient {
	colors := []uint32{0xfff7fb, 0xece2f0, 0xd0d1e6, 0xa6bddb, 0x67a9cf, 0x3690c0, 0x02818a, 0x016c59, 0x014636}
	return preset(colors)
}

func PuBu() Gradient {
	colors := []uint32{0xfff7fb, 0xece7f2, 0xd0d1e6, 0xa6bddb, 0x74a9cf, 0x3690c0, 0x0570b0, 0x045a8d, 0x023858}
	return preset(colors)
}

func PuRd() Gradient {
	colors := []uint32{0xf7f4f9, 0xe7e1ef, 0xd4b9da, 0xc994c7, 0xdf65b0, 0xe7298a, 0xce1256, 0x980043, 0x67001f}
	return preset(colors)
}

func RdPu() Gradient {
	colors := []uint32{0xfff7f3, 0xfde0dd, 0xfcc5c0, 0xfa9fb5, 0xf768a1, 0xdd3497, 0xae017e, 0x7a0177, 0x49006a}
	return preset(colors)
}

func YlGnBu() Gradient {
	colors := []uint32{0xffffd9, 0xedf8b1, 0xc7e9b4, 0x7fcdbb, 0x41b6c4, 0x1d91c0, 0x225ea8, 0x253494, 0x081d58}
	return preset(colors)
}

func YlGn() Gradient {
	colors := []uint32{0xffffe5, 0xf7fcb9, 0xd9f0a3, 0xaddd8e, 0x78c679, 0x41ab5d, 0x238443, 0x006837, 0x004529}
	return preset(colors)
}

func YlOrBr() Gradient {
	colors := []uint32{0xffffe5, 0xfff7bc, 0xfee391, 0xfec44f, 0xfe9929, 0xec7014, 0xcc4c02, 0x993404, 0x662506}
	return preset(colors)
}

func YlOrRd() Gradient {
	colors := []uint32{0xffffcc, 0xffeda0, 0xfed976, 0xfeb24c, 0xfd8d3c, 0xfc4e2a, 0xe31a1c, 0xbd0026, 0x800026}
	return preset(colors)
}
