package load

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)

func LoadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func LoadImageFromURL(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	switch {
	case strings.Contains(contentType, "jpeg"), strings.Contains(contentType, "jpg"):
		return jpeg.Decode(resp.Body)
	case strings.Contains(contentType, "png"):
		return png.Decode(resp.Body)
	default:
		img, _, err := image.Decode(resp.Body)
		return img, err
	}
}
