package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"os"
)

const dx, dy = 300, 500

var font = [][]byte{

	{ // 0

		0, 1, 1, 1, 0,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		0, 1, 1, 1, 0,
	},

	{ // 1

		0, 0, 1, 0, 0,

		0, 1, 1, 0, 0,

		1, 0, 1, 0, 0,

		0, 0, 1, 0, 0,

		0, 0, 1, 0, 0,

		0, 0, 1, 0, 0,

		0, 0, 1, 0, 0,

		1, 1, 1, 1, 1,
	},

	{ // 2

		0, 1, 1, 1, 0,

		1, 0, 0, 0, 1,

		0, 0, 0, 0, 1,

		0, 0, 0, 1, 1,

		0, 1, 1, 0, 0,

		1, 0, 0, 0, 0,

		1, 0, 0, 0, 0,

		1, 1, 1, 1, 1,
	},

	{ // 3

		1, 1, 1, 1, 0,

		0, 0, 0, 0, 1,

		0, 0, 0, 1, 0,

		0, 1, 1, 1, 0,

		0, 0, 0, 1, 0,

		0, 0, 0, 0, 1,

		0, 0, 0, 0, 1,

		1, 1, 1, 1, 0,
	},

	{ // 4

		1, 0, 0, 1, 0,

		1, 0, 0, 1, 0,

		1, 0, 0, 1, 0,

		1, 0, 0, 1, 0,

		1, 1, 1, 1, 1,

		0, 0, 0, 1, 0,

		0, 0, 0, 1, 0,

		0, 0, 0, 1, 0,
	},

	{ // 5

		1, 1, 1, 1, 1,

		1, 0, 0, 0, 0,

		1, 0, 0, 0, 0,

		1, 1, 1, 1, 0,

		0, 0, 0, 0, 1,

		0, 0, 0, 0, 1,

		0, 0, 0, 0, 1,

		1, 1, 1, 1, 0,
	},

	{ // 6

		0, 0, 1, 1, 1,

		0, 1, 0, 0, 0,

		1, 0, 0, 0, 0,

		1, 1, 1, 1, 0,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		0, 1, 1, 1, 0,
	},

	{ // 7

		1, 1, 1, 1, 1,

		0, 0, 0, 0, 1,

		0, 0, 0, 0, 1,

		0, 0, 0, 1, 0,

		0, 0, 1, 0, 0,

		0, 1, 0, 0, 0,

		0, 1, 0, 0, 0,

		0, 1, 0, 0, 0,
	},

	{ // 8

		0, 1, 1, 1, 0,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		0, 1, 1, 1, 0,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		0, 1, 1, 1, 0,
	},

	{ // 9

		0, 1, 1, 1, 0,

		1, 0, 0, 0, 1,

		1, 0, 0, 0, 1,

		1, 1, 0, 0, 1,

		0, 1, 1, 1, 1,

		0, 0, 0, 0, 1,

		0, 0, 0, 0, 1,

		1, 1, 1, 1, 0,
	},
}

func init() {

}

func main() {
	http.HandleFunc("/pic1/", pic1)
	http.HandleFunc("/pic2/", pic2)
	http.HandleFunc("/pic3/", pic3)
	http.HandleFunc("/pic4/", pic4)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func drawpng() {
	imgfile, _ := os.Create("1.png")
	defer imgfile.Close()
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), uint8(y % 256), uint8(x % 256)})
		}
	}
	if err := png.Encode(imgfile, img); err != nil {
		log.Fatal(err)
	}
}
func pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j % 255)
		}
	}
	return pic
}
func pic1(w http.ResponseWriter, req *http.Request) {
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), uint8(y % 256), uint8(x % 256)})
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
func pic2(w http.ResponseWriter, req *http.Request) {
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {

			if x%10 == 0 {
				img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
			}
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
func pic3(w http.ResponseWriter, req *http.Request) {
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			if x%10 == 0 {
				img.Set(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
			}
		}
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
func pic4(w http.ResponseWriter, req *http.Request) {
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if x%2 == 0 {
				img.Set(x, y, color.RGBA{uint8(x * y % 255), uint8(x * y % 255), 0, 255})
			}
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
