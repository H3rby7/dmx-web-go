package dmx

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func createReader() (reader *dmxusbpro.EnttecDMXUSBProController) {
	opts := options.GetAppOptions()
	if !opts.HasDMXReader() {
		log.Warnf("No port specified to READ from DMX - skipping")
		return
	}
	channels := opts.DmxChannelCount
	port := opts.DmxReadPort
	baud := opts.DmxReadBaudrate
	log.Infof("Opening DMX Serial for READING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	reader = dmxusbpro.NewEnttecDMXUSBProController(config, channels, false)
	reader.SetLogVerbosity(logVerbosity())
	if err := reader.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for READING: %s", err)
	}
	return
}

func shutdownReader(reader *dmxusbpro.EnttecDMXUSBProController) {
	if reader != nil {
		log.Debugf("Shutting down DMX reader...")
		if err := reader.Disconnect(); err != nil {
			log.Fatal("Error disconnecting DMX reader:", err)
		} else {
			log.Infof("DMX reader was shut down gracefully")
		}
	}
}
