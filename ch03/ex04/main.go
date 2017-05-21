// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var width int64
var height int64   // canvas size in pixels
var xyscale int64  // pixels per x or y unit
var zscale float64 // pixels per z unit

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {

	r.ParseForm()

	width = 600
	val := r.Form.Get("width")
	fmt.Printf("val=%v\n", val)
	getint, err := strconv.ParseInt(val, 10, 64)
	fmt.Printf("err=%v\n", err)
	if err == nil {
		width = getint
	}
	fmt.Printf("width=%d\n", width)

	height = 320
	val = r.Form.Get("height")
	getint, err = strconv.ParseInt(val, 10, 64)
	if err == nil {
		height = getint
	}
	fmt.Printf("height=%d\n", height)

	color := "red"
	val = r.Form.Get("color")
	if len(val) > 0 {
		color = val
	}
	fmt.Printf("color=%v\n", color)

	xyscale = width / 2 / xyrange
	zscale = float64(height) * 0.4

	// 出力
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	xxx := (x - y) * cos30 * float64(xyscale)
	yyy := (x + y) * sin30 * float64(xyscale)
	sx := float64(width/2) + xxx
	sy := float64(height/2) + yyy - z*zscale
	return sx, sy, z
}

func f(x, y float64) (value float64) {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
