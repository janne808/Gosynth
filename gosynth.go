package main

import (
    "gosynth/modules/phasor"
    "gosynth/io/alsa"
    "math"
    "fmt"
    "os"
)

const sampleRate = 44100
const bufferSize = 256

func main() {
     // instantiate ALSA driver
     driver, error := alsa.NewDriver(sampleRate, 2, 2, bufferSize)
     if error != nil {
     	fmt.Println(error)
     	os.Exit(1)
     }
     fmt.Println("[+] ALSA driver initialized")

     // audio synthesis buffer
     data := make([]byte, 2*2*bufferSize)

     // modules
     carrier := phasor.New(0.0, 0.0)
     carrier.SetPhaseIncrement(carrier.FrequencyToPhaseIncrement(60.0, float64(sampleRate)))
     
     modulator := phasor.New(0.0, 0.0)
     modulator.SetPhaseIncrement(modulator.FrequencyToPhaseIncrement(120.0, float64(sampleRate)))

     lfo := phasor.New(0.0, 0.0)
     lfo.SetPhaseIncrement(lfo.FrequencyToPhaseIncrement(0.1, float64(sampleRate)))

     lfo2 := phasor.New(0.0, 0.0)
     lfo2.SetPhaseIncrement(lfo.FrequencyToPhaseIncrement(0.03, float64(sampleRate)))

     var out float64 = 0.0

     for {
     	 // synthesize audio buffer
     	 for i := 0; i < bufferSize; i++ {
	     // compute output sample
	     phase := carrier.GetPhase()
	     phase += math.Sin(modulator.GetPhase()) * (3.5 + 2.5 * math.Sin(lfo.GetPhase()))

	     // operator phase feedback
	     phase += out * (0.5 + 0.5 * math.Sin(lfo2.GetPhase()))
	     
	     // tick all states
     	     carrier.Tick()
     	     modulator.Tick()
     	     lfo.Tick()	     
     	     lfo2.Tick()	     
	     
	     out = math.Sin(phase)

	     // write output to driver
	     b := int16(out * 0.1 * 32767)

	     // left channel
	     data[2*2*i] = byte(b)
	     data[2*2*i+1] = byte(b >> 8)
	     
	     // right channel
	     data[2*2*i+2] = byte(b)
	     data[2*2*i+1+2] = byte(b >> 8)
     	 }

	 // write buffer to ALSA
     	 numBytes, writeError := driver.TryWrite(data)
	 if writeError != nil {
	    fmt.Println(writeError)
	    os.Exit(1)
	 }
	 if numBytes != 2*2*bufferSize {
	    fmt.Printf("[+] Wrote %d bytes to ALSA buffer\n", numBytes)
	 }
     }
     
     driver.Close()
}
