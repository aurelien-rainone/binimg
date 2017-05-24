package imgscan

import (
	"fmt"
	"image"
	"image/color"

	"github.com/aurelien-rainone/imgtools/binimg"
)

// A Scanner behaves as an image.Image, with added color scanning capabilities.
//
// For example, it can report wether a particular region of the image it embeds
// is uniform (i.e made of an unique color), what is this uniform color, or
// compute the average color.
type Scanner interface {
	image.Image

	// IsUniformColor indicates if the region r is only made of pixels of color c.
	//
	// The scan stops at the first pixel encountered that is different from c.
	IsUniformColor(r image.Rectangle, c color.Color) bool

	// IsUniform indicates if the region r is uniform. If that is the case, the
	// uniform color is returned, otherwise the returned color is nil.
	//
	// The scan stops at the first pixel encountered that is different from the
	// previous one.
	IsUniform(r image.Rectangle) (bool, color.Color)

	// AverageColor indicates wether the region is uniform and the average color
	// of the region r. If all the pixels have the same color (i.e the region is
	// uniform) then the average color is that color.
	//
	// A full scan of the region is performed in order to determine the average
	// color.
	AverageColor(r image.Rectangle) (bool, color.Color)
}

// NewScanner returns a new Scanner of the given image.Image.
//
// The actual scanner implementation depends on the image bit depth and the
// availability of an implementation.
func NewScanner(img image.Image) (Scanner, error) {
	var (
		s   Scanner
		err error
	)
	switch img.(type) {
	case *binimg.Binary:
		s = NewBinaryScanner(img.(*binimg.Binary))
	case *image.Gray:
		s = NewGrayScanner(img.(*image.Gray))
	default:
		err = fmt.Errorf("unsupported image type")

	}
	return s, err
}
