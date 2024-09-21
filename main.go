package main

import (
	"flag"
)

const device_name = "/dev/video0"

func main() {
	device := flag.String("d", device_name, "the dir used for webcam")
	flag.Parse()

	camera, err := webcamera.Open(*device)
	if err != nil {
		panic(err.Error())
	}
}
