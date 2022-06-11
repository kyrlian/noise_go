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

//xyPair is a simple x,y struct
type xyPair struct {
	x float64
	y float64
}

func customShape(src [](float64), duration float64) [](xyPair) { //build an array of xy from a list of y
	var nbpoints = len(src)
	var xratio = duration / float64(nbpoints)
	var cs = [](xyPair){}
	for i := 1; i <= nbpoints; i++ {
		x := float64(i) * xratio
		cs = append(cs, xyPair{x, src[i]})
	}
	return cs
}

//var cs = customShape([](float64){.0, .8, .4, .0})

//Oscillator structure
type Oscillator struct {
	shape       int
	freq        float64
	phase       float64
	ampl        float64
	width       float64 // optional, to set the width of the top of the PULSE shape
	lfofreq     Signal  //expects V/Oct (ie -1 half the freq, +1 doubles the freq) as in: f * 2 ^ lfofreq
	lfophase    Signal
	lfoampl     Signal
	customshape [](xyPair) // for CUSTOM shape
}

//CONSTRUCTORS
func oscillator(shape int, freq float64) Oscillator {
	return Oscillator{shape, freq, .0, 1.0, .0, nil, nil, nil, nil}
}

func enveloppe(points [](xyPair)) Oscillator {
	var lastx = points[len(points)-1].x
	var freq = 1 / lastx
	return Oscillator{CUSTOM, freq, .0, 1.0, .0, nil, nil, nil, points}
}

//MODIFIERS
func (o Oscillator) setLfo(itype int, lfo Signal) Oscillator {
	switch itype {
	case FREQ:
		o.lfofreq = lfo
	case PHASE:
		o.lfophase = lfo
	case AMPL:
		o.lfoampl = lfo
	}
	return o
}

//GETERS
func (o *Oscillator) getfreq(t float64) float64 { //Freq in Hz
	var r = o.freq
	if o.lfofreq != nil {
		var lfoval = o.lfofreq.getval(t)
		var m = math.Pow(2, lfoval) //expect V/Oct : f * 2 ^ lfofreq
		r *= m
		// fmt.Printf("getfreq;base;%v;lfo;%v;mult;%v;res;%v\n", o.freq, lfoval, m, r)
	}
	return r
}

func (o *Oscillator) getphase(t float64) float64 { //phase should be 0-1, as it's added after normalisation. so .5 is a half period shift
	var r = o.phase
	if o.lfophase != nil {
		r *= o.lfophase.getval(t) //expect mult factor
	}
	return r
}

func (o *Oscillator) getampl(t float64) float64 { // amplitude should be 0-1, if it's more than 1 it will saturate - no attempt is done to normalize
	var r = o.ampl
	if o.lfoampl != nil {
		r *= o.lfoampl.getval(t) //expect mult factor
	}
	return r
}

func (o *Oscillator) getcustomval(x float64) float64 { //getval for custom shape
	if o.customshape != nil {
		var points = o.customshape //.points
		//var lastx = points[len(points)-1].x
		var previousx = points[0].x
		var previousy = points[0].y
		var xmod = x //math.Mod(x, lastx) //we loop the shape
		for _, point := range points {
			if xmod < point.x { //previousx <= xmod && xmod < point.x
				var nextx = point.x
				var nexty = point.y
				var r = previousy + (xmod-previousx)/(nextx-previousx)*(nexty-previousy)
				//fmt.Printf("getcustomval;x;%v;y;%v;\n", x, r)
				return r
			}
			previousx = point.x
			previousy = point.y
		}
	}
	return 0.0
}

func (o *Oscillator) getval(t float64) float64 {
	var f = o.getfreq(t)
	var p = 1.0 / f           //period
	var tmod = math.Mod(t, p) //O-p
	//var x = t*f + o.getphase(t)
	var xmod = math.Mod(tmod*f+o.getphase(t), 1) //O-1
	//var xmod = math.Mod(x, 1)
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
		if xmod < o.width { //with width=.5 it's just a square
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
		y = o.getcustomval(xmod)
	}
	return y * o.getampl(t) //can be negative - ex for LFOs
}
