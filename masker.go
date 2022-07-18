package imagemasker

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

/* 	This function finds the first instance where one image is a
complete subset of another */
func findAbsoluteMatch(origImageFilename string, maskImageFilename string) (xPos int, yPos int, width int, length int) {
	origImg := getImage(origImageFilename)
	maskImage := getImage(maskImageFilename)

	// Grab the top-left pixel of the mask image
	maskPix := maskImage.At(0, 0)

	for i := 0; i < origImg.Bounds().Dx(); i++ {
		for j := 0; j < origImg.Bounds().Dy(); j++ {
			// Find the first instance of exact same color
			if origImg.At(i, j) == maskPix {
				//maskPix = maskImage.At
			}
		}
	}
	log.Fatalf("Function not complete.")
	return 0, 0, 0, 0
}

// reusable func for getting an image from a .png file
func getImage(filename string) image.Image {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}
	defer reader.Close()

	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal("Failed to decode image ", filename, err)
	}

	return img
}

func incrementPixelInImage(entireImage image.Image, posX *int, posY *int) {
	// Move the pixel up by one
	maxX := entireImage.Bounds().Dx()
	maxY := entireImage.Bounds().Dy()

	if *posX < maxX {
		*posX++
	} else if *posY < maxY {
		*posY++
	} else {
		*posX = maxX
		*posY = maxY
	}
}

func resetPixelCoords(posX *int, posY *int) {
	*posX = 0
	*posY = 0
}
