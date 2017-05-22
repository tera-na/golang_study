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
	"log"
	"math/big"
	"math/cmplx"
	"net/http"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := new(big.Rat).SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := new(big.Rat).SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			fx, _ := x.Float64()
			fy, _ := y.Float64()
			z := complex(fx, fy)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img) // NOTE: ignoring errors
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// func mandelbrot(real_val, imag_val *big.Rat) color.Color {
// 	const iterations = 20
// 	const contrast = 15

// 	var v_real = new(big.Rat).SetFloat64(0)
// 	var v_imag = new(big.Rat).SetFloat64(0)
// 	c := new(big.Rat).SetInt64(2)
// 	for n := uint8(0); n < iterations; n++ {
// 		v_real, v_imag = seki(v_real, v_imag, v_real, v_imag)
// 		v_real = v_real.Add(v_real, real_val)
// 		v_imag = v_imag.Add(v_imag, imag_val)
// 		tmp := v_real.Abs(v_real)
// 		if tmp.Cmp(c) == -1 {
// 			return color.Gray{255 - contrast*n}
// 		}
// 	}
// 	return color.Black
// }

func seki(real1, imag1, real2, imag2 *big.Rat) (*big.Rat, *big.Rat) {

	real_z := real1.Mul(real1, real2)
	imag_z := imag1.Mul(imag1, imag2)

	real_z = real_z.Sub(real_z, imag_z)
	imag_z = imag_z.Add(real1.Mul(real1, imag2), real2.Mul(real2, imag1))

	return real_z, imag_z
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{255 - contrast*i, 127 - contrast*i, 127 - contrast*i, 255}
		}
	}
	return color.Black
}
