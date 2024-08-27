package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, _ := strconv.Atoi(r.URL.Query().Get("cycles"))
		if c <= 0 {
			c = 1
		}
		lissajous(w, c)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(w http.ResponseWriter, cycles int) {
	const (
		res     = 0.001 // resolução angular
		size    = 100   // tamanho da imagem [-size..+size]
		nframes = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*float64(cycles)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) // NOTA: ignorando erros de codificação
}
