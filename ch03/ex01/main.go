package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				fmt.Fprintf(os.Stderr, "error:nan\n")
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				fmt.Fprintf(os.Stderr, "error:nan\n")
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				fmt.Fprintf(os.Stderr, "error:nan\n")
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				fmt.Fprintf(os.Stderr, "error:nan\n")
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ok
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	result := math.Sin(r) / r
	return result, !math.IsNaN(result) && !math.IsInf(result, 0) && !math.IsInf(result, -1)
}
