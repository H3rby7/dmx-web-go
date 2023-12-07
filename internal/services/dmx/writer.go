package dmx

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

// DMXReaderService handles reading from DMX
type DMXWriterService struct {
	writer *dmxusbpro.EnttecDMXUSBProController
}

// NewDMXWriterService creates a new DMXWriterService
//
// Connects to an [EnttecDMXUSBProController] using the Port specified in [AppOptions]
func NewDMXWriterService() (service *DMXWriterService) {
	opts := options.GetAppOptions()
	if ok, objection := opts.CanWriteDMX(); !ok {
		log.Warnf("%s - skipping", objection)
		return
	}
	service = &DMXWriterService{}
	service.ConnectDMX()
	return
}

// GetStage grants direct access to the DMX writer's 'GetStage'
//
// Gets a copy of all staged channel values
func (s *DMXWriterService) GetStage() []byte {
	return s.writer.GetStage()
}

// Stage grants direct access to the DMX writer's 'Stage'
//
// # Prepare a channel to be changed to the given value
//
// Note: This does not send out the changes, you must call the 'Commit' method to apply the stage live.
func (s *DMXWriterService) Stage(channel int16, value byte) error {
	return s.writer.Stage(channel, value)
}

// Commit grants direct access to the DMX writer's 'Commit'
//
// Apply the 'staged' values to go live.
//
// Note: This does not clear the Stage!
func (s *DMXWriterService) Commit() error {
	return s.writer.Commit()
}

// Connect to DMX
func (s *DMXWriterService) ConnectDMX() {
	opts := options.GetAppOptions()
	channels := opts.DmxChannelCount
	port := opts.DmxWritePort
	baud := opts.DmxWriteBaudrate
	log.Infof("Opening DMX Serial for WRITING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	writer := dmxusbpro.NewEnttecDMXUSBProController(config, channels, true)
	writer.SetLogVerbosity(logVerbosity())
	if err := writer.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for WRITING: %s", err)
	}
	s.writer = writer
}

// Disconnect from DMX
func (s *DMXWriterService) DisconnectDMX() {
	if s.writer != nil {
		log.Debugf("Shutting down DMX writer...")
		shouldClear := options.GetAppOptions().DmxClearOnQuit
		if shouldClear {
			log.Infof("Clearing DMX output to zeros")
			s.writer.ClearStage()
			s.writer.Commit()
		} else {
			log.Debugf("Skipping DMX output cleanup")
		}
		if err := s.writer.Disconnect(); err != nil {
			log.Fatal("Error disconnecting DMX writer:", err)
		} else {
			log.Infof("DMX writer was shut down gracefully")
		}
	}
}
