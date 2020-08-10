# colorgrad 🎨

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/colorgrad?tab=doc)
[![go report](https://goreportcard.com/badge/github.com/mazznoer/colorgrad)](https://goreportcard.com/report/github.com/mazznoer/colorgrad)
[![Build Status](https://travis-ci.org/mazznoer/colorgrad.svg?branch=master)](https://travis-ci.org/mazznoer/colorgrad)
[![codecov](https://codecov.io/gh/mazznoer/colorgrad/branch/master/graph/badge.svg)](https://codecov.io/gh/mazznoer/colorgrad)

Fun & easy way to create _color gradient_ / _color scales_ in __Go__ (__Golang__).

![color-scale](img/color-scale-1.png)
<!--
## Index

* [Usages](#usages)
  - [Basic](#basic)
  - [Custom Colors](#custom-colors)
  - [Hex Colors](#using-hex-colors)
  - [Named Colors](#named-colors)
  - [Custom Domain](#custom-domain)
  - [Blending Mode](#blending-mode)
  - [Invalid RGB](#beware-of-invalid-rgb-color)
  - [Hard-Edged Gradient](#hard-edged-gradient)
* [Preset Gradients](#preset-gradients)
* [Color Scheme](#color-scheme)
* [Gallery](#gallery)
* [Playground](#playground)
* [Dependencies](#dependencies)
* [Inspirations](#inspirations)
-->
```go
import "github.com/mazznoer/colorgrad"
```

## Preset Gradients
<!--
```go
grad := colorgrad.Rainbow()

c := grad.At(t) // t in the range 0..1
colors1 := grad.ColorfulColors(5)
colors2 := grad.Colors(17)
grad2 := grad.Sharp(13)
```
-->
#### Diverging

`colorgrad.BrBG()`
![img](img/gradient-BrBG.png)

`colorgrad.PRGn()`
![img](img/gradient-PRGn.png)

`colorgrad.PiYG()`
![img](img/gradient-PiYG.png)

`colorgrad.PuOr()`
![img](img/gradient-PuOr.png)

`colorgrad.RdBu()`
![img](img/gradient-RdBu.png)

`colorgrad.RdGy()`
![img](img/gradient-RdGy.png)

`colorgrad.RdYlBu()`
![img](img/gradient-RdYlBu.png)

`colorgrad.RdYlGn()`
![img](img/gradient-RdYlGn.png)

`colorgrad.Spectral()`
![img](img/gradient-spectral.png)

#### Sequential (Single Hue)

`colorgrad.Blues()`
![img](img/gradient-blues.png)

`colorgrad.Greens()`
![img](img/gradient-greens.png)

`colorgrad.Greys()`
![img](img/gradient-greys.png)

`colorgrad.Oranges()`
![img](img/gradient-oranges.png)

`colorgrad.Purples()`
![img](img/gradient-purples.png)

`colorgrad.Reds()`
![img](img/gradient-reds.png)

#### Sequential (Multi-Hue)

`colorgrad.Turbo()`
![img](img/gradient-turbo.png)

`colorgrad.Viridis()`
![img](img/gradient-viridis.png)

`colorgrad.Inferno()`
![img](img/gradient-inferno.png)

`colorgrad.Magma()`
![img](img/gradient-magma.png)

`colorgrad.Plasma()`
![img](img/gradient-plasma.png)

`colorgrad.Cividis()`
![img](img/gradient-cividis.png)

`colorgrad.Warm()`
![img](img/gradient-warm.png)

`colorgrad.Cool()`
![img](img/gradient-cool.png)

`colorgrad.CubehelixDefault()`
![img](img/gradient-cubehelix-default.png)

#### Cyclical

`colorgrad.Rainbow()`
![img](img/gradient-rainbow.png)

`colorgrad.Sinebow()`
![img](img/gradient-sinebow.png)

## Custom Gradient

#### Basic

```go
grad, err := colorgrad.NewGradient().Build()

if err != nil {
    panic(err)
}

dmin, dmax := grad.Domain()

// Get single color at certain position.
// t in the range dmin..dmax (default to 0..1)
c1 := grad.At(t)       // colorful.Color
c2 := grad.At(t).Hex() // hex color string
var c3 color.Color = grad.At(t) // color.Color

// Get n colors evenly spaced across gradient.
colors1 := grad.ColorfulColors(9) // []colorful.Color
colors2 := grad.Colors(23)        // []color.Color
```
![img](img/black-to-white.png)

#### Custom Colors

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
    Mode(colorgrad.HCL).
    Build()
```
![img](img/custom-colors.png)

#### Using Hex Colors

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#FFD700", "#00BFFF", "#FFD700").
    Build()
```
![img](img/hex-colors.png)

#### Named Colors

We can also use named colors as defined in the SVG 1.1 spec.

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("gold", "hotpink", "darkturquoise").
    Build()
```
![img](img/named-colors.png)

#### Custom Domain & Color Position

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#DC143C", "#FFD700", "#4682b4").
    Build()
```
![img](img/color-scale-1.png)

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#DC143C", "#FFD700", "#4682b4").
    Domain(0, 0.35, 1).
    Build()
```
![img](img/color-scale-2.png)

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#DC143C", "#FFD700", "#4682b4").
    Domain(15, 60, 80).
    Build()

grad.At(15).Hex() // #DC143C
grad.At(75)
grad.At(80).Hex() // #4682b4
```
![img](img/color-scale-3.png)

#### Blending Mode

```go
grad, err := colorgrad.NewGradient().
    HtmlColors("#ff0", "#008ae5").
    Mode(colorgrad.LRGB).
    Build()
```
![blend-modes](img/blend-modes.png)

#### Beware of Invalid RGB Color
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
![invalig rgb](img/not-clamped.png)

With `Clamped()`
![valid rgb](img/clamped.png)

## Gradient Methods

_TODO_

#### Hard-Edged Gradient

```go
grad1, err := colorgrad.NewGradient().
    HtmlColors("#18dbf4", "#f6ff56").
    Domain(0, 100).
    Build()
```
![img](img/normal-gradient.png)

```go
grad2 := grad1.Sharp(7)

dmin, dmax := grad2.Domain() // 0, 100 -- same as original gradient
```
![img](img/classes-gradient.png)

## Color Scheme

`colorgrad.Scheme.Accent`
![img](img/scheme-accent.png)

`colorgrad.Scheme.Category10`
![img](img/scheme-category10.png)

`colorgrad.Scheme.Dark2`
![img](img/scheme-dark2.png)

`colorgrad.Scheme.Paired`
![img](img/scheme-paired.png)

`colorgrad.Scheme.Pastel1`
![img](img/scheme-pastel1.png)

`colorgrad.Scheme.Pastel2`
![img](img/scheme-pastel2.png)

`colorgrad.Scheme.Set1`
![img](img/scheme-set1.png)

`colorgrad.Scheme.Set2`
![img](img/scheme-set2.png)

`colorgrad.Scheme.Set3`
![img](img/scheme-set3.png)

## Gallery

Noise + hard-edged gradient
![noise](img/noise-2.png)

Random _cool_ colors
![random-color](img/random-cool.png)

## Playground

* [Basic](https://play.golang.org/p/qoUQvzOkceg)
* [Random colors](https://play.golang.org/p/d67x9di4sAF)

## Dependencies

* [colorful](https://github.com/lucasb-eyer/go-colorful)

## Inspirations

* [chroma.js](https://github.com/gka/chroma.js)
* [d3-scale-chromatic](https://github.com/d3/d3-scale-chromatic/)
