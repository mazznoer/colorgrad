package main

import (
    "image"
    "image/png"
    "os"
    
    "github.com/mazznoer/colorgrad"
)

func main() {
    grad := colorgrad.NewGradient().Build()
    w := 800
    h := 90
    fw := float64(w)

    img := image.NewRGBA(image.Rect(0, 0, w, h))
        
    for x := 0; x < w; x++ {
        col := grad.At(float64(x) / fw)
        
        for y := 0; y < h; y++ {
            img.Set(x, y, col)
        }
    }

    file, err := os.Create("gradient.png")

    if err != nil {
        panic(err.Error())
    }
    defer file.Close()
    
    png.Encode(file, img)
}
