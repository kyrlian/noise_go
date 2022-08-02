package main

import "math"

//TrackElement with a start and end
type TrackElement struct {
	signal Signal
	start  float64
	end    float64
	ampl   Signal //base+lfoa
}

//Track is a collection of track elements
type Track struct {
	elements [](TrackElement) //slice
}

//CONSTRUCTORS
func trackElement(signal Signal, start float64, end float64, ampl Signal) TrackElement {
	if( end < start) { end = start }
	return  TrackElement{signal, start, end, ampl}
}
func track() Track {
	return Track{}
}

//MODIFIERS
func (trackpointer *Track) appendSignal(signal Signal, start float64, end float64, ampl Signal) *Track { //pass pointer to modify in place
	trackpointer.elements = append(trackpointer.elements, trackElement(signal, start, end, ampl))
	return trackpointer //return the pointer, the given track has been modified
}

//GETTERS
func (tr Track) getval(t float64) float64 {
	var r = 0.0
	//fmt.Printf("tr.elements.len: %v\n", len(tr.elements))
	for _, elem := range tr.elements {
		if t >= elem.start && t < elem.end { //this signal is active
			var subt = t - elem.start
			v := elem.signal.getval(subt)
			a := elem.ampl.getval(subt)
			r += v * math.Max(0, a)
			//fmt.Printf("tr.getval.t,r: %v , %v\n", t, r)
		}
	}
	return r
}
