package drivers

import (
	"time"
)

type ButtonDriver struct {
	pressed bool
	interval time.Duration
	debounceCount int
	pin Pin
	r Reader
	stop chan bool
}

func NewButtonDriver(initial bool, r Reader, p Pin, i time.Duration, dCount int) *ButtonDriver {
	b := &ButtonDriver{
		pressed: initial,
		interval: i,
		debounceCount: dCount,
		pin: p,
		r: r,
		stop: make(chan bool),
	}

	return b
}

func (b *ButtonDriver) Start() {
	pressed := b.pressed
	counter := 0

	go func() {
		for {
			newVal, err := b.r.Read(b.pin)
			if err == nil { // skip errors for now
				// for debug logging
				//
				//var oldVal int
				//if pressed {
				//	oldVal = 1
				//} else {
				//	oldVal = 0
				//}
				//log.Printf("old: %d, new: %d, counter: %d\n", oldVal, newVal, counter)

				newPressed := newVal == 1
				if newPressed == pressed && counter > 0 {
					counter = 0
				}
				if newPressed != pressed {
					counter++
				}
				if counter >= b.debounceCount {
					//log.Println("UPDATING")
					pressed = newPressed
					b.pressed = pressed
					counter = 0
				}
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
