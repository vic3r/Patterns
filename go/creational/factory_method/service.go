package factory

//Phone is a phone interface
type Phone interface {
	Call(int) error
	Send(string) error
}
