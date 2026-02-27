// Package enttecmock defines enttec mocked DMX services that do not require actual connected DMX hardware.
package enttecmock

import (
	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	log "github.com/sirupsen/logrus"
)

// Mocked DMX Writer that manages state of fading channels and logging their values upon change.
type MockedDmxWriterService struct {
	isActive bool
	faders   []models_fader.DMXFader
}

// Create a MockedDmxWriterService
func NewMockedDmxWriterService() *MockedDmxWriterService {
	log.Debugf("Creating new MockedDMXWriterService")
	f := &MockedDmxWriterService{}
	f.ConnectDMX()
	return f
}

// Write values to DMX
func (f *MockedDmxWriterService) Write(values map[int16]byte) {
	f.logFaders()
}

// Connect to DMX
func (s *MockedDmxWriterService) ConnectDMX() {
	log.Infof("Connected mocked DMX writer")
}

// Disconnect from DMX
func (s *MockedDmxWriterService) DisconnectDMX() {
	log.Infof("Disconnected mocked DMX writer")
}

// logFaders prints out all faders to the log that are active or have values > 0
func (s *MockedDmxWriterService) logFaders() {
	m := make(map[int]byte)
	for i := range s.faders {
		f := s.faders[i]
		if f.IsActive() || f.GetCurrentValue() > 0 {
			m[i] = f.GetCurrentValue()
		}
	}
	log.Infof("%v", m)
}
