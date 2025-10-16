// Package printer defines a Service reading from DMX and printing to log
package printer

import (
	"github.com/H3rby7/dmx-web-go/internal/services/reader"
	log "github.com/sirupsen/logrus"
)

// DMXLoggerService reads from DMX and logs the received changesets
type DMXLoggerService struct {
	reader reader.DMXReader
}

// NewDMXLoggerService creates a new [DMXLoggerService] instance with proper defaults
func NewDMXLoggerService(reader reader.DMXReader) *DMXLoggerService {
	log.Debugf("Creating new DMXLoggerService")
	b := &DMXLoggerService{
		reader: reader,
	}
	return b
}

// Register On-DMX-Change Channel
// Call this function as goroutine as it is blocking!
func (b *DMXLoggerService) PrintDMX() {
	c := make(chan map[int]byte)
	go b.reader.OnDMXChange(c)
	for cs := range c {
		log.Debugf("%v", cs)
	}
}
