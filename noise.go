package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("MAKE SOME NOISE")
	initSpeaker()

	var ex = "drums"

	if ex == "simple" {
		var o2 = oscillator_f(PULSE, 440.0)
		var finalTrack = Track{[]TrackElement{{&o2, 0.0, 5.0, tf(.8)}}}
		runSampler(&finalTrack)
	}

	if ex == "majeurmineur" {
		//var envshape = []xyPair{xyPair{0.0, 0.0}, xyPair{0.1, 0.8}, xyPair{0.2, 0.4}, xyPair{0.4, 0.4}, xyPair{0.5, 0.0}}
		var envshape = customShape([](float64){.0, .8, .4, .4, .0}, .5)
		var lfoa = oscillator(CUSTOM, tf(2.0), tf(0), tf(1 / 1.3), nil, envshape)
		var lfoa2 = oscillator(CUSTOM, tf(2.0), tf(0.5), tf(1 / 1.3), nil, envshape)
		var tmajeur = accordMajeur(440.0)
		var tmineur = accordMineur(440.0)
		var finalTrack = Track{[]TrackElement{{&tmajeur, 0.0, 5, timedFloat(.8, &lfoa)}, {&tmineur, .0, 5.0, timedFloat(.8, &lfoa2)}}}
		runSampler(&finalTrack)
	}

	if ex == "enveloppe" {
		var envshape = []xyPair{{0.0, 0.0}, {0.1, 0.8}, {0.2, 0.5}, {0.4, 0.0}, {0.5, 0.0}}
		var lfoa = enveloppe(envshape)		
		var o2 = oscillator(SIN, tf(440.0), tf(0), timedFloat(0.8, &lfoa), nil, nil)
		var finalTrack = Track{[]TrackElement{{&o2, 0.0, 5.0, tf(1.0)}}}
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
		tOscs.elements = append(tOscs.elements, TrackElement{oscillator(SIN, tf(55.0), tf(0), tf(0.7), nil, nil), .0, 5.0, tf(1.0)})
		tOscs.elements = append(tOscs.elements, TrackElement{oscillator(SIN, tf(110.0), tf(0.2), tf(0.6), nil, nil), .0, 5.0, tf(1.0)})
		tOscs.elements = append(tOscs.elements, TrackElement{oscillator(NOISE, tf(0.0), tf(0.0), tf(0.1), nil, nil), .0, 5.0, tf(.05)})

		var ampl =  timedFloat(0.8, oscillator(CUSTOM, tf(2.0), tf(0), tf(1.0), nil, timedPairList([]xyPair{{0.0, 0.0}, {0.2, 0.8}, {0.3, 0.6}, {0.4, 0.1}, {1.0, 0.0}})))

		finalTrack.elements = append(finalTrack.elements, TrackElement{tOscs, .0, 5.0,ampl})
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
		var lfoaBip = oscillator(PULSE, tf(4), tf(0), tf(.9), tf(.2), nil)
		var oBip = oscillator(SIN, tf(baseFreq * f2pi), tf(0), tf(1.0 / f2pi), nil, nil)
		finalTrack.elements = append(finalTrack.elements, TrackElement{oBip, float64(i) / 5, float64(i), timedFloat(.8, lfoaBip)})
		//
		var oHighKicks = oscillator(NOISE, tf(0), tf(0), timedFloat(.8, oscillator(PULSE, tf(8), tf(0.1), tf(.2), tf(0.1), nil)), nil, nil)
		finalTrack.elements = append(finalTrack.elements, TrackElement{oHighKicks, float64(i) / 5, 5.0, tf(1)})
		//finalise
		runSampler(&finalTrack)
	}

}
