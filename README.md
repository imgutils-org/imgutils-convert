# imgutils-convert

[![Go Reference](https://pkg.go.dev/badge/github.com/imgutils-org/imgutils-convert.svg)](https://pkg.go.dev/github.com/imgutils-org/imgutils-convert)
[![Go Report Card](https://goreportcard.com/badge/github.com/imgutils-org/imgutils-convert)](https://goreportcard.com/report/github.com/imgutils-org/imgutils-convert)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for converting images between different formats. Part of the [imgutils](https://github.com/imgutils-org) collection.

## Features

- Support for JPEG, PNG, GIF, BMP, and TIFF formats
- Automatic format detection from file extension
- Configurable JPEG quality
- Simple one-liner file conversion
- Stream-based encoding for flexibility

## Installation

```bash
go get github.com/imgutils-org/imgutils-convert
```

## Quick Start

```go
package main

import (
    "log"

    "github.com/imgutils-org/imgutils-convert"
)

func main() {
    // Convert PNG to JPEG with one line
    err := convert.ConvertFile("input.png", "output.jpg", convert.Options{
        Quality: 85,
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## Usage Examples

### File Conversion

```go
// PNG to JPEG
convert.ConvertFile("image.png", "image.jpg", convert.DefaultOptions())

// JPEG to PNG
convert.ConvertFile("photo.jpg", "photo.png", convert.DefaultOptions())

// Any format to WebP-friendly PNG
convert.ConvertFile("image.bmp", "image.png", convert.DefaultOptions())
```

### Format-Specific Conversion

```go
// Open source image
file, _ := os.Open("input.png")
src, _, _ := image.Decode(file)
file.Close()

// Convert to JPEG
out, _ := os.Create("output.jpg")
convert.ToJPEG(src, out, 90)
out.Close()

// Convert to PNG
out, _ = os.Create("output.png")
convert.ToPNG(src, out)
out.Close()

// Convert to GIF
out, _ = os.Create("output.gif")
convert.ToGIF(src, out)
out.Close()
```

### Stream-Based Conversion

```go
// Convert from reader to writer
inputFile, _ := os.Open("input.png")
outputFile, _ := os.Create("output.jpg")

convert.Convert(inputFile, outputFile, convert.JPEG, convert.Options{
    Quality: 85,
})

inputFile.Close()
outputFile.Close()
```

### Detect Format from Extension

```go
format := convert.FormatFromExtension("photo.jpg")  // Returns convert.JPEG
format := convert.FormatFromExtension("image.png")  // Returns convert.PNG
format := convert.FormatFromExtension("anim.gif")   // Returns convert.GIF
```

## API Reference

### Types

#### Format

```go
type Format string

const (
    JPEG Format = "jpeg"
    PNG  Format = "png"
    GIF  Format = "gif"
    BMP  Format = "bmp"
    TIFF Format = "tiff"
)
```

#### Options

```go
type Options struct {
    Quality int // JPEG quality (1-100), default 85
}
```

### Functions

| Function | Description |
|----------|-------------|
| `ConvertFile(in, out, opts)` | Convert file to format based on extension |
| `Convert(r, w, format, opts)` | Convert from reader to writer |
| `Encode(w, img, format, opts)` | Encode image to specific format |
| `Decode(r)` | Decode image from reader |
| `FormatFromExtension(path)` | Detect format from file extension |
| `ToJPEG(img, w, quality)` | Convert to JPEG |
| `ToPNG(img, w)` | Convert to PNG |
| `ToGIF(img, w)` | Convert to GIF |
| `ToBMP(img, w)` | Convert to BMP |
| `ToTIFF(img, w)` | Convert to TIFF |
| `DefaultOptions()` | Returns default options (quality 85) |

## Supported Formats

| Format | Extension(s) | Read | Write |
|--------|--------------|------|-------|
| JPEG | .jpg, .jpeg | Yes | Yes |
| PNG | .png | Yes | Yes |
| GIF | .gif | Yes | Yes |
| BMP | .bmp | Yes | Yes |
| TIFF | .tiff, .tif | Yes | Yes |

## Requirements

- Go 1.16 or later

## Related Packages

- [imgutils-compress](https://github.com/imgutils-org/imgutils-compress) - Image compression
- [imgutils-resize](https://github.com/imgutils-org/imgutils-resize) - Image resizing
- [imgutils-sdk](https://github.com/imgutils-org/imgutils-sdk) - Unified SDK

## License

MIT License - see [LICENSE](LICENSE) for details.
