package main

//xyPair is a simple x,y struct, for custom shapes
type xyPair struct {
	x float64
	y float64
}

type TimedFloat struct {
	base float64
	lfo  Signal //expects V/Oct (ie -1 half the freq, +1 doubles the freq) as in: f * 2 ^ lfofreq
}

type TimedPair struct {
	x TimedFloat
	y TimedFloat
}

//CONSTRUCTORS
func timedFloat(base float64, lfo Signal) TimedFloat {
	return TimedFloat{base, lfo}
}

func tf(base float64) TimedFloat { //short alias
	return timedFloat(base, nil)
}

func timedPair(x TimedFloat, y TimedFloat) TimedPair {
	return TimedPair{x, y}
}

func tp(x float64, y float64) TimedPair {//alias with conversion float-timedFloat
	return timedPair(tf(x), tf(y))
}

func timedPairList(xylist [](xyPair)) []TimedPair { //convert list of xyPair to list of TimedPair
	var res = [](TimedPair){}
	for _, xy := range xylist {
		res = append(res, tp(xy.x, xy.y))
	}
	return res
}

func customShape(src [](float64), duration float64) [](TimedPair) { //build an array of xy from a list of y and a total duration
	var nbpoints = len(src)
	var xratio = duration / float64(nbpoints)
	var res = [](TimedPair){}
	for i := 1; i <= nbpoints; i++ {
		x := float64(i) * xratio
		res = append(res, timedPair(tf(x), tf(src[i])))
	}
	return res
}

// Special enveloppe constructor
func enveloppe(points [](xyPair)) OscillatorType {
	var lastx = points[len(points)-1].x
	var freq = 1 / lastx
	return oscillator(CUSTOM, tf(freq), tf(.0), tf(1.0), nil, timedPairList(points))
}

//GETTERS
func (tf TimedFloat) getval(t float64) float64 {
	r := tf.base
	if tf.lfo != nil {
		r *= tf.lfo.getval(t)
	}
	return r
}

func getcustomshapeval(points []TimedPair, xmod float64, t float64) float64 { //getval for custom shape
	//var lastx = points[len(points)-1].x.getval(t)
	var previousx = points[0].x.getval(t)
	var previousy = points[0].y.getval(t)
	//var xmod = math.Mod(t, lastx) //we loop the shape
	for _, point := range points {
		var nextx = point.x.getval(t)
		var nexty = point.y.getval(t)
		if xmod < nextx { //previousx <= xmod && xmod < point.x
			var r = previousy + (xmod-previousx)/(nextx-previousx)*(nexty-previousy)
			//fmt.Printf("getcustomval;x;%v;y;%v;\n", x, r)
			return r
		}
		previousx = nextx
		previousy = nexty
	}
	return 0.0
}
