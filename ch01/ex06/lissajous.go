package main

import (
	"image/color"
	"os"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"net/http"
	"log"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		return

	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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
		for t := 0.0; t < cycles*2*math.Pi; t += res {
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