# colorgrad 🎨

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mazznoer/colorgrad)](https://pkg.go.dev/github.com/mazznoer/colorgrad)
[![Build Status](https://travis-ci.org/mazznoer/colorgrad.svg?branch=master)](https://travis-ci.org/mazznoer/colorgrad)
[![Build Status](https://github.com/mazznoer/colorgrad/workflows/Go/badge.svg)](https://github.com/mazznoer/colorgrad/actions)
[![go report](https://goreportcard.com/badge/github.com/mazznoer/colorgrad)](https://goreportcard.com/report/github.com/mazznoer/colorgrad)
[![codecov](https://codecov.io/gh/mazznoer/colorgrad/branch/master/graph/badge.svg)](https://codecov.io/gh/mazznoer/colorgrad)
[![Release](https://img.shields.io/github/release/mazznoer/colorgrad.svg)](https://github.com/mazznoer/colorgrad/releases/latest)

Go (Golang) _color scales_ library.

## Index

* [Usage](#usage)
    - [Basic](#basic)
    - [Custom Colors](#custom-colors)
    - [Hex Colors](#using-hex-colors)
    - [Named Colors](#named-colors)
    - [CSS Color String](#css-color-string)
    - [Domain & Color Position](#domain--color-position)
    - [Blending Mode](#blending-mode)
    - [Invalid RGB](#beware-of-invalid-rgb-color)
* [Preset Gradients](#preset-gradients)
* [Hard-Edged Gradient](#hard-edged-gradient)
* [Color Schemes](#color-schemes)
* [Gallery](#gallery)
* [Playground](#playground)
* [Dependency](#dependency)
* [Inspirations](#inspirations)

```go
import "github.com/mazznoer/colorgrad"
```

```go
type Gradient interface {
    // Get color at certain position
    At(float64) colorful.Color

    // Get n colors evenly spaced across gradient
    ColorfulColors(uint) []colorful.Color

    // Get n colors evenly spaced across gradient
    Colors(uint) []color.Color

    // Get the gradient domain min and max
    Domain() (float64, float64)

    // Return a new hard-edge gradient
    Sharp(uint) Gradient
}
```

## Usage

### Basic

```go
grad, err := colorgrad.NewGradient().Build()

if err != nil {
    panic(err)
}
```
![img](doc/images/custom-default.png)

### Custom Colors

`Colors()` method accept anything that implement [color.Color](https://golang.org/pkg/image/color/#Color) interface.

```go
import "image/color"
import "github.com/lucasb-eyer/go-colorful"

grad, err := colorgrad.NewGradient().
    Colors(
        color.RGBA{0, 206, 209, 255},
        color.RGBA{255, 105, 180, 255},
        colorful.Color{R: 0.274, G: 0.5, B: 0.7},
        colorful.Hsv(50, 1, 1),
        colorful.Hsv(348, 0.9, 0.8),
    ).
    Build()
```
![img](doc/images/custom-colors.png)

### Using Hex Colors

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#c41189", "#00BFFF", "#FFD700").
    Build()
```
![img](doc/images/custom-hex-colors.png)

### Named Colors

We can also use named colors as defined in the SVG 1.1 spec.

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("gold", "hotpink", "darkturquoise").
    Build()
```
![img](doc/images/custom-named-colors.png)

### CSS Color String

```go
grad, err := colorgrad.NewGradient().
    HtmlColors(
        "rgb(125,110,221)",
        "rgb(90%,45%,97%)",
        "hsl(229,79%,85%)",
    ).
    Build()
```
![img](doc/images/custom-css-colors.png)

### Domain & Color Position

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("deeppink", "gold", "seagreen").
    Build()
```
![img](doc/images/domain-default.png)

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("deeppink", "gold", "seagreen").
    Domain(0, 100).
    Build()
```
![img](doc/images/domain-100.png)

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("deeppink", "gold", "seagreen").
    Domain(0, 0.7, 1).
    Build()
```
![img](doc/images/color-position-1.png)

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("deeppink", "gold", "seagreen").
    Domain(15, 30, 80).
    Build()
```
![img](doc/images/color-position-2.png)

### Blending Mode

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#ff0", "#008ae5").
    Mode(colorgrad.LRGB).
    Build()
```
![blend-modes](doc/images/blend-modes.png)

### Beware of Invalid RGB Color
Read it [here](https://github.com/lucasb-eyer/go-colorful#blending-colors).

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#DC143C", "#FFD700", "#4682b4").
    Mode(colorgrad.HCL).
    Build()

grad.At(t) // might get invalid RGB color
grad.At(t).Clamped() // return closest valid RGB color
```

Without `Clamped()`

![invalid rgb](doc/images/custom-invalid-colors.png)

With `Clamped()`

![valid rgb](doc/images/custom-clamped.png)

## Preset Gradients

All preset gradients are in the domain 0..1.

### Diverging

`colorgrad.BrBG()`
![img](doc/images/preset/BrBG.png)

`colorgrad.PRGn()`
![img](doc/images/preset/PRGn.png)

`colorgrad.PiYG()`
![img](doc/images/preset/PiYG.png)

`colorgrad.PuOr()`
![img](doc/images/preset/PuOr.png)

`colorgrad.RdBu()`
![img](doc/images/preset/RdBu.png)

`colorgrad.RdGy()`
![img](doc/images/preset/RdGy.png)

`colorgrad.RdYlBu()`
![img](doc/images/preset/RdYlBu.png)

`colorgrad.RdYlGn()`
![img](doc/images/preset/RdYlGn.png)

`colorgrad.Spectral()`
![img](doc/images/preset/Spectral.png)

### Sequential (Single Hue)

`colorgrad.Blues()`
![img](doc/images/preset/Blues.png)

`colorgrad.Greens()`
![img](doc/images/preset/Greens.png)

`colorgrad.Greys()`
![img](doc/images/preset/Greys.png)

`colorgrad.Oranges()`
![img](doc/images/preset/Oranges.png)

`colorgrad.Purples()`
![img](doc/images/preset/Purples.png)

`colorgrad.Reds()`
![img](doc/images/preset/Reds.png)

### Sequential (Multi-Hue)

`colorgrad.Turbo()`
![img](doc/images/preset/Turbo.png)

`colorgrad.Viridis()`
![img](doc/images/preset/Viridis.png)

`colorgrad.Inferno()`
![img](doc/images/preset/Inferno.png)

`colorgrad.Magma()`
![img](doc/images/preset/Magma.png)

`colorgrad.Plasma()`
![img](doc/images/preset/Plasma.png)

`colorgrad.Cividis()`
![img](doc/images/preset/Cividis.png)

`colorgrad.Warm()`
![img](doc/images/preset/Warm.png)

`colorgrad.Cool()`
![img](doc/images/preset/Cool.png)

`colorgrad.CubehelixDefault()`
![img](doc/images/preset/CubehelixDefault.png)

### Cyclical

`colorgrad.Rainbow()`
![img](doc/images/preset/Rainbow.png)

`colorgrad.Sinebow()`
![img](doc/images/preset/Sinebow.png)

## Hard-Edged Gradient

```go
grad1, err := colorgrad.NewGradient().
    HtmlColors("#18dbf4", "#f6ff56").
    Build()
```
![img](doc/images/gradient-normal.png)

```go
grad2 := grad1.Sharp(7)
```
![img](doc/images/gradient-sharp.png)

```go
grad := colorgrad.Spectral().Sharp(19)
```
![img](doc/images/spectral-sharp.png)

## Color Schemes

```go
import "github.com/mazznoer/colorgrad/scheme"
```

`scheme.Category10`

![img](doc/images/scheme/Category10.png)

`scheme.Accent`

![img](doc/images/scheme/Accent.png)

`scheme.Dark2`

![img](doc/images/scheme/Dark2.png)

`scheme.Paired`

![img](doc/images/scheme/Paired.png)

`scheme.Pastel1`

![img](doc/images/scheme/Pastel1.png)

`scheme.Pastel2`

![img](doc/images/scheme/Pastel2.png)

`scheme.Set1`

![img](doc/images/scheme/Set1.png)

`scheme.Set2`

![img](doc/images/scheme/Set2.png)

`scheme.Set3`

![img](doc/images/scheme/Set3.png)

`scheme.Tableau10`

![img](doc/images/scheme/Tableau10.png)

## Gallery

Noise + hard-edged gradient

![noise](doc/images/noise.png)

Random colors from `colorgrad.Cool()`

![random-colors](doc/images/random-colors.png)

## Playground

* [Basic](https://play.golang.org/p/qoUQvzOkceg)
* [Random colors](https://play.golang.org/p/d67x9di4sAF)

## Dependency

* [colorful](https://github.com/lucasb-eyer/go-colorful)

## Inspirations

* [chroma.js](https://github.com/gka/chroma.js)
* [d3-scale-chromatic](https://github.com/d3/d3-scale-chromatic/)
* colorful's [gradientgen.go](https://github.com/lucasb-eyer/go-colorful/blob/master/doc/gradientgen/gradientgen.go)

