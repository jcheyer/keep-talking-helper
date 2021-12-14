package deps

type Bomb interface {
	GetSerial() (string, error)
}
