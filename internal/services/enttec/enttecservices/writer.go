// Package enttecservices provides services to interact with the Enttec DMX USB Pro Controller
package enttecservices

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

// DMX Writer writes values to DMX using .
type WriterServiceEnttecDMXUSBPro struct {
	writer *dmxusbpro.EnttecDMXUSBProController
}

// Create a WriterServiceEnttecDMXUSBPro
func NewWriterServiceEnttecDMXUSBPro() *WriterServiceEnttecDMXUSBPro {
	log.Debugf("Creating new DMXWriterService")
	f := &WriterServiceEnttecDMXUSBPro{}
	f.ConnectDMX()
	return f
}

// Write values to DMX
func (f *WriterServiceEnttecDMXUSBPro) Write(values map[int16]byte) {
	// FLAG to help us detect if we need to write to DMX
	for i := range values {
		f.writer.Stage(i, values[i])
	}
	f.writer.Commit()
}

// Connect to DMX
func (s *WriterServiceEnttecDMXUSBPro) ConnectDMX() {
	opts := options.GetAppOptions()

	if ok, objection := opts.CanWriteDMX(); !ok {
		log.Infof("%s - Skipping DMX Writer Creation", objection)
		return
	}

	channels := opts.DmxChannelCount
	port := opts.DmxWritePort
	baud := opts.DmxWriteBaudrate
	log.Infof("Opening DMX Serial for WRITING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	writer := dmxusbpro.NewEnttecDMXUSBProController(config, channels, true)
	writer.SetLogVerbosity(opts.DmxLogLevel)
	if err := writer.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for WRITING: %s", err)
	}
	s.writer = writer
}

// Disconnect from DMX
func (s *WriterServiceEnttecDMXUSBPro) DisconnectDMX() {
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
