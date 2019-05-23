package main

import (
    "gosynth/modules/operator"
    "gosynth/io/alsa"
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
     carrier := operator.New(0.0, 80.0, float64(sampleRate))
     modulator := operator.New(0.0, 160.0, float64(sampleRate))
     lfo := operator.New(0.0, 0.1, float64(sampleRate))
     lfo2 := operator.New(0.0, 0.0667, float64(sampleRate))

     var phasemod float64 = 0.0
     var out float64 = 0.0
     
     for {
     	 // synthesize audio buffer
     	 for i := 0; i < bufferSize; i++ {
	     // compute output sample
	     phasemod = modulator.GetOutput() *
	     		(8.0 + 8.0 * lfo.GetOutput())
	     phasemod += out * (0.5 + 0.5 * lfo2.GetOutput())
	     carrier.SetPhaseModulation(phasemod)
	     
	     // tick all states
     	     carrier.Tick()
     	     modulator.Tick()
     	     lfo.Tick()
     	     lfo2.Tick()
	     
	     out = carrier.GetOutput()

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
