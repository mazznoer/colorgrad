# colorgrad 🎨

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/colorgrad?tab=doc)

Fun & easy way to create _color gradient_ / _color scales_ in __Go__ (__Golang__).

![color-scale](img/color-scale-1.png)

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

```go
grad, err := colorgrad.NewGradient().
    HexColors("#DC143C", "#FFD700", "#4682b4").
    Domain(15, 47.5, 80).
    Build()
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

`grad := colorgrad.Oranges()`
![img](img/gradient-oranges.png)

`grad := colorgrad.Purples()`
![img](img/gradient-purples.png)

`grad := colorgrad.Reds()`
![img](img/gradient-reds.png)

`grad := colorgrad.Greys()`
![img](img/gradient-greys.png)

### Random Colors

![random-color](img/random-cool.png)
[Try it online](https://play.golang.org/p/d67x9di4sAF)

### Online Playground

[Try it online](https://play.golang.org/p/rE8OI50PsQA)

### Dependencies

* [colorful](https://github.com/lucasb-eyer/go-colorful)

### Inspirations

* [chroma.js](https://github.com/gka/chroma.js)
* [d3-scale-chromatic](https://github.com/d3/d3-scale-chromatic/)
