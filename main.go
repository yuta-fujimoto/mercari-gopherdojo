package main

import (
	"convert/convertImage"
	"flag"
	"fmt"
	"os"
)

func main() {
	inForm := flag.String("i", "jpg", "input file format(jpg, png, gif)")
	outForm := flag.String("o", "png", "output file format(jpg, png, gif, pgm, ppm)")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		return
	}
	err := convertImage.ConvertImage(args[0], *inForm, *outForm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
