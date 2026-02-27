// Package enttecmock defines enttec mocked DMX services that do not require actual connected DMX hardware.
package enttecmock

import (
	log "github.com/sirupsen/logrus"
)

// MockedDMXReaderService mocks reading from DMX
type MockedDMXReaderService struct {
	reader chan map[int]byte
}

// NewMockedDMXReaderService creates a new [MockedDMXReaderService]
func NewMockedDMXReaderService() (service *MockedDMXReaderService) {
	log.Debugf("Creating new MockedDMXReaderService")
	service = &MockedDMXReaderService{}
	service.ConnectDMX()
	return
}

// OnDMXChange grants direct access to the DMX reader's 'OnDMXChange'
//
// Calls to 'ReadChangeset' are the mocked DMX input, that is dispatched via channel.
// Call this function as goroutine as it is blocking!
func (s *MockedDMXReaderService) OnDMXChange(c chan map[int]byte) {
	for msg := range s.reader {
		c <- msg
	}
}

// ConnectDMX creates the internal reader channel.
func (s *MockedDMXReaderService) ConnectDMX() {
	log.Debugf("Creating mocked DMX reader channel...")
	s.reader = make(chan map[int]byte)
	log.Infof("Created mocked DMX reader channel")
}

// DisconnectDMX closes the internal reader channel.
func (s *MockedDMXReaderService) DisconnectDMX() {
	if s.reader != nil {
		log.Debugf("Closing mocked DMX reader channel...")
		close(s.reader)
		log.Infof("Closed mocked DMX reader channel")
	}
}

// ReadChangeset does the mock magic by 'reading' the input and passing it to the application via the 'OnDMXChange' channel.
func (s *MockedDMXReaderService) ReadChangeset(mockedInput map[int]byte) {
	s.reader <- mockedInput
}
