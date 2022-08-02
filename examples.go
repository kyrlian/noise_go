package main

import "math"

func example_simple() Track {
	var o2 = oscillator_sf(PULSE, 440.0)
	var finalTrack = Track{[]TrackElement{{&o2, 0.0, 5.0, tf(.8)}}}
	return finalTrack
}

var _ = example_simple() //Avoid not used error

func example_majeurmineur() Track {
	//var envshape = []xyPair{xyPair{0.0, 0.0}, xyPair{0.1, 0.8}, xyPair{0.2, 0.4}, xyPair{0.4, 0.4}, xyPair{0.5, 0.0}}
	var envshape = timedPairList_y([](float64){.0, .8, .4, .4, .0}, .5)
	//var lfoa = oscillator_custom_spa(tf(2.0), tf(0), tf(1/1.3), envshape).set(AMPL,tf(1/1.3))
	var lfoa = oscillator_custom_y([](float64){.0, .8, .4, .4, .0}, .5).set(AMPL,tf(1/1.3))
	var lfoa2 = oscillator_custom_spa(tf(2.0), tf(0.5), tf(1/1.3), envshape)
	var tmajeur = accordMajeur(440.0)
	var tmineur = accordMineur(440.0)
	var finalTrack = Track{[]TrackElement{{&tmajeur, 0.0, 5, timedFloat(.8, &lfoa)}, {&tmineur, .0, 5.0, timedFloat(.8, &lfoa2)}}}
	return finalTrack
}

var _ = example_majeurmineur() //Avoid not used error

func example_enveloppe() Track {
	var envshape = []xyPair{{0.0, 0.0}, {0.1, 0.8}, {0.2, 0.5}, {0.4, 0.0}, {0.5, 0.0}}
	var lfoa = oscillator_custom_xy(envshape)
	var o2 = oscillator(SIN, tf(440.0), tf(0), timedFloat(0.8, &lfoa))
	var finalTrack = Track{[]TrackElement{{&o2, 0.0, 5.0, tf(1.0)}}}
	return finalTrack
}

var _ = example_enveloppe() //Avoid not used error

func example_harmonics() Track {
	//var finalTrack = harmonics(22.5, 10)
	//var finalTrack = harmonicsTuning(1.71875, 10)
	var finalTrack = harmonicsTuning(22.5, 10)
	return finalTrack
}

var _ = example_harmonics() //Avoid not used error

func example_drums() Track {
	var finalTrack = track()
	//var oHighKicks = Oscillator{NOISE, 0, 0, 0, 0, nil, nil, &Oscillator{CUSTOM, 2.0, 0, 1, 0, nil, nil, nil, &CustomShape{[]xyPair{xyPair{0.0, 0.8}, xyPair{0.3, 0.1}, xyPair{1.0, 0.0}}}}, nil}
	//finalTrack.appendSignal(&oHighKicks, .0, 5.0, 1.0, nil)

	var tOscs = track()
	tOscs.appendSignal(oscillator(SIN, tf(55.0), tf(0), tf(0.7)), .0, 5.0, tf(1.0)) //appendSignal modifies the track itself
	tOscs.appendSignal(oscillator(SIN, tf(110.0), tf(0.5), tf(0.6)), .0, 5.0, tf(1.0))
	tOscs.appendSignal(oscillator_noise(tf(0.1)), .0, 5.0, tf(.05))
	//tOscs.appendSignal(oscillator(NOISE, tf(0.0), tf(0.0), tf(0.1), nil, nil), .0, 5.0, tf(.05))

	var enveloppe = timedPairList_xy([]xyPair{{0.0, 0.0}, {0.2, 0.8}, {0.3, 0.6}, {0.4, 0.1}, {1.0, 0.0}})//Envelope of the hit
	var fslope = oscillator_custom_xy([]xyPair{{0.0, 1.0},  {1.0, 2.0}})//frequence of the hits
	var ampl = timedFloat(0.8, oscillator_custom_spa(timedFloat(1.0,fslope), tf(0), tf(1.0), enveloppe))//global enveloppe uses the hit enveloppe, but with a viariable repetition frequency

	finalTrack.appendSignal(tOscs, .0, 5.0, ampl)
	//finalTrack.appendSignal(tOscs, .0, 5.0, ampl)

	return finalTrack
}

var _ = example_drums() //Avoid not used error

func example_combined1() Track {
	//intro
	var baseFreq = 22.5
	var finalTrack = harmonics(baseFreq, 6)
	//bip
	var i = 7
	var f2pi = math.Pow(2, float64(i))
	//var lfoaBip = oscillator(PULSE, tf(4), tf(0), tf(.9), tf(.2), nil)
	var lfoaBip = oscillator_pulse(tf(4), tf(0), tf(.9), tf(.2))

	var oBip = oscillator(SIN, tf(baseFreq*f2pi), tf(0), tf(1.0/f2pi))
	finalTrack.appendSignal(oBip, float64(i) / 5, float64(i), timedFloat(.8, lfoaBip))
	//
	var oHighKicks = oscillator(NOISE, tf(0), tf(0), timedFloat(.8, oscillator_pulse(tf(8), tf(0.1), tf(.2), tf(0.1))))
	finalTrack.appendSignal(oHighKicks, float64(i) / 5, 5.0, tf(1))
	//finalise
	return finalTrack
}

var _ = example_combined1() //Avoid not used error
