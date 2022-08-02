package main

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type streamer struct {
	track          *Track
	samplePerSec   int
	samplePosition int
}

var valmax = 1.0

func (str *streamer) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		time := float64(str.samplePosition) / float64(str.samplePerSec) //seconds
		val := str.track.getval(time)
		if val > valmax {
			valmax = val
			fmt.Printf("WARNING - sampler:valmax=%v\n", val)
		}
		samples[i][0] = val
		samples[i][1] = val
		str.samplePosition++
	}
	return len(samples), true
}

func (str streamer) Err() error {
	return nil
}

//Noise function that integrates the track
func createStreamer(trk *Track) beep.Streamer {
	fmt.Printf("Preparing sound stream\n")
	var str = streamer{trk, 0, 0}
	str.samplePerSec = beep.SampleRate(44100).N(time.Second)
	return &str
}

func initSpeaker() {
	fmt.Printf("Preparing speakers\n")
	//integrate with audio lib
	var sampleRate = beep.SampleRate(44100)
	var buffSize = 5 * sampleRate.N(time.Second) //buffer for 5 seconds
	speaker.Init(sampleRate, buffSize)
}

func runSampler(trk *Track) {
	speaker.Play(createStreamer(trk))
	fmt.Printf("Playing sound\n")
	select {}
}
