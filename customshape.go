package main

//xyPair is a simple x,y struct, for custom shapes
type xyPair struct {
	x float64
	y float64
}

type TimedPair struct {
	x TimedFloat
	y TimedFloat
}

//CONSTRUCTORS
func timedPair(x TimedFloat, y TimedFloat) TimedPair {
	return TimedPair{x, y}
}

func tp(x float64, y float64) TimedPair { //alias with conversion float-timedFloat
	return timedPair(tf(x), tf(y))
}

func timedPairList_xy(xylist [](xyPair)) []TimedPair { //convert list of xyPair to list of TimedPair
	var res = [](TimedPair){}
	for _, xy := range xylist {
		res = append(res, tp(xy.x, xy.y))
	}
	return res
}

func timedPairList_y(src [](float64), duration float64) [](TimedPair) { //build an array of xy from a list of y and a total duration
	var nbpoints = len(src)
	var xratio = duration / float64(nbpoints)
	var res = [](TimedPair){}
	for i := 0; i < nbpoints; i++ {
		x := float64(i) * xratio
		res = append(res, timedPair(tf(x), tf(src[i])))
	}
	return res
}

//CONSTRUCTORS for CUSTOM oscillator
func oscillator_custom_spa(freq Signal, phase Signal, ampl Signal, customshape [](TimedPair)) Oscillator {
	return oscillator_full(CUSTOM, freq, phase, ampl, nil, customshape)
}
func oscillator_custom(customshape [](TimedPair)) Oscillator {
	lastx := customshape[len(customshape)-1].x
	return oscillator_full(CUSTOM, lastx.invert(), tf(.0), tf(1.0), nil, customshape)
}
func oscillator_custom_xy(points [](xyPair)) Oscillator {
	//var lastx = points[len(points)-1].x
	//var freq = 1 / lastx
	//return oscillator_full(CUSTOM, tf(freq), tf(.0), tf(1.0), nil, timedPairList_xy(points))
	return oscillator_custom(timedPairList_xy(points))
}
func oscillator_custom_y(src [](float64), duration float64) Oscillator {
	return oscillator_custom(timedPairList_y(src,duration))
}

//GETTERS
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
