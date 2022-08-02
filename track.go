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

func (tr Track) append(signal Signal, start float64, end float64, ampl Signal) Track {
	tr.elements = append(tr.elements, TrackElement{signal, start, end, ampl})
	return tr
}

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
