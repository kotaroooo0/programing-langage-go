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
	file, err := os.Create("sample.svg")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-wdith: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	var cs = [][]float64{}
	var max, min = math.Inf(-1), math.Inf(0)

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

			y := (ay + by + cy + dy) * .25
			if y > max {
				max = y
			}
			if y < min {
				min = y
			}
			cs = append(cs, []float64{ax, ay, bx, by, cx, cy, dx, dy})
		}
	}
	for _, c := range cs {
		y := (c[1] + c[3] + c[5] + c[7]) * .25
		norm := (y - min) / (max - min)

		fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: rgba(%v,%v,%v,0.5); stroke-width: 0.3' />\n",
			c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7], 255*(1-norm), 0, 255*norm)
	}

	fmt.Fprintf(file, "</svg>\n")
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
