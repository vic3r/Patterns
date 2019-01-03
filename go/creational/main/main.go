package main

import "github.com/vic3r/patterns/go/creational/factory_method"

func main() {
	phone := factory.NewPhone(factory.Iphone)
	phone.Call(1233)

	broPhone := factory.NewPhone(factory.Pixel)
	broPhone.Call(1234)
}
