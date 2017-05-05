package main

import (
	"image"
	"image/color"
	"image/draw"
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
var (
	blue color.Color = color.RGBA{0, 0, 255, 255}
)

func init() {

}

func main() {
	http.HandleFunc("/pic1/", pic1)
	http.HandleFunc("/pic2/", pic2)
	http.HandleFunc("/pic3/", pic3)
	http.HandleFunc("/pic4/", pic4)
	http.HandleFunc("/pic5/", pic5)
	http.HandleFunc("/pic6/", pic6)
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
func pic5(w http.ResponseWriter, req *http.Request) {
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
func pic6(w http.ResponseWriter, req *http.Request) {
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		img.Set(x, dy/2, color.RGBA{uint8(x % 255), uint8(x % 255), 0, 255})
	}
	for y := 0; y < dy/5; y++ {
		img.Set(5, y, color.RGBA{uint8(y % 255), uint8(y % 255), 0, 255})
	}
	for y := 5; y < dy/5; y++ {
		img.Set(56, y, color.RGBA{uint8(y % 255), uint8(y % 255), 0, 255})
	}
	for x := 11; x < 56; x++ {
		img.Set(x, 5, color.RGBA{uint8(x % 255), uint8(x % 255), 0, 255})
	}
	for x := 11; x < 56; x++ {
		img.Set(x, 35, color.RGBA{uint8(x % 255), uint8(x % 255), 0, 255})
	}
	for x := 11; x < 56; x++ {
		img.Set(x, dy/5, color.RGBA{uint8(x % 255), uint8(x % 255), 0, 255})
	}
	zero(img, 61, 5, 20, 50)
	one(img, 85, 5, 1, 50)
	eight(img, 90, 5, 20, 50)
	two(img, 120, 5, 20, 50)
	three(img, 150, 5, 20, 50)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
func zero(img *image.NRGBA, offx, offy, dx, dy int) {
	for y := offy; y < dy+offy; y++ {
		img.Set(offx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for y := offy; y < dy+offy; y++ {
		img.Set(offx+dx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
}
func one(img *image.NRGBA, offx, offy, dx, dy int) {
	for y := offy; y < dy+offy; y++ {
		img.Set(offx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
}
func two(img *image.NRGBA, offx, offy, dx, dy int) {
	for y := offy; y < dy+offy; y++ {
		img.Set(offx+dx, y/2, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for y := offy + dy/2; y < dy+offy; y++ {
		img.Set(offx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy/2, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
}
func three(img *image.NRGBA, offx, offy, dx, dy int) {
	for y := offy; y < dy+offy; y++ {
		img.Set(offx+dx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy/2, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
}
func eight(img *image.NRGBA, offx, offy, dx, dy int) {
	for y := offy; y < dy+offy; y++ {
		img.Set(offx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for y := offy; y < dy+offy; y++ {
		img.Set(offx+dx, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy/2, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
	for x := offx; x < dx+offx; x++ {
		img.Set(x, offy+dy, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}
}
