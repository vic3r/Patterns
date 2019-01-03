package factory

import "fmt"

// SamsungStruct is an samsung struct
type SamsungStruct struct{}

var _ Phone = &SamsungStruct{}

func newSamsung() Phone {
	return &SamsungStruct{}
}

// Call calls to a number
func (p *SamsungStruct) Call(number int) error {
	fmt.Printf("Call to %d made via an Samsung", number)
	return nil
}

// Send sends a message
func (p *SamsungStruct) Send(receiver string) error {
	fmt.Printf("Message sent to: %s via an Samsung", receiver)
	return nil
}
