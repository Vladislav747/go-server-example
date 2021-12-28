// Lissajous генерирует анимированный GIF из случайных
// фигур Лиссажу.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 //Первый цвет палитры
	blackIndex = 0 //Следующий цвет палитры
)

func main() {
	var phrase = "Go Lissajous"
	fmt.Println(phrase)
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	//Генерируем случайное значение
	rand.Seed(time.Now().UTC().UnixNano())

	//Приводим
	freq := rand.Float64() * 3.0 // Относительная частота колебаний у
	//Создание экземляра gif
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		//Создание
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
