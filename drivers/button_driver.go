package drivers

import (
	"time"
)

type ButtonDriver struct {
	pressed bool
	interval time.Duration
	pin Pin
	r Reader
	stop chan bool
}

func NewButtonDriver(r Reader, p Pin, i time.Duration) *ButtonDriver {
	b := &ButtonDriver{
		pressed: false,
		interval: i,
		pin: p,
		r: r,
		stop: make(chan bool),
	}

	return b
}

func (b *ButtonDriver) Start() {
	pressed := b.pressed

	go func() {
		for {
			newVal, _ := b.r.Read(b.pin) // skip errors for now
			newPressed := newVal == 1
			if newPressed != pressed {
				pressed = newPressed
				b.pressed = pressed
			}

			select {
			case <- time.After(b.interval):
			case <- b.stop:
				return
			}
		}
	}()
}

func (b *ButtonDriver) Stop() {
	b.stop <- true
}

func (b *ButtonDriver) Pressed() bool {
	return b.pressed
}
