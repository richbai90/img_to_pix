package utils

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

func DecodeImg(filename string) image.Image {
	var img image.Image
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		log.Fatal("Failed to open image ", filename, ".\nERROR: ", err.Error())
	}

	// check if png file
	if img, err = png.Decode(f); err == nil {
		return img
	}
	if img, err = jpeg.Decode(f); err == nil {
		return img
	}
	if img, err = bmp.Decode(f); err == nil {
		return img
	}
	if img, err = tiff.Decode(f); err == nil {
		return img
	}
	if img, err = gif.Decode(f); err == nil {
		return img
	}

	log.Fatal("Could not determine image type for ", filename)
	return nil
}

func GetBytes(img image.Image) []byte {
	// buffer := bytes.Buffer{}
	rgba, ok := (img).(*image.RGBA)
	if ok {
		return rgba.Pix

	} else {
		rgba := img.(*image.NRGBA)
		return rgba.Pix

	}
}
