package dmx

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro/messages"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

// DMXReaderService handles reading from DMX
type DMXReaderService struct {
	reader *dmxusbpro.EnttecDMXUSBProController
}

// NewDMXReaderService creates a new DMXReaderService
//
// Connects to an [EnttecDMXUSBProController] using the Port specified in [AppOptions]
func NewDMXReaderService() (service *DMXReaderService) {
	service = &DMXReaderService{}
	service.ConnectDMX()
	return
}

// OnDMXChange grants direct access to the DMX reader's 'OnDMXChange'
//
// Read from DMX and get the results back via channel.
// Call this function as goroutine as it is blocking!
func (s *DMXReaderService) OnDMXChange(c chan messages.EnttecDMXUSBProApplicationMessage) {
	s.reader.OnDMXChange(c, 15)
}

// Connect to DMX
func (s *DMXReaderService) ConnectDMX() {
	opts := options.GetAppOptions()
	channels := opts.DmxChannelCount
	port := opts.DmxReadPort
	baud := opts.DmxReadBaudrate
	log.Infof("Opening DMX Serial for READING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	reader := dmxusbpro.NewEnttecDMXUSBProController(config, channels, false)
	reader.SetLogVerbosity(opts.DmxLogLevel)
	if err := reader.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for READING: %s", err)
		return
	}
	s.reader = reader
	if opts.DmxBridge {
		log.Infof("Switching Read Mode to 'changes only'")
		s.reader.SwitchReadMode(1)
	}
}

// Disconnect from DMX
func (s *DMXReaderService) DisconnectDMX() {
	if s.reader != nil {
		log.Debugf("Shutting down DMX reader...")
		if err := s.reader.Disconnect(); err != nil {
			log.Fatal("Error disconnecting DMX reader:", err)
		} else {
			log.Infof("DMX reader was shut down gracefully")
		}
	}
}
