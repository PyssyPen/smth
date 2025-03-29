package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"github.com/golang/freetype"
)

func main() {
	imgFile, err := os.Open("/home/pyssy/VSC/GO/meme_bot/jpg/24.jpg")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	fontBytes, err := os.ReadFile("/home/pyssy/VSC/GO/meme_bot/fonts/FredericBlack.ttf")
	if err != nil {
		panic(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	originalImage, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	// create a new RGBA image
	rgba := image.NewRGBA(originalImage.Bounds())
	draw.Draw(rgba, rgba.Bounds(), originalImage, image.Point{}, draw.Src)

	// Create a freetype context for drawing text
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(30)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.NewUniform(color.White))

	topText := "TEST  1"
	bottomText := "HUEST 2"

	pt := freetype.Pt(10, 40) // Position for top text (X, Y)
	_, err = c.DrawString(topText, pt)

	if err != nil {
		log.Fatalf("failed to draw top text: %v", err)
	}

	// Measure and add bottom text
	pt = freetype.Pt(10, rgba.Bounds().Dy()-10) // Position for bottom text (X, Y)
	_, err = c.DrawString(bottomText, pt)

	if err != nil {
		log.Fatalf("failed to draw bottom text: %v", err)
	}

	outputFile, err := os.Create("/home/pyssy/VSC/GO/meme_bot/jpg/ready_meme.jpg")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, rgba, nil)
	if err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}
}
