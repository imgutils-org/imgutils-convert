// Package convert provides image format conversion utilities.
package convert

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// Format represents an image format.
type Format string

const (
	JPEG Format = "jpeg"
	PNG  Format = "png"
	GIF  Format = "gif"
	BMP  Format = "bmp"
	TIFF Format = "tiff"
)

// Options configures the conversion.
type Options struct {
	Quality int // JPEG quality (1-100), default 85
}

// DefaultOptions returns sensible defaults.
func DefaultOptions() Options {
	return Options{Quality: 85}
}

// Encode writes an image to the writer in the specified format.
func Encode(w io.Writer, img image.Image, format Format, opts Options) error {
	if opts.Quality <= 0 || opts.Quality > 100 {
		opts.Quality = 85
	}

	switch format {
	case JPEG:
		return jpeg.Encode(w, img, &jpeg.Options{Quality: opts.Quality})
	case PNG:
		return png.Encode(w, img)
	case GIF:
		return gif.Encode(w, img, nil)
	case BMP:
		return bmp.Encode(w, img)
	case TIFF:
		return tiff.Encode(w, img, nil)
	default:
		return errors.New("unsupported format")
	}
}

// Decode reads an image from the reader.
func Decode(r io.Reader) (image.Image, Format, error) {
	img, formatStr, err := image.Decode(r)
	if err != nil {
		return nil, "", err
	}
	return img, Format(formatStr), nil
}

// Convert reads an image and converts it to a different format.
func Convert(r io.Reader, w io.Writer, format Format, opts Options) error {
	img, _, err := Decode(r)
	if err != nil {
		return err
	}
	return Encode(w, img, format, opts)
}

// ConvertFile converts an image file to a different format.
func ConvertFile(inputPath, outputPath string, opts Options) error {
	in, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer in.Close()

	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	format := FormatFromExtension(outputPath)
	return Encode(out, img, format, opts)
}

// FormatFromExtension determines the format from a file extension.
func FormatFromExtension(path string) Format {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".jpg", ".jpeg":
		return JPEG
	case ".png":
		return PNG
	case ".gif":
		return GIF
	case ".bmp":
		return BMP
	case ".tiff", ".tif":
		return TIFF
	default:
		return JPEG
	}
}

// ToJPEG converts an image to JPEG format.
func ToJPEG(img image.Image, w io.Writer, quality int) error {
	return Encode(w, img, JPEG, Options{Quality: quality})
}

// ToPNG converts an image to PNG format.
func ToPNG(img image.Image, w io.Writer) error {
	return Encode(w, img, PNG, DefaultOptions())
}

// ToGIF converts an image to GIF format.
func ToGIF(img image.Image, w io.Writer) error {
	return Encode(w, img, GIF, DefaultOptions())
}

// ToBMP converts an image to BMP format.
func ToBMP(img image.Image, w io.Writer) error {
	return Encode(w, img, BMP, DefaultOptions())
}

// ToTIFF converts an image to TIFF format.
func ToTIFF(img image.Image, w io.Writer) error {
	return Encode(w, img, TIFF, DefaultOptions())
}
