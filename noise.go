package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("MAKE SOME NOISE")
	initSpeaker()

	var ex = "harmonics"

	if ex == "simple" {
		var o2 = Oscillator{PULSE, 440.0, 0, 1, .1, nil, nil, nil, nil}
		var finalTrack = Track{[]TrackElement{TrackElement{&o2, 0.0, 5.0, .8, nil}}}
		runSampler(&finalTrack)
	}

	if ex == "majeurmineur" {
		//var envshape = []xyPair{xyPair{0.0, 0.0}, xyPair{0.1, 0.8}, xyPair{0.2, 0.4}, xyPair{0.4, 0.4}, xyPair{0.5, 0.0}}
		var envshape = customShape([](float64){.0, .8, .4, .4, .0}, .5)
		var lfoa = Oscillator{CUSTOM, 2.0, 0, 1 / 1.3, 0, nil, nil, nil, envshape}
		var lfoa2 = Oscillator{CUSTOM, 2.0, 0.5, 1 / 1.3, 0, nil, nil, nil, envshape}
		var tmajeur = accordMajeur(440.0)
		var tmineur = accordMineur(440.0)
		var finalTrack = Track{[]TrackElement{TrackElement{&tmajeur, 0.0, 5, .8, &lfoa}, TrackElement{&tmineur, .0, 5.0, .8, &lfoa2}}}
		runSampler(&finalTrack)
	}

	if ex == "enveloppe" {
		var envshape = []xyPair{xyPair{0.0, 0.0}, xyPair{0.1, 0.8}, xyPair{0.2, 0.5}, xyPair{0.4, 0.0}, xyPair{0.5, 0.0}}
		var lfoa = Oscillator{CUSTOM, 2.0, 0, 1, 0, nil, nil, nil, envshape}
		var o2 = Oscillator{SIN, 440.0, 0, 0.8, 0, nil, nil, &lfoa, nil}
		var finalTrack = Track{[]TrackElement{TrackElement{&o2, 0.0, 5.0, 1.0, nil}}}
		runSampler(&finalTrack)
	}

	if ex == "harmonics" {
		//var finalTrack = harmonics(22.5, 10)
		//var finalTrack = harmonicsTuning(1.71875, 10)
		var finalTrack = harmonicsTuning(22.5, 10)

		runSampler(&finalTrack)
	}

	if ex == "drums" {
		var finalTrack = Track{}
		//var oHighKicks = Oscillator{NOISE, 0, 0, 0, 0, nil, nil, &Oscillator{CUSTOM, 2.0, 0, 1, 0, nil, nil, nil, &CustomShape{[]xyPair{xyPair{0.0, 0.8}, xyPair{0.3, 0.1}, xyPair{1.0, 0.0}}}}, nil}
		//finalTrack.elements = append(finalTrack.elements, TrackElement{&oHighKicks, .0, 5.0, 1.0, nil})

		var tOscs = Track{}
		tOscs.elements = append(tOscs.elements, TrackElement{&Oscillator{SIN, 55.0, 0, 0.7, 0.1, nil, nil, nil, nil}, .0, 5.0, 1.0, nil})
		tOscs.elements = append(tOscs.elements, TrackElement{&Oscillator{SIN, 110.0, 0.2, 0.6, 0.1, nil, nil, nil, nil}, .0, 5.0, 1.0, nil})
		tOscs.elements = append(tOscs.elements, TrackElement{&Oscillator{NOISE, 0.0, 0.0, 0.1, 0.0, nil, nil, nil, nil}, .0, 5.0, .05, nil})

		finalTrack.elements = append(finalTrack.elements, TrackElement{&tOscs, .0, 5.0, 0.8, &Oscillator{CUSTOM, 2.0, 0, 1.0, 0, nil, nil, nil, []xyPair{xyPair{0.0, 0.0}, xyPair{0.2, 0.8}, xyPair{0.3, 0.6}, xyPair{0.4, 0.1}, xyPair{1.0, 0.0}}}})
		//finalTrack.elements = append(finalTrack.elements, TrackElement{&tOscs, .0, 5.0, 1.0, nil})

		runSampler(&finalTrack)
	}

	if ex == "combined1" {
		//intro
		var baseFreq = 22.5
		var finalTrack = harmonics(baseFreq, 6)
		//bip
		var i = 7
		var f2pi = math.Pow(2, float64(i))
		var lfoaBip = Oscillator{PULSE, 4, 0, .9, .2, nil, nil, nil, nil}
		var oBip = Oscillator{SIN, baseFreq * f2pi, 0, 1.0 / f2pi, 0, nil, nil, nil, nil}
		finalTrack.elements = append(finalTrack.elements, TrackElement{&oBip, float64(i) / 5, float64(i), .8, &lfoaBip})
		//
		var oHighKicks = Oscillator{NOISE, 0, 0, 0, 0, nil, nil, &Oscillator{PULSE, 8, 0.1, .2, 0.1, nil, nil, nil, nil}, nil}
		finalTrack.elements = append(finalTrack.elements, TrackElement{&oHighKicks, float64(i) / 5, 5.0, 1, nil})
		//finalise
		runSampler(&finalTrack)
	}

}
