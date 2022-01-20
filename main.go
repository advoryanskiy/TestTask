package main

import (
	"iSoftTestTask/drivers"
	"time"
)

func main() {
	b := drivers.NewButtonDriver(nil, 7, time.Second)
	b.Pressed()
}
