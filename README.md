# colorgrad

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/colorgrad?tab=doc)

Fun & easy way to create _color gradient_ / _color scales_ in __Go__ (__Golang__).

### Usages

#### Basic

```go
import "github.com/mazznoer/colorgrad"
```

```go
gradient, err := colorgrad.NewGradient().Build()

if err != nil {
    panic(err)
}

// Get single color at certain position.
gradient.At(0) // colorful.Color
gradient.At(0.5).Hex() // hex color string
gradient.At(1)

// Get n colors evenly spaced across gradient.
gradient.Colors(27) // []colorful.Color
colorgrad.IntoColors(gradient.Colors(10)) // []color.Color
```

![img](img/black-to-white.png)

#### Custom Colors

`Colors()` method accept anything that implement [color.Color](https://golang.org/pkg/image/color/#Color) interface.

```go
import "image/color"
import "github.com/lucasb-eyer/go-colorful"

gradient, err := colorgrad.NewGradient().
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
gradient, err := colorgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Build()
```

![img](img/basic-hex.png)

#### Custom Domain

```go
gradient, err := colorgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Domain(0, 75, 100).
    Build()
```

![img](img/basic-hex.png)

#### Blending Mode

```go
gradient, err := colorgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Mode(colgrad.HCL).
    Build()
```

![img](img/basic-2.png)
![img](img/basic-hex.png)

#### Random Color from Gradient

```go
import "math/rand"

for i := 0; i < 100; i++ {
    fmt.Println(gradient.At(rand.Float64()))
}
```

### Online Playground

[Try it online](https://play.golang.org/p/7zaL_OQ4Gbf)

### Dependencies

* [colorful](https://github.com/lucasb-eyer/go-colorful)

### Inspirations

* [chroma.js](https://github.com/gka/chroma.js)
* [d3-scale-chromatic](https://github.com/d3/d3-scale-chromatic/)
