package main

import (
	"fmt"
	"math"
	"math/rand"
)

func root(x float64, n int) float64 {
	var lower = .0
	var upper = x
	var r = .0
	for upper-lower >= 0.000000001 {
		r = (upper + lower) / 2.0
		var temp = math.Pow(r, float64(n))
		if temp > x {
			upper = r
		} else {
			lower = r
		}
	}
	return r
}

var semiToneConst = root(2.0, 12) //1,05946309435929

func getSemiToneFreq(fstart float64, nsemitones int) float64 {
	var nfreq = fstart * math.Pow(semiToneConst, float64(nsemitones))
	//fmt.Printf("getFreq(%v,%v)=%v\n", fstart, nsemitones, nfreq)
	return nfreq
}

/* func simpleEnv() []xyPair {
	return []xyPair{{0.0, 0.0}, {0.1, 0.8}, {0.2, 0.4}, {0.4, 0.4}, {0.5, 0.0}}
} */

func accord3(fstart float64, gap1 int, gap2 int) Track { //3,4
	var finalTrack = Track{}
	var baseFreq = 110.0
	var oNote1 = oscillator_f(SIN, baseFreq)
	//finalTrack.elements = append(finalTrack.elements, TrackElement{&oNote1, .0, 5.0, .7, nil})
	finalTrack.append(&oNote1, .0, 5.0, tf(.7))
	var oNote2 = oscillator_f(SIN, getSemiToneFreq(baseFreq, gap1))
	finalTrack.elements = append(finalTrack.elements, TrackElement{&oNote2, .0, 5.0, tf(.6)})
	var oNote3 = oscillator_f(SIN, getSemiToneFreq(baseFreq, gap1+gap2))
	finalTrack.elements = append(finalTrack.elements, TrackElement{&oNote3, .0, 5.0, tf(.5)})
	return finalTrack
}

func accordMineur(fstart float64) Track {
	return accord3(fstart, 3, 4)
}
func accordMajeur(fstart float64) Track {
	return accord3(fstart, 4, 3)
}

func harmonics(baseFreq float64, nharmonics int) Track {
	var finalTrack = Track{}
	for i := 1; i <= nharmonics; i++ {
		var f2pi = math.Pow(2, float64(i))
		var oNote = oscillator_fpa(SIN, baseFreq * f2pi, 0, 1.0 / f2pi)
		finalTrack.elements = append(finalTrack.elements, TrackElement{&oNote, float64(i) / 5, float64(nharmonics), tf(1.0)})
	}
	return finalTrack
}

func harmonicsTuning(baseFreq float64, nharmonics int) Track {
	var finalTrack = Track{}
	var tuningTime = 10.0  //secs
	var outOfTuneMax = 2.0 //V/Oct
	var tuningSlope = timedPairList([]xyPair{{.0, 1.0}, {.9, 0.0}, {1.0, 0.0}})
	for i := 1; i <= nharmonics; i++ {
		var f2pi = math.Pow(2, float64(i))
		var outOfTuneStart = (rand.Float64() - .5) * outOfTuneMax
		var tuningLfof = oscillator(CUSTOM, tf(.1), tf(.0), tf(outOfTuneStart), tf(.0), tuningSlope)
		var n1Freq = baseFreq * f2pi
		var n2Freq = getSemiToneFreq(n1Freq, 4)
		var n3Freq = getSemiToneFreq(n1Freq, 7)
		fmt.Printf("harmonicsTuning:n1Freq:%v	,n2Freq:%v	,n3Freq:%v\n", n1Freq, n2Freq, n3Freq)
		//var oNote1 = oscillator(SIN, n1Freq).setLfo(FREQ, &tuningLfof) //custom oscillator constructor and modifier
		var oNote1 = oscillator(SIN, timedFloat(n1Freq,&tuningLfof), tf(.0), tf(1.0), nil, nil)
		var oNote2 = oscillator(SIN, timedFloat(n2Freq,&tuningLfof), tf(.0), tf(1.0), nil, nil)
		var oNote3 = oscillator(SIN, timedFloat(n3Freq,&tuningLfof), tf(.0), tf(1.0), nil, nil)
		var accAmp = tf(1.0 / f2pi / 3)
		finalTrack.append(oNote1, .0, tuningTime, accAmp) //custom append method
		finalTrack.elements = append(finalTrack.elements, TrackElement{oNote2, .0, tuningTime, accAmp})
		finalTrack.elements = append(finalTrack.elements, TrackElement{oNote3, .0, tuningTime, accAmp})
	}
	return finalTrack
}
