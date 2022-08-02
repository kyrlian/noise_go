package main

import "math"

//FILTER
type Filter struct {
	in   Signal
	low  float64
	high float64
}

func filter(s Signal, low float64, high float64) Filter {
	return Filter{s, low, high}
}
func (o Oscillator) filter(low float64, high float64) Filter {
	return filter(o, low, high)
}
func (t Track) filter(low float64, high float64) Filter {
	return filter(t, low, high)
}
func (f Filter) getval(t float64) float64 {
	var v = f.in.getval(t)
	if v < f.low {
		v = f.low
	} else if v > f.high {
		v = f.high
	}
	return v
}
func demo_filter() Signal {
	var s_dotfilter = oscillator_sf(SIN, 120).filter(1.0, 2.0)
	var s_funcfilter = filter(oscillator_sf(SIN, 120), 1.0, 2.0)
	var tr = track()
	tr.appendSignal(s_dotfilter, .0, 1.0, tf(.0))
	tr.appendSignal(s_funcfilter, .0, 1.0, tf(.0))
	var trf = tr.filter(.0, .9)
	return trf
}

var _ = demo_filter() //Avoid not used error

//INVERT - takes an input signal and gives 1/signal(t)
type Invert struct {
	in Signal
}

func (i Invert) getval(t float64) float64 {
	return 1 / i.in.getval(t)
}
func invert(s Signal) Invert {
	return Invert{s}
}
func (tf TimedFloat) invert() Invert {
	return invert(tf)
}

//POWER - takes an input signal and gives 2^signal(t)
type Power struct {
	in   Signal
	base float64
}

func (p Power) getval(t float64) float64 {
	return math.Pow(p.base, p.in.getval(t))
}
func power(s Signal, base float64) Power {
	return Power{s, base}
}
func (o Oscillator) power(base float64) Power {
	return power(o, base)
}
func (o Oscillator) power2() Power {
	return o.power(2.0)
}
func demo_power() Signal {
	var s_dotfilter = oscillator(SIN, timedFloat(120, oscillator_custom_xy([]xyPair{{.0, .0}, {5.0, 1.0}}).power2()), tf(.0), tf(.1))
	var s_funcfilter = oscillator(SIN, timedFloat(120, power(oscillator_custom_xy([]xyPair{{.0, .0}, {5.0, 1.0}}), 2.0)), tf(.0), tf(.1))
	var tr = track()
	tr.appendSignal(s_dotfilter, .0, 1.0, tf(.0))
	tr.appendSignal(s_funcfilter, .0, 1.0, tf(.0))
	return tr
}

var _ = demo_power() //Avoid not used error
