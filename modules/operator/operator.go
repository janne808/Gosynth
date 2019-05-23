package operator

import (
    "gosynth/modules/phasor"
    "math"
)

type Operator struct {
     p phasor.Phasor
     frequency float64
     phaseModulation float64
}

// constructor
func New(phaseModulation float64, frequency float64, sampleRate float64) Operator {
     // instantiate phasor
     p := phasor.New(0.0, 0.0)
     
     // instantiate operator
     op := Operator { p: p, frequency: frequency, phaseModulation: phaseModulation }

     // init phasor frequency
     op.p.SetFrequency(frequency, sampleRate)
     
     return op
}

func (op *Operator) Tick() {
     op.p.Tick()
}

func (op *Operator) SetFrequency(frequency float64, sampleRate float64) {
     op.p.SetFrequency(frequency, sampleRate)
}

func (op *Operator) SetPhaseModulation(phaseModulation float64) {
     op.phaseModulation = phaseModulation
}

func (op Operator) GetOutput() float64 {
     return math.Sin(op.p.GetPhase()+op.phaseModulation)
}
