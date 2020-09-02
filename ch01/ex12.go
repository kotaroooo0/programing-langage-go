package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"log"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	settings := Settings{
		Cycles:  5,
		Res:     0.001,
		Size:    100,
		Nframes: 64,
		Delay:   8,
	}

	// TODO: 複数クエリがあった場合どうなる(ex: ?cycles=10,20&)
	cyclesStr := r.URL.Query().Get("cycles")
	if cyclesStr != "" {
		cycles, err := strconv.ParseFloat(cyclesStr, 32)
		if err != nil {
			log.Fatal(err)
		}
		settings.Cycles = cycles
	}

	// TODO: 他のクエリも同様に

	lissajous(w, settings)
}

var palette = []color.Color{color.Black, color.RGBA{1, 200, 1, 1}, color.RGBA{200, 1, 1, 1}, color.RGBA{1, 1, 200, 1}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	blueIndex  = 3
)

type Settings struct {
	Cycles  float64
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func lissajous(out io.Writer, settings Settings) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: settings.Nframes}
	phase := 0.0
	for i := 0; i < settings.Nframes; i++ {
		rect := image.Rect(0, 0, 2*settings.Size+1, 2*settings.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < settings.Cycles*2*math.Pi; t += settings.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			r := rand.Float32()
			if r < 0.33 {
				img.SetColorIndex(settings.Size+int(x*float64(settings.Size)+0.5), settings.Size+int(y*float64(settings.Size)+0.5), greenIndex)
			} else if r < 0.66 {
				img.SetColorIndex(settings.Size+int(x*float64(settings.Size)+0.5), settings.Size+int(y*float64(settings.Size)+0.5), redIndex)
			} else {
				img.SetColorIndex(settings.Size+int(x*float64(settings.Size)+0.5), settings.Size+int(y*float64(settings.Size)+0.5), blueIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, settings.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
