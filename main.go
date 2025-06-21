package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/ARXXIII/ascii-artist/ascii"
	"github.com/ARXXIII/ascii-artist/load"
)

func main() {
	var imagePath string

	fmt.Print("URL or path: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		imagePath = scanner.Text()
	}
	if imagePath == "" {
		fmt.Println("Путь не введён.")
		return
	}

	var img interface{}
	var err error

	if strings.HasPrefix(imagePath, "http://") || strings.HasPrefix(imagePath, "https://") {
		img, err = load.LoadImageFromURL(imagePath)
	} else {
		img, err = load.LoadImage(imagePath)
	}

	if err != nil {
		panic(err)
	}

	ascii.PrintASCII(img.(image.Image), 80)
}
