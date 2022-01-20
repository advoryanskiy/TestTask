package drivers

type Pin uint8

type Reader interface {
	Read(Pin) (val int, err error)
}

type Driver interface {
	Start()
	Stop()
}
