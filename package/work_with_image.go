package manager

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"log"
	"os"
)

func Work_from_number(number, file_RGB string) {
	// number = "3742692872024071707132178101814610325487376049613852002170525879949807192177"
}

func Work_from_image(image_path, file_RGB string) {
	file, err := os.Open(image_path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	imageSize := img.Bounds().Max.X * img.Bounds().Max.Y
	channels := map[string][]uint8{
		"red":   make([]uint8, imageSize),
		"green": make([]uint8, imageSize),
		"blue":  make([]uint8, imageSize),
	}

	n := 0
	for Y := 0; Y < img.Bounds().Max.Y; Y++ {
		for X := 0; X < img.Bounds().Max.X; X++ {
			col := img.At(X, Y).(color.NRGBA)
			channels["red"][n] = col.R
			channels["green"][n] = col.G
			channels["blue"][n] = col.B
			n++
		}
	}

	fmt.Println("Work from image: Done.")
}
