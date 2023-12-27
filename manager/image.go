package manager

import (
	"encoding/json"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"fmt"
	"os"

	"phil.com/types"
)

func Get_image_channels(old_image_path, image_RGB string) types.Image {
	file, err := os.Open(old_image_path)
	Err_check(err)

	defer file.Close()

	img, _, err := image.Decode(file)
	Err_check(err)

	var img_data types.Image

	n := 0
	for Y := 0; Y < img.Bounds().Max.Y; Y++ {
		for X := 0; X < img.Bounds().Max.X; X++ {
			col := img.At(X, Y).(color.NRGBA)

			img_data.Red = append(img_data.Red, []byte(fmt.Sprint(col.R))...)
			img_data.Green = append(img_data.Green, []byte(fmt.Sprint(col.G))...)
			img_data.Blue = append(img_data.Blue, []byte(fmt.Sprint(col.B))...)

			n++
		}
	}

	return img_data
}

func Save_image_channels(img_data types.Image, image_RGB string) {
	img_data_bytes, err := json.Marshal(img_data)
	Err_check(err)

	file, err := os.OpenFile(image_RGB, os.O_WRONLY, os.ModePerm)
	Err_check(err)

	_, err = file.Write(img_data_bytes)
	Err_check(err)

	err = file.Close()
	Err_check(err)
}

func Load_image_channels(image_RGB string) types.Image {
	file, err := os.OpenFile(image_RGB, os.O_RDONLY, os.ModePerm)
	Err_check(err)

	var img_data types.Image

	img_data_bytes, err := io.ReadAll(file)
	Err_check(err)

	err = json.Unmarshal(img_data_bytes, &img_data)
	Err_check(err)

	err = file.Close()
	Err_check(err)

	return img_data
}
