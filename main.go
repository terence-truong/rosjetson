package main

import (
	"log"
	// "os/exec"
	// "strconv"

	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
	"periph.io/x/conn/v3/gpio"
)

func main() {

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	en := gpioreg.ByName("GPIO20")
	if en == nil {
		log.Fatal("Failed to find GPIO20")
	}

	in1 := gpioreg.ByName("GPIO14")
	if in1 == nil {
		log.Fatal("Failed to find GPIO14")
	}

	in2 := gpioreg.ByName("GPIO16")
	if in2 == nil {
		log.Fatal("Failed to find GPIO14")
	}

	en.PWM(gpio.DutyMax, 10*physic.Hertz)
	in1.Out(gpio.High)
	in2.Out(gpio.Low)
}