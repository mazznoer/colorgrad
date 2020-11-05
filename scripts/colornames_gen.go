// Code generator for colornames.go
// Run `gofmt` on output code

// +build ignore

package main

import (
	"fmt"

	cn "golang.org/x/image/colornames"
)

var header = `
package colorgrad

import "github.com/lucasb-eyer/go-colorful"

var colornames = map[string]colorful.Color{
`

func main() {
	fmt.Print(header)
	for _, name := range cn.Names {
		c := cn.Map[name]
		r := float64(c.R) / 255
		g := float64(c.G) / 255
		b := float64(c.B) / 255
		fmt.Printf("\t\"%s\": {R:%v, G:%v, B:%v},\n", name, r, g, b)
	}
	fmt.Println("}")
}
