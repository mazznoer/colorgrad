# colorgrad 🎨

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/colorgrad?tab=doc)
[![go report](https://goreportcard.com/badge/github.com/mazznoer/colorgrad)](https://goreportcard.com/report/github.com/mazznoer/colorgrad)
[![Build Status](https://travis-ci.org/mazznoer/colorgrad.svg?branch=master)](https://travis-ci.org/mazznoer/colorgrad)
[![codecov](https://codecov.io/gh/mazznoer/colorgrad/branch/master/graph/badge.svg)](https://codecov.io/gh/mazznoer/colorgrad)

Fun & easy way to create _color gradient_ / _color scales_ in __Go__ (__Golang__).

![color-scale](img/color-scale-1.png)

### Index

* [Usages](#usages)
  - [Basic](#basic)
  - [Custom Colors](#custom-colors)
* [Preset Gradients](#preset-gradients)
* [Color Scheme](#color-scheme)
* [Random Colors](#random-colors)
* [Playground](#go-playground)
* [Dependencies](#dependencies)

### Usages

#### Basic

```go
import "github.com/mazznoer/colorgrad"
```

```go
grad, err := colorgrad.NewGradient().Build()

if err != nil {
    panic(err)
}

// Get single color at certain position.
grad.At(0) // colorful.Color
grad.At(0.5).Hex() // hex color string
grad.At(1)

// Get n colors evenly spaced across gradient.
grad.Colors(27) // []colorful.Color
colorgrad.IntoColors(grad.Colors(10)) // []color.Color
```
![img](img/black-to-white.png)

#### Custom Colors

`Colors()` method accept anything that implement [color.Color](https://golang.org/pkg/image/color/#Color) interface.

```go
import "image/color"
import "github.com/lucasb-eyer/go-colorful"

grad, err := colorgrad.NewGradient().
    Colors(
        color.RGBA{255, 0, 255, 255},
        color.Gray{100},
        color.White,
        colorful.Hsv(210, 1, 0.8),
    ).
    Build()
```
![img](img/basic-2.png)

#### Using Hex Colors

```go
grad, err := colorgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Build()
```
![img](img/basic-hex.png)

#### Custom Domain

By default domain are in the range `0..1`, but you can change it using using the `Domain()` method.

```go
grad, err := colorgrad.NewGradient().
    HexColors("#DC143C", "#FFD700", "#4682b4").
    Domain(15, 47.5, 80).
    Build()

grad.At(15).Hex() // #DC143C
grad.At(75)
grad.At(80).Hex() // #4682b4
```
![img](img/color-scale-2.png)

#### Blending Mode

```go
grad, err := colorgrad.NewGradient().
    HexColors("#ff0", "#008ae5").
    Mode(colorgrad.LRGB).
    Build()
```
![blend-modes](img/blend-modes.png)

#### Beware of Invalid RGB Color
Read it [here](https://github.com/lucasb-eyer/go-colorful#blending-colors).

```go
grad, err := colorgrad.NewGradient().
    HexColors("#DC143C", "#FFD700", "#4682b4").
    Mode(colorgrad.HCL).
    Build()

grad.At(t) // might get invalid RGB color
grad.At(t).Clamped() // return closest valid RGB color
```

Without `Clamped()`
![invalig rgb](img/not-clamped.png)

With `Clamped()`
![valid rgb](img/clamped.png)

### Preset Gradients

```go
grad := colorgrad.Rainbow()
grad.At(t) // t in the range 0..1
grad.Colors(15)
```

`grad := colorgrad.Turbo()`
![img](img/gradient-turbo.png)

`grad := colorgrad.Warm()`
![img](img/gradient-warm.png)

`grad := colorgrad.Cool()`
![img](img/gradient-cool.png)

`grad := colorgrad.Rainbow()`
![img](img/gradient-rainbow.png)

`grad := colorgrad.CubehelixDefault()`
![img](img/gradient-cubehelix-default.png)

`grad := colorgrad.Sinebow()`
![img](img/gradient-sinebow.png)

`grad := colorgrad.Spectral()`
![img](img/gradient-spectral.png)

`grad := colorgrad.Viridis()`
![img](img/gradient-viridis.png)

`grad := colorgrad.Magma()`
![img](img/gradient-magma.png)

`grad := colorgrad.Plasma()`
![img](img/gradient-plasma.png)

`grad := colorgrad.Inferno()`
![img](img/gradient-inferno.png)

`grad := colorgrad.Cividis()`
![img](img/gradient-cividis.png)

`grad := colorgrad.Blues()`
![img](img/gradient-blues.png)

`grad := colorgrad.Greens()`
![img](img/gradient-greens.png)

`grad := colorgrad.Greys()`
![img](img/gradient-greys.png)

`grad := colorgrad.Oranges()`
![img](img/gradient-oranges.png)

`grad := colorgrad.Purples()`
![img](img/gradient-purples.png)

`grad := colorgrad.Reds()`
![img](img/gradient-reds.png)

### Color Scheme

It just a slice of colors. But you can create gradient from color scheme.

```go
grad, err := colorgrad.NewGradient().
    Colors(colorgrad.Scheme.Set1...).
    Build()
```

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

### Random Colors

![random-color](img/random-cool.png)

### Go Playground

* [Basic](https://play.golang.org/p/rE8OI50PsQA)
* [Random colors](https://play.golang.org/p/d67x9di4sAF)

### Dependencies

* [colorful](https://github.com/lucasb-eyer/go-colorful)

### Inspirations

* [chroma.js](https://github.com/gka/chroma.js)
* [d3-scale-chromatic](https://github.com/d3/d3-scale-chromatic/)
