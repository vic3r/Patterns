package factory

// PhoneType is type definitions
type PhoneType string

const (
	// Iphone is a type of phone
	Iphone PhoneType = "iphone"
	// Samsung is a type of phone
	Samsung PhoneType = "samsung"
	// Pixel is a type of phone
	Pixel PhoneType = "pixel"
)

// NewPhone creates a new Phone
func NewPhone(phone PhoneType) Phone {
	switch phone {
	case Iphone:
		return newIphone()
	case Samsung:
		return newSamsung()
	default:
		return newPixel()
	}
}
