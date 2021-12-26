/*
Package convert enables to convert between JPEG and PNG. Also, it can make black/white image(PGM) from
color image(JPEG, PNG).
*/
package convertImage

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
)

/*
Convert all image files in directory specified as string arg. inForm and outForm are I/O image format. If some sort of error occurs(failed to read directory, invalid format(txt, pdf, etc)), ConvertImage prints appropriate error messages and calls exit. Format .jpg .png .pgm are ignored even if they are not assigned as argument.
*/
func ConvertImage(arg string, inForm string, outForm string) error {
	params, err := initParams(arg, inForm, outForm)
	defer func() {
		for i := 0; i < len(params.Infile); i++ {
			err = params.Infile[i].Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			err = params.Outfile[i].Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		}
	}()
	if err != nil {
		return err
	}
	for i := 0; i < len(params.Infile); i++ {
		img, _, err := image.Decode(params.Infile[i])
		if err != nil {
			return err
		}
		switch params.Outform {
		case PNG:
			png.Encode(params.Outfile[i], img)
		case JPEG:
			jpeg.Encode(params.Outfile[i], img, &jpeg.Options{Quality: 100})
		case GIF:
			gif.Encode(params.Outfile[i], img, &gif.Options{NumColors: 256})
		case PGM:
			pgmEncode(params.Outfile[i], img)
		}
	}
	return nil
}

func pgmEncode(file *os.File, img image.Image) {
	bounds := img.Bounds()
	dstline := make([]string, bounds.Dx())

	pgmWriteHeader(file, bounds)
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			dstline[x] = strconv.Itoa(int(c.Y))
		}
		file.Write([]byte(strings.Join(dstline, " ") + "\n"))
	}
}

func pgmWriteHeader(file *os.File, bounds image.Rectangle) {
	file.Write([]byte("P2\n#gray file\n" + strconv.Itoa(bounds.Dx()) +
		" " + strconv.Itoa(bounds.Dy()) + "\n255\n"))
}
