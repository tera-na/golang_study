// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	var z_max float64 = 0.0
	var z_min float64 = 0.0
	// Z軸の最大最小値を求める
	for i := 0; i < xyrange; i++ {
		for j := 0; j < xyrange; j++ {
			z := f(float64(i), float64(j)) * zscale
			if z > z_max {
				z_max = z
			}
			if z < z_min {
				z_min = z
			}
		}
	}

	// 出力
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ac := corner(i+1, j)
			bx, by, bc := corner(i, j)
			cx, cy, cc := corner(i, j+1)
			dx, dy, dc := corner(i+1, j+1)

			red, blue := 0, 0
			color := (ac*zscale + bc*zscale + cc*zscale + dc*zscale) / 4
			if color > 0 {
				red = int((color / z_max) * 255.0)
				if red > 255 {
					red = 255
				}
			}
			if color < 0 {
				blue = int((math.Abs(color) / math.Abs(z_min)) * 255.0)
				if blue > 255 {
					blue = 255
				}
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%02x00%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, red, blue)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) (value float64) {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
