package phasor

import (
     "math"
)

type Phasor struct {
     // phase increment per tick
     phaseIncrement float64
     
     // internal phasor state
     phase float64
}

// constructor
func New(phaseIncrement float64, phase float64) Phasor {
     p := Phasor { phaseIncrement, phase }
     return p
}

// tick over one sample
func (p *Phasor) Tick() {
     // increment phase
     p.phase += p.phaseIncrement

     // stay in [0 .. 2pi]
     if p.phase > 2.0*math.Pi {
     	p.phase -= 2.0*math.Pi
     }
}

// set frequency
func (p *Phasor) SetFrequency(frequency float64, samplingFrequency float64) {
     p.phaseIncrement = 2.0*math.Pi*frequency/samplingFrequency
}

// set phasor phase increment
func (p *Phasor) SetPhaseIncrement(phaseIncrement float64) {
     p.phaseIncrement = phaseIncrement
}

// get phasor phase
func (p Phasor) GetPhase() float64 {
     return p.phase
}
