package convertImage

import (
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
)

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
	file.Write([]byte("P2\n#gray image\n" + strconv.Itoa(bounds.Dx()) +
		" " + strconv.Itoa(bounds.Dy()) + "\n255\n"))
}

func ppmEncode(file *os.File, img image.Image) {
	bounds := img.Bounds()
	dstline := make([]string, bounds.Dx())

	ppmWriteHeader(file, bounds)
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			dstline[x] = strconv.Itoa(int(c.R)) + " " +
				strconv.Itoa(int(c.G)) + " " + strconv.Itoa(int(c.B))
		}
		file.Write([]byte(strings.Join(dstline, "\n") + "\n"))
	}
}

func ppmWriteHeader(file *os.File, bounds image.Rectangle) {
	file.Write([]byte("P3\n#color image\n" + strconv.Itoa(bounds.Dx()) +
		" " + strconv.Itoa(bounds.Dy()) + "\n255\n"))
}
