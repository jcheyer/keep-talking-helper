package bomb

import "github.com/jcheyer/keep-talking-helper/internals/serial"

func (b *bomb) AddSerial(s string) error {
	ser, err := serial.New(serial.WithSerial(s))
	if err != nil {
		return err
	}

	b.serial = ser
	return nil
}
