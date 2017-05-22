// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		py2 := (py + 1) % height
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py2)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			px2 := (px + 1) % width
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px2)/width*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			co1 := mandelbrot(z1)
			co2 := mandelbrot(z2)
			co3 := mandelbrot(z3)
			co4 := mandelbrot(z4)

			r1, g1, b1, _ := co1.RGBA()
			r2, g2, b2, _ := co2.RGBA()
			r3, g3, b3, _ := co3.RGBA()
			r4, g4, b4, _ := co4.RGBA()
			r := (r1 + r2 + r3 + r4) / 4
			g := (g1 + g2 + g3 + g4) / 4
			b := (b1 + b2 + b3 + b4) / 4
			a := 255

			// Image point (px, py) represents complex value z.
			img.Set(px, py, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// r := 255 - n*15
			// g := n * 15
			// b := 127 - n*15

			// return color.RGBA{r, g, b, 255}
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Black
}
