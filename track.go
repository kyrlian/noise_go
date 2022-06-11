package main

import "math"

//TrackElement with a start and end
type TrackElement struct {
	signal Signal
	start  float64
	end    float64
	ampl   float64
	lfoa   Signal
}

//Track is a collection of track elements
type Track struct {
	elements [](TrackElement) //slice
}

func (tre TrackElement) setlfoa(s Signal) {
	tre.lfoa = s
}

func (tr Track) append(signal Signal, start float64, end float64, ampl float64, lfoa Signal) Track {
	tr.elements = append(tr.elements, TrackElement{signal, start, end, ampl, lfoa})
	return tr
}

func (tr Track) getval(t float64) float64 {
	var r = 0.0
	//fmt.Printf("tr.elements.len: %v\n", len(tr.elements))
	for _, elem := range tr.elements {
		if t >= elem.start && t < elem.end { //this signal is active
			var subt = t - elem.start
			v := elem.signal.getval(subt)
			fa := elem.ampl
			if elem.lfoa != nil { //add lfoa
				fa *= elem.lfoa.getval(subt) //lfo is multiplies base amplitude
			}
			r += v * math.Max(0, fa)
			//fmt.Printf("tr.getval.t,r: %v , %v\n", t, r)
		}
	}

	return r
}
