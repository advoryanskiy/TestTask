package main

import (
	"iSoftTestTask/drivers"
	"testing"
	"time"
)

type TestReader struct {
	vals []int
	i int
}

func NewTestReader(vals []int) *TestReader {
	return &TestReader{
		vals: vals,
		i: 0,
	}
}

func (t *TestReader) Read(pin drivers.Pin) (int, error) {
	val := t.vals[t.i]
	if t.i < len(t.vals) - 1 {
		t.i = t.i + 1
	}

	return val, nil
}

func (t *TestReader) I() int {
	return t.i
}

func TestButtonDriverPressed(t *testing.T) {
	seq := []int{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1}
	assertPressed(seq, true, t)
}

func TestButtonDriverUnpressed(t *testing.T) {
	seq := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0}
	assertPressed(seq, false, t)
}

func TestButtonDriverContactBounds(t *testing.T) {
	seq := []int {0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0}
	assertPressed(seq, false, t)
}

func assertPressed(seq []int, expected bool, t *testing.T) {
	initial := seq[0] == 1
	r := NewTestReader(seq)
	b := drivers.NewButtonDriver(initial, r, 7, time.Millisecond * 5, 5)
	b.Start()

	for i := 0 ; i < len(seq) - 3; i++ {
		if initial != expected && b.Pressed() == expected {
			t.Fail()
		}

		time.Sleep(time.Millisecond * 5)
	}

	time.Sleep(time.Millisecond * 30)

	if b.Pressed() != expected {
		t.Fail()
	}

	b.Stop()
}
