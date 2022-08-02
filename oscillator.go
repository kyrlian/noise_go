package main

import (
	"math"
	"math/rand"
)

//type shapeType int

//shape type
const (
	SIN int = iota
	SQR
	TRI
	SAW
	ISAW
	FLAT
	NOISE
	PULSE
	CUSTOM
)

//lfo type
const (
	FREQ int = iota
	PHASE
	AMPL
	WIDTH
)

//Oscillator structure
type Oscillator struct {
	shape       int
	freq        Signal
	phase       Signal
	ampl        Signal
	width       Signal        // optional, to set the width of the top of the PULSE shape
	customshape [](TimedPair) // for CUSTOM shape
}

//CONSTRUCTORS
func oscillator_full(shape int, freq Signal, phase Signal, ampl Signal, width Signal, customshape [](TimedPair)) Oscillator {
	return Oscillator{shape, freq, phase, ampl, width, customshape}
}
func oscillator(shape int, freq Signal, phase Signal, ampl Signal) Oscillator {
	return oscillator_full(shape, freq, phase, ampl, tf(.5), nil)
}
func oscillator_sf(shape int, freq float64) Oscillator {
	return oscillator(shape, tf(freq), tf(.0), tf(1.0))
}
func oscillator_sfpa(shape int, freq float64, phase float64, ampl float64) Oscillator {
	return oscillator(shape, tf(freq), tf(phase), tf(ampl))
}

//CONSTRUCTORS dedicated to shapes
func oscillator_pulse(freq Signal, phase Signal, ampl Signal, width Signal) Oscillator {
	return oscillator_full(PULSE, freq, phase, ampl, width, nil)
}
func oscillator_noise(ampl Signal) Oscillator {
	return oscillator(NOISE, tf(1.0), tf(1.0), ampl)
}

// See customshape.go for CUSTOM constructors

//SETTERS
func (o Oscillator) set(elem int, s Signal) Oscillator {
	switch elem {
	case FREQ:
		o.freq = s
	case PHASE:
		o.phase = s
	case AMPL:
		o.ampl = s
	case WIDTH:
		o.width = s
	}
	return o
}

//GETERS
func (o Oscillator) getval(t float64) float64 {
	var f = o.freq.getval(t)
	var p = 1.0 / f                                  //period
	var tmod = math.Mod(t, p)                        //O-p
	var xmod = math.Mod(tmod*f+o.phase.getval(t), 1) //O-1
	//fmt.Printf("Oscillator:getval:x:%v,xmod:%v\n", x, xmod)
	//switch shape
	var y = 0.0
	switch o.shape {
	case SIN:
		y = math.Sin(2.0 * math.Pi * xmod)
	case FLAT:
		y = 1.0
	case SQR:
		if xmod < .5 {
			y = 1.0
		} else {
			y = -1.0
		}
	case PULSE:
		if xmod < o.width.getval(t) { //with width=.5 it's just a square
			y = 1.0
		} else {
			y = -1.0
		}
	case SAW: //ramp up
		y = -1.0 + 2*xmod
	case ISAW: //ramp down
		y = 1.0 - 2*xmod
	case TRI: //ramp up and down
		if xmod < .5 {
			y = -1.0 + 4.0*xmod
		} else {
			y = 3.0 - 4.0*xmod
		}
	case NOISE:
		y = rand.Float64()
	case CUSTOM:
		if o.customshape != nil {
			y = getcustomshapeval(o.customshape, xmod, t)
		}
	}
	return y * o.ampl.getval(t) //can be negative - ex for LFOs
}
