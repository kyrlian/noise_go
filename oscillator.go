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
)

//Oscillator structure
type OscillatorType struct {
	shape       int
	freq        Signal
	phase       Signal
	ampl        Signal
	width       Signal        // optional, to set the width of the top of the PULSE shape
	customshape [](TimedPair) // for CUSTOM shape
}

//CONSTRUCTORS
func oscillator(shape int, freq Signal, phase Signal, ampl Signal, width Signal, customshape [](TimedPair)) OscillatorType {
	return OscillatorType{shape, freq, phase, ampl, width, customshape}
}
func oscillator_f(shape int, freq float64) OscillatorType {
	return oscillator(shape, tf(freq), tf(.0), tf(1.0), tf(.0), nil)
}
func oscillator_fpa(shape int, freq float64, phase float64, ampl float64) OscillatorType {
	return oscillator(shape, tf(freq), tf(phase), tf(ampl), tf(.0), nil)
}

//GETERS
func (o OscillatorType) getval(t float64) float64 {
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
