package factory

import "fmt"

// IphoneStruct is an iphone struct
type IphoneStruct struct{}

var _ Phone = &IphoneStruct{}

func newIphone() Phone {
	return &IphoneStruct{}
}

// Call calls to a number
func (p *IphoneStruct) Call(number int) error {
	fmt.Printf("Call to %d made via an Iphone", number)
	return nil
}

// Send sends a message
func (p *IphoneStruct) Send(receiver string) error {
	fmt.Printf("Message sent to: %s via an Iphone", receiver)
	return nil
}
