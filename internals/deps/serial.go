package deps

type Serial interface {
	SetSerial(s string) error
	IsEven() (bool, error)
	IsValid() bool
}
