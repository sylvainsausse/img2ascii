package utils

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func Compress(img image.Image, w int, h int) image.Image {
	var a, b int
	var neo *image.RGBA

	neo = image.NewRGBA(image.Rect(0, 0, w, h))

	a = img.Bounds().Max.X
	b = img.Bounds().Max.Y
	ratiow := a / w
	ratioh := b / h
	fmt.Println(a, b, w, h)
	fmt.Println(ratiow, ratioh)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			cpt := [3]uint32{0, 0, 0}
			for k := 0; k < ratiow; k++ {
				for l := 0; l < ratioh; l++ {
					r, g, blue, _ := img.At(ratiow*i+k, ratioh*j+l).RGBA()
					cpt[0] += r >> 8
					cpt[1] += g >> 8
					cpt[2] += blue >> 8
				}
			}
			cpt[0] /= uint32(ratiow * ratioh)
			cpt[1] /= uint32(ratiow * ratioh)
			cpt[2] /= uint32(ratiow * ratioh)
			neo.Set(i, j, color.RGBA{uint8(cpt[0]), uint8(cpt[1]), uint8(cpt[2]), 255})
		}
	}
	return neo
}

func BW(img image.Image) image.Image {
	var a, b int

	a = img.Bounds().Max.X
	b = img.Bounds().Max.Y

	neo := image.NewRGBA(image.Rect(0, 0, a, b))

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			r, g, blue, _ := img.At(i, j).RGBA()
			grey := uint8(((r >> 8) + (g >> 8) + (blue >> 8)) / 3)
			neo.Set(i, j, color.RGBA{grey, grey, grey, 255})
		}
	}

	return neo
}

func Ascii(img image.Image) string {
	var a, b int
	result := ""
	list := " .:-=+*#%@"
	a = img.Bounds().Max.X
	b = img.Bounds().Max.Y
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			grey, _, _, _ := img.At(j, i).RGBA()
			result += string(list[(int(grey>>8)*len(list))/256])
		}
		result += "\n"
	}
	return result
}
