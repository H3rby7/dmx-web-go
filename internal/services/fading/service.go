package fading

import (
	"time"

	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/dmx"
	log "github.com/sirupsen/logrus"
)

// DMX Writer that takes care of fading channels to the desired values over time.
type FadingService struct {
	isActive bool
	writer   *dmx.DMXWriterService
	faders   []models_fader.DMXFader
}

// Create a new fading writer with the current DMX stage
func NewFadingService(writer *dmx.DMXWriterService) *FadingService {
	if writer == nil {
		log.Panicf("Writer is nil, cannot create FadingWriter.")
	}
	opts := options.GetAppOptions()
	f := &FadingService{
		isActive: false,
		writer:   writer,
		faders:   make([]models_fader.DMXFader, opts.DmxChannelCount+1),
	}
	for i := range f.faders {
		f.faders[i] = models_fader.NewDMXFader(int16(i))
	}
	f.GetUpdateFromWriter()
	return f
}

/*
Get current stage from the writer and update internal faders with it

A call to this function from the outside is only necessary, if the fading writer is not using the actual writer exclusively.
*/
func (f *FadingService) GetUpdateFromWriter() {
	log.Infof("Updating fader with values from writer stage")
	stage := f.writer.GetStage()
	for i := range f.faders {
		log.Tracef("Updated channel '%v' to '%v'", i, stage[i])
		f.faders[i].SetValue(float32(stage[i]))
	}
}

// Fade a given channel to a given value over a given duration
func (f *FadingService) FadeTo(channel int16, value byte, fadeDurationMillis int64) {
	highestChannel := int16(len(f.faders))
	if channel < 1 || channel > highestChannel {
		log.Errorf("Skipping update for channel '%d', because it is out of range, must be between 1 and %d", channel, highestChannel)
		return
	}
	f.faders[channel].FadeTo(value, fadeDurationMillis)
}

// Fade a given scene over a given duration
func (f *FadingService) FadeScene(scene models_scene.Scene, fadeDurationMillis int64) {
	log.Debugf("Using a fade duration of %d millis, setting scene: %v", fadeDurationMillis, scene.List)
	for _, entry := range scene.List {
		f.FadeTo(entry.Channel, entry.Value, fadeDurationMillis)
	}
}

// Immediately set all DMX values to 0
func (f *FadingService) ClearAll() {
	for i := range f.faders {
		f.faders[i].FadeTo(0, 0)
	}
}

// Start a go-routine that runs the update loop
func (f *FadingService) Start() {
	go f.loop()
}

// Stop the update loop go-routine
func (f *FadingService) Stop() {
	log.Infof("Stopping FadingService")
	f.isActive = false
}

// Blocking loop that calculates a nd runs updates on the faders.
func (f *FadingService) loop() {
	log.Infof("Started fading writer")
	f.isActive = true
	for f.isActive {
		// FLAG to help us detect if we need to write to DMX
		dirty := false
		for i := range f.faders {
			if f.faders[i].IsActive() {
				dirty = true
				f.writer.Stage(int16(i), f.faders[i].UpdateValue())
			}
		}
		if dirty {
			f.writer.Commit()
		}
		time.Sleep(time.Millisecond * models_fader.TICK_INTERVAL_MILLIS)
	}
	log.Infof("Stopped fading writer")
}
