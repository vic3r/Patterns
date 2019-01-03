package factory

import "fmt"

// PixelStruct is an pixel struct
type PixelStruct struct{}

var _ Phone = &PixelStruct{}

func newPixel() Phone {
	return &PixelStruct{}
}

// Call calls to a number
func (p *PixelStruct) Call(number int) error {
	fmt.Printf("Call to %d made via an Pixel", number)
	return nil
}

// Send sends a message
func (p *PixelStruct) Send(receiver string) error {
	fmt.Printf("Message sent to: %s via an Pixel", receiver)
	return nil
}
