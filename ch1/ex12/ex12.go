package main

import (
	"image/color"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"net/http"
	"log"
	"time"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	cycles := 5 // number of complete x oscillator revolutions
	if cyclesStr := r.FormValue("cycles"); cyclesStr != "" {
		if value, err := strconv.Atoi(cyclesStr); err == nil {
			cycles = value
		}
	}
	lissajous(cycles, w)
}

func lissajous(cycles int, out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{
			R: uint8(55 + 200*scale),
			G: uint8(55 + 200*scale),
			B: uint8(55 + 200*scale),
			A: 255,
		}
		palette = append(palette, c)
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		index := 0
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i%(len(palette)-1))+1))
			index++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
