package imagemasker

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

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

func findAbsoluteMatch(origImage string, maskImage string) {
	fmt.Println(getImage(origImage).Bounds().Max)
	fmt.Println(getImage(maskImage).Bounds().Max)
}

func main() {
	//subjectFile := os.Args[1]
	//maskFile := os.Args[2]

}
