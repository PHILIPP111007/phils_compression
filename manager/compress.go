package manager

import (
	"fmt"

	"phil.com/boyermoore"
	"phil.com/types"
)

var (
	channel, generated_numbers types.Channel
	channel_size               int
)

func Compress_image_channels(gen_numbers types.Channel, img_data types.Image) {
	generated_numbers = gen_numbers

	channel = img_data.Red
	compress_channel("red")

	// channel = img_data.Green
	// compress_channel("green", img_data.Green)

	// channel = img_data.Blue
	// compress_channel("blue", img_data.Blue)
}

func compress_channel(key string) {

	for {
		channel_size = len(channel)
		if channel_size == 0 {
			break
		}

		n, _ := recursion(1, -1)
		fmt.Println(n, channel_size)

		channel = channel[n:]
	}
}

func recursion(n, index int) (N int, Index int) {
	var index_new = finding_1(channel[:n])

	if index_new != -1 && n < channel_size {
		return recursion(n+1, index_new)
	}
	return n, index
}

func finding_1(channel types.Channel) int {
	return boyermoore.Index(generated_numbers, channel)
}
