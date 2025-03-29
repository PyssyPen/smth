package image

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/golang/freetype"
)

type Service struct {
}

func NewImageService() *Service {
	return &Service{}
}

func (i *Service) DrawText(inputFileName, topText, bottomText string) (string, error) {
	imgFile, err := os.Open(inputFileName)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	fontBytes, err := os.ReadFile("/home/pyssy/VSC/GO/meme_bot/fonts/Pragmpla.ttf")
	if err != nil {
		return "", err
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return "", err
	}

	originalImage, _, err := image.Decode(imgFile)
	if err != nil {
		return "", err
	}

	// Получаем размеры изображения
	// Bounds := originalImage.Bounds()
	// Width := Bounds.Dx()
	// Height := Bounds.Dy()
	// Выводим размеры изображения
	//fmt.Printf("Ширина: %d, Высота: %d\n", Width, Height)

	// так мы можем настраивать положение текста на картинке (начало текста в середине по Х)
	x := originalImage.Bounds().Dx() / 2
	// дальше нужно узнать длину текста в пикселях (наверно)

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

	pt := freetype.Pt(x, 40) // Position for top text (X, Y)
	_, err = c.DrawString(topText, pt)
	if err != nil {
		return "", err
	}

	// Measure and add bottom text
	pt = freetype.Pt(x, rgba.Bounds().Dy()-10) // Position for bottom text (X, Y)
	_, err = c.DrawString(bottomText, pt)

	if err != nil {
		return "", err
	}

	outFile, err := os.CreateTemp("", "meme")
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, rgba, nil)
	if err != nil {
		return "", err
	}

	return outFile.Name(), nil
}
