// Package fader provides a services to fade DMX writes
package fader

import (
	"time"

	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	"github.com/H3rby7/dmx-web-go/internal/options"
	enttecinterfaces "github.com/H3rby7/dmx-web-go/internal/services/enttec/interfaces"
	log "github.com/sirupsen/logrus"
)

// DMX Writer that takes care of fading channels to the desired values over time.
type FadingService struct {
	isActive bool
	writer   enttecinterfaces.EnttecDMXWriter
	faders   []models_fader.DMXFader
}

// Create a new fading writer with the current DMX stage
func NewFadingService(writer enttecinterfaces.EnttecDMXWriter) *FadingService {
	log.Debugf("Creating new FadingService")
	opts := options.GetAppOptions()
	f := &FadingService{
		isActive: false,
		writer:   writer,
		faders:   make([]models_fader.DMXFader, opts.DmxChannelCount+1),
	}
	for i := range f.faders {
		f.faders[i] = models_fader.NewDMXFader(int16(i))
	}
	if ok, objection := opts.CanWriteDMX(); ok {
		f.Start()
	} else {
		log.Warnf("%s - Skipping Start", objection)
	}
	return f
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

// Stop the update loop go-routine
func (f *FadingService) CleanUp() {
	log.Infof("Cleaning up FadingService")
	f.writer.DisconnectDMX()
}

// Blocking loop that calculates a nd runs updates on the faders.
func (f *FadingService) loop() {
	log.Infof("Started fading writer")
	f.isActive = true
	for f.isActive {
		// FLAG to help us detect if we need to write to DMX
		dirty := false
		values := make(map[int16]byte, len(f.faders))
		for i := range f.faders {
			if f.faders[i].IsActive() {
				dirty = true
				values[int16(i)] = f.faders[i].UpdateValue()
			}
		}
		if dirty {
			f.writer.Write(values)
		}
		time.Sleep(time.Millisecond * models_fader.TICK_INTERVAL_MILLIS)
	}
	log.Infof("Stopped fading writer")
}
