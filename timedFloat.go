package main

type TimedFloat struct {
	base float64
	lfo  Signal //expects V/Oct (ie -1 half the freq, +1 doubles the freq) as in: f * 2 ^ lfofreq
}

//CONSTRUCTORS
func timedFloat(base float64, lfo Signal) TimedFloat {
	return TimedFloat{base, lfo}
}

func tf(base float64) TimedFloat { //short alias
	return timedFloat(base, nil)
}

//GETTERS
func (tf TimedFloat) getval(t float64) float64 {
	r := tf.base
	if tf.lfo != nil {
		r *= tf.lfo.getval(t)
	}
	return r
}
