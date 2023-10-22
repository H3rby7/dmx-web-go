package dmx

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func createWriter() *dmxusbpro.EnttecDMXUSBProController {
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
	return writer
}

func shutdownWriter(writer *dmxusbpro.EnttecDMXUSBProController) {
	log.Debugf("Shutting down DMX writer...")
	shouldClear := options.GetAppOptions().DmxClearOnQuit
	if shouldClear {
		log.Infof("Clearing DMX output to zeros")
		writer.ClearStage()
		writer.Commit()
	} else {
		log.Debugf("Skipping DMX output cleanup")
	}
	if err := writer.Disconnect(); err != nil {
		log.Fatal("Error disconnecting DMX writer:", err)
	} else {
		log.Infof("DMX writer was shut down gracefully")
	}
}
