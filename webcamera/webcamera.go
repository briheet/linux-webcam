package webcamera

import (
	"fmt"

	"golang.org/x/sys/unix"
)

type Webcam struct {
	fd          uintptr
	buffercount uint32
	buffers     [][]byte
	streaming   bool
	pollFds     []unix.PollFd
}

func Open(name string) (*Webcam, error) {
	handle, err := unix.Open(name, unix.O_RDWR|unix.O_NONBLOCK, 0666)
	if err != nil {
		return nil, err
	}

	if handle < 0 {
		return nil, fmt.Errorf("failed to open %v", name)
	}

	success := false

	defer func() {
		if !success {
			unix.Close(handle)
		}
	}()

	fd := uintptr(handle)

	supportsVideoCamera, supportsVideoStreaming, err := checkCapabilities(fd)

	if supportsVideoCamera == false {
		return nil, fmt.Errorf("this does not support video camera")
	}

	if supportsVideoStreaming == false {
		return nil, fmt.Errorf("this does not supports video streaming")
	}

	w := new(Webcam)

	w.fd = uintptr(handle)
	w.buffercount = 256
	w.streaming = true
	w.pollFds = []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}

	success = true

	return w, err
}
