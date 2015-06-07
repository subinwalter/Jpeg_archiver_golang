/*
 * contains utilities for jpeg image compression
 */
package util

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

//reads jpeg image from a file and store it in buffer
func JpegReader(filename string) *bytes.Buffer {
	infile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = infile.Close(); err != nil {
			panic(err)
		}
	}()
	var buff bytes.Buffer
	_, err = buff.ReadFrom(infile)
	if err != nil {
		panic(err)
	}
	return &buff
}

//writes jpeg image to a file
func JpegWriter(filename string, buff *bytes.Buffer) {
	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = outfile.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = buff.WriteTo(outfile)
	if err != nil {
		panic(err)
	}
}

//decodes jpeg image
func JpegDecode(buff *bytes.Buffer) image.Image {
	sbyte := buff.Bytes()
	sreader := bytes.NewReader(sbyte)

	img, err := jpeg.Decode(sreader)
	if err != nil {
		panic(err)
	}
	return img
}

//encode jpeg image accoring to quality,quality range 0-100
func JpegEncode(img image.Image, quality int) *bytes.Buffer {
	var buf bytes.Buffer
	var opt jpeg.Options
	opt.Quality = quality
	err := jpeg.Encode(&buf, img, &opt)

	if err != nil {
		panic(err)
	}
	return &buf
}

//convert decoded jpeg image to greyscale for comparison
func Grayscale(img image.Image) image.Image {
	bound := img.Bounds()

	w, h := bound.Max.X, bound.Max.Y
	gray := image.NewGray(bound)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}
	return gray
}

//convert encoded jpeg image to greyscale for comparison
func JpegDecodeGray(buff *bytes.Buffer) image.Image {
	sbyte := buff.Bytes()
	sreader := bytes.NewReader(sbyte)

	img, err := jpeg.Decode(sreader)
	if err != nil {
		panic(err)
	}
	bound := img.Bounds()

	w, h := bound.Max.X, bound.Max.Y
	gray := image.NewGray(bound)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}
	return gray

}

//image comparison technique mean pixel error
func MeanPixelError(grayoriginal, graycompress image.Image) float64 {
	var ogr, cgr color.Gray
	var mpe float64
	mpe = 0
	bound := grayoriginal.Bounds()
	w, h := bound.Max.X, bound.Max.Y
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			ogr = (grayoriginal.At(x, y)).(color.Gray)
			cgr = (graycompress.At(x, y)).(color.Gray)
			temp := float64(ogr.Y)
			temp1 := float64(cgr.Y)
			mpe = mpe + math.Abs(temp-temp1)
		}
	}
	mpe = mpe / (float64(w) * float64(h))
	return mpe
}
