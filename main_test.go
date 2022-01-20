package main

import (
	"iSoftTestTask/drivers"
	"math/rand"
	"testing"
	"time"
)

// threshold is voltage percentage to represent 1 or 0 logical value
// calculated as 1.19 / 3.3, where 3.3 - max voltage for pin, 1.19 - threshold voltage
const threshold = 0.36

type TestReader struct {}

func (t *TestReader) Read(pin drivers.Pin) (int, error) {
	if rand.Float32() > threshold {
		return 1, nil
	} else {
		return 0, nil
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestButtonDriverStateRead(t *testing.T) {
	r := &TestReader{}
	b := drivers.NewButtonDriver(r, 7, time.Millisecond * 10)
	b.Start()

	for i := 0; i < 10; i++ {
		if b.Pressed() {
			t.Log("Button PRESSED")
		} else {
			t.Log("Button NOT PRESSED")
		}

		time.Sleep(time.Millisecond * 20)
	}

	b.Stop()
}

func TestReaderSequence(t *testing.T) {
	r := &TestReader{}

	for i := 0; i < 10; i++ {
		if v, _ := r.Read(7); v == 1 {
			t.Log("Pin ON")
		} else {
			t.Log("Pin OFF")
		}
	}
}
