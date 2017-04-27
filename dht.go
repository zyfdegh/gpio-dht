package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const pin = 4

var p rpio.Pin

func initGPIO() {
	err := rpio.Open()
	if err != nil {
		log.Fatalf("open gpio error: %v\n", err)
		return
	}
	// 	defer rpio.Close()

	p = rpio.Pin(pin)
}

// write start signal to DHT22
func initDHT22() {
	p.Output()

	// keep low for at least 10ms
	p.Low()
	time.Sleep(20 * time.Microsecond)

	// pull up for 20~40us
	p.PullUp()
	time.Sleep(30 * time.Millisecond)

}

func read() {
	initGPIO()
	initDHT22()
	// set as input and wait DHT response for 80us
	p.Input()

	for i := 0; i < 1000; i++ {
		fmt.Println(p.Read())
	}
}

func parseTH(stat rpio.State) (h, t float32, err error) {
	h, t = 23.0, 25.2
	return
}
