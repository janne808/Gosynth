package phasor

import (
     "math"
)

type phasor struct {
     // phase increment per tick
     phaseIncrement float64
     
     // internal phasor state
     phase float64
}

// constructor
func New(phaseIncrement float64, phase float64) phasor {
     p := phasor { phaseIncrement, phase }
     return p
}

// tick over one sample
func (p *phasor) Tick() {
     // increment phase
     p.phase += p.phaseIncrement

     // stay in [0 .. 2pi]
     if p.phase > 2.0*math.Pi {
     	p.phase -= 2.0*math.Pi
     }
}

// convert frequency
func (p phasor) FrequencyToPhaseIncrement(frequency float64, samplingFrequency float64) float64 {
     return 2.0*math.Pi*frequency/samplingFrequency
}

// set phasor phase increment
func (p *phasor) SetPhaseIncrement(phaseIncrement float64) {
     p.phaseIncrement = phaseIncrement
}

// get phasor phase
func (p phasor) GetPhase() float64 {
     return p.phase
}
