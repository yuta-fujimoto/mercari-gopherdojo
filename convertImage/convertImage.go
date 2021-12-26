/*
Package convert enables to convert between JPEG and PNG, GIF. Also, it can make black/white image(PGM) from
color image(JPEG, PNG, GIF).
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
Convert all image files in directory or filepath itself specified as string arg. inForm and outForm are I/O image format. If some sort of error occurs(failed to read directory, invalid format(txt, pdf, etc)), ConvertImage returns proper error and do nothing. Unnecessary format .jpg .png .pgm .gif are ignored if arg is specified as directory.
*/
func ConvertImage(arg string, inForm string, outForm string) error {
	params, err := initParams(arg, inForm, outForm)
	if err != nil {
		return err
	}
	input := make([]*os.File, params.size)
	output := make([]*os.File, params.size)

	for i := 0; i < params.size; i++ {
		input[i], err = os.Open(params.Infile[i])
		if err != nil {
			closeAllFiles(input, output, i, i)
			return getError(err)
		}
		output[i], err = os.Create(params.Outfile[i])
		if err != nil {
			closeAllFiles(input, output, i+1, i)
			return getError(err)
		}
	}
	defer func() {
		closeAllFiles(input, output, params.size, params.size)
	}()
	for i := 0; i < params.size; i++ {
		img, _, err := image.Decode(input[i])
		if err != nil {
			return err
		}
		switch params.Outform {
		case PNG:
			png.Encode(output[i], img)
		case JPEG:
			jpeg.Encode(output[i], img, &jpeg.Options{Quality: 100})
		case GIF:
			gif.Encode(output[i], img, &gif.Options{NumColors: 256})
		case PGM:
			pgmEncode(output[i], img)
		}
	}
	return nil
}

func closeAllFiles(input []*os.File, output []*os.File, inCnt int, outCnt int) {
	var err error

	for i := 0; i < inCnt; i++ {
		err = input[i].Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}
	for i := 0; i < outCnt; i++ {
		err = output[i].Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}
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
