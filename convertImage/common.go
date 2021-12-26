package convertImage

// valid image format(PGM is for output only)
const (
	JPEG = ".jpg"
	PNG  = ".png"
	PGM  = ".pgm"
	GIF  = ".gif"
)

// ConvertImage at first opens all files and specifies formats to store them into Params
type Params struct {
	Infile  []string
	Outfile []string
	Inform  string
	Outform string
	size    int
}
