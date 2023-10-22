package dmxfader

import (
	"time"

	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
)

// 25 per second
const TICK_INTERVAL_MILLIS = 1000 / 25

type FadingWriter struct {
	isActive bool
	writer   *dmxusbpro.EnttecDMXUSBProController
	faders   []DMXFader
}

// Create a new fading writer with the current DMX stage
func NewFadingWriter(writer *dmxusbpro.EnttecDMXUSBProController) *FadingWriter {
	opts := options.GetAppOptions()
	f := &FadingWriter{
		isActive: false,
		writer:   writer,
		faders:   make([]DMXFader, opts.DmxChannelCount+1),
	}
	for i := range f.faders {
		f.faders[i] = DMXFader{
			isActive: false,
			channel:  int16(i),
		}
	}
	f.GetUpdateFromWriter()
	return f
}

/*
Get current stage from the writer and update internal faders with it

A call to this function from the outside is only necessary, if the fading writer is not using the actual writer exclusively.
*/
func (f *FadingWriter) GetUpdateFromWriter() {
	log.Infof("Updating fader with values from writer stage")
	stage := f.writer.GetStage()
	for i := range f.faders {
		log.Tracef("Updated channel '%v' to '%v'", i, stage[i])
		f.faders[i].currentValue = float32(stage[i])
	}
}

// Fade a given channel to a given value over a given duration
func (f *FadingWriter) FadeTo(channel int16, value byte, fadeDurationMillis int64) {
	f.faders[channel].FadeTo(value, fadeDurationMillis)
}

// Immediately set all DMX values to 0
func (f *FadingWriter) ClearAll() {
	for i := range f.faders {
		f.faders[i].FadeTo(0, 0)
	}
}

func (f *FadingWriter) Start() {
	go f.loop()
}

func (f *FadingWriter) Stop() {
	f.isActive = false
}

func (f *FadingWriter) loop() {
	log.Infof("Started fading writer")
	f.isActive = true
	for f.isActive {
		dirty := false
		for i := range f.faders {
			if f.faders[i].IsActive() {
				dirty = true
				f.writer.Stage(int16(i), f.faders[i].GetNextValue())
			}
		}
		if dirty {
			f.writer.Commit()
		}
		time.Sleep(time.Millisecond * TICK_INTERVAL_MILLIS)
	}
	log.Infof("Stopped fading writer")
}
