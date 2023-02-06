package main

import (
	"fmt"
	utils "fredoinc/img2ascii/utils"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)

func help() {
	fmt.Println("Convert png or jpeg to ascii")
	fmt.Println("usage : img2ascii [-h] {filename} [-o|-d|-a]")
	fmt.Println("-h display this text")
	fmt.Println("-o {filename}\tset the output file")
	fmt.Println("-d {int} {int}\t set the size of the ascii art")
	fmt.Println("-d {int} {int}\t set the size of the ascii art")
	fmt.Println("-a \t\t output all the stage of conversion in current dir")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(3)
	}

	vboz := false
	arg := os.Args[2:]
	out := "out.txt"
	dim := [2]int{200, 200}

	if utils.Occur(arg, "-a") != -1 {
		vboz = true
	}

	v := utils.Occur(os.Args, "-h")

	if v != -1 {
		help()
		os.Exit(0)
	}

	v = utils.Occur(arg, "-o")

	if v != -1 {
		if v > len(arg)-1 || utils.Startwith(arg[v+1], "-") {
			fmt.Fprintf(os.Stderr, "you need to put a filename after -o option\n")
			os.Exit(2)
		}
		out = arg[v+1]
	}

	v = utils.Occur(arg, "-d")

	if v != -1 {
		if v > len(arg)-2 || utils.Startwith(arg[v+1], "-") || utils.Startwith(arg[v+2], "-") {
			fmt.Fprintf(os.Stderr, "You need to make 2 ints args after -d option\n")
			os.Exit(2)
		}
		var err1, err2 error
		dim[0], err1 = strconv.Atoi(arg[v+1])
		dim[1], err2 = strconv.Atoi(arg[v+2])
		if err1 != nil || err2 != nil {
			os.Exit(1)
		}
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(2)
	}
	file.Close()

	if dim[0] > img.Bounds().Max.X || dim[1] > img.Bounds().Max.Y {
		fmt.Fprintln(os.Stderr, "Size of the conversion exceed the size of the image")
		os.Exit(3)
	}

	// Compressing image
	imaj := utils.Compress(img, dim[0], dim[1])

	if vboz {
		// Saving compressed image
		fo, err := os.Create("img.jpg")
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(4)
		}
		err = jpeg.Encode(fo, imaj, nil)
		fo.Close()
	}

	// Rendering BW image
	imaj = utils.BW(imaj)

	if vboz {
		// Saving compressed image
		fo, err := os.Create("imgBW.jpg")
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(4)
		}
		err = jpeg.Encode(fo, imaj, nil)
		fo.Close()
	}

	st := utils.Ascii(imaj)
	fo, err := os.Create(out)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(4)
	}
	fo.WriteString(st)
}
