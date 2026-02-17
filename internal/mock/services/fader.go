package mockservice

import (
	"time"

	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	"github.com/H3rby7/dmx-web-go/internal/options"
	log "github.com/sirupsen/logrus"
)

// DMX Writer that takes care of fading channels to the desired values over time.
type MockedFadingService struct {
	isActive bool
	faders   []models_fader.DMXFader
}

// Create a new MockedFadingService with an empty stage
func NewMockedFadingService() *MockedFadingService {
	log.Debugf("Creating new FadingService")
	opts := options.GetAppOptions()
	s := &MockedFadingService{
		isActive: false,
		faders:   make([]models_fader.DMXFader, opts.DmxChannelCount+1),
	}
	for i := range s.faders {
		s.faders[i] = models_fader.NewDMXFader(int16(i))
	}
	s.Start()
	return s
}

// Fade a given channel to a given value over a given duration
func (s *MockedFadingService) FadeTo(channel int16, value byte, fadeDurationMillis int64) {
	highestChannel := int16(len(s.faders))
	if channel < 1 || channel > highestChannel {
		log.Errorf("Skipping update for channel '%d', because it is out of range, must be between 1 and %d", channel, highestChannel)
		return
	}
	s.faders[channel].FadeTo(value, fadeDurationMillis)
}

// Fade a given scene over a given duration
func (s *MockedFadingService) FadeScene(scene models_scene.Scene, fadeDurationMillis int64) {
	log.Debugf("Using a fade duration of %d millis, setting scene: %v", fadeDurationMillis, scene.List)
	for _, entry := range scene.List {
		s.FadeTo(entry.Channel, entry.Value, fadeDurationMillis)
	}
}

// Immediately set all DMX values to 0
func (s *MockedFadingService) ClearAll() {
	for i := range s.faders {
		s.faders[i].FadeTo(0, 0)
	}
}

// Start a go-routine that runs the update loop
func (s *MockedFadingService) Start() {
	go s.loop()
}

// Stop the update loop go-routine
func (s *MockedFadingService) Stop() {
	log.Infof("Stopping loop")
	s.isActive = false
}

// Blocking loop that calculates and runs updates on the faders.
func (s *MockedFadingService) loop() {
	log.Infof("Started loop")
	s.isActive = true
	for s.isActive {
		// FLAG to help us detect if we need to write to DMX
		dirty := false
		for i := range s.faders {
			if s.faders[i].IsActive() {
				dirty = true
				s.faders[i].UpdateValue()
			}
		}
		if dirty {
			s.logFaders()
		}
		time.Sleep(time.Millisecond * models_fader.TICK_INTERVAL_MILLIS)
	}
	log.Infof("Stopped loop")
}

// Connect to DMX
func (s *MockedFadingService) ConnectDMX() {
	log.Info("Connect faked")
}

// Disconnect from DMX
func (s *MockedFadingService) DisconnectDMX() {
	log.Info("Disconnect faked")
}

// logFaders prints out all faders to the log that are active or have values > 0
func (s *MockedFadingService) logFaders() {
	m := make(map[int]byte)
	for i := range s.faders {
		f := s.faders[i]
		if f.IsActive() || f.GetCurrentValue() > 0 {
			m[i] = f.GetCurrentValue()
		}
	}
	log.Infof("%v", m)
}
