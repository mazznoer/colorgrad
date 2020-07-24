# colorgrad

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/colorgrad?tab=doc)

Fun & easy way to create _color gradient_ / _color scales_ in __Go__ (__Golang__).

## Usages

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
gradient.At(0.5)
gradient.At(1)

// Get n colors evenly spaced across gradient.
gradient.Colors(27) // []colorful.Color
```

![img](doc/black-to-white.png)

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

![img](doc/basic-2.png)

#### Using Hex Colors

```go
gradient, err := colgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Build()
```

![img](doc/basic-hex.png)

#### Custom Domain

```go
gradient, err := colgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Domain(0, 75, 100).
    Build()
```

![img](doc/basic-hex.png)

#### Blending Mode

```go
gradient, err := colgrad.NewGradient().
    HexColors("#B22222", "#FFD700", "#2E8B57").
    Mode(colgrad.HCL).
    Build()
```

![img](doc/basic-2.png)
![img](doc/basic-hex.png)
![img](doc/basic-2.png)
![img](doc/basic-hex.png)

#### Get Hex Color

```go
for _, c := range gradient.Colors(15) {
    fmt.Println(c.Hex())
}
```

#### Random Color from Gradient

```go
import "math/rand"

for i := 0; i < 100; i++ {
    fmt.Println(gradient.At(rand.Float64()))
}
```

## Online Playground

[Try it online](https://play.golang.org/p/7zaL_OQ4Gbf)

## Dependencies

* [colorful](https://github.com/lucasb-eyer/go-colorful)

## Author

* [Mazznoer](https://github.com/mazznoer)
